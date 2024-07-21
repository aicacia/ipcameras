package playback

import (
	"log"
	"log/slog"
	"os"
	"path"
	"strconv"
	"sync"
	"time"

	"github.com/aicacia/ipcameras/api/app/service/format"
	"github.com/deepch/vdk/av"
)

const playerChanSize = 1024

const (
	PlaybackBackward int8 = -1
	PlaybackForward  int8 = 1
)

type Player struct {
	mutex       sync.RWMutex
	currentTime *time.Time
	direction   int8
	rate        float32
	paused      bool
	demuxers    []*format.Demuxer
	closed      bool
	stream      chan *av.Packet
}

func NewPlayer(folder string, currentTime *time.Time, direction int8, rate float32) (*Player, error) {
	entries, err := os.ReadDir(folder)
	if err != nil {
		return nil, err
	}
	demuxers := make([]*format.Demuxer, len(entries))
	for _, entry := range entries {
		if !entry.IsDir() {
			idxInt64, err := strconv.ParseInt(entry.Name(), 10, 8)
			if err != nil {
				return nil, err
			}
			file, err := os.OpenFile(path.Join(folder, entry.Name()), os.O_RDONLY, 0644)
			if err != nil {
				return nil, err
			}
			demuxer, err := format.NewDemuxer(file, int8(idxInt64))
			if err != nil {
				return nil, err
			}
			demuxers[int(idxInt64)] = demuxer
		}
	}
	return &Player{
		currentTime: currentTime,
		direction:   direction,
		rate:        rate,
		paused:      false,
		demuxers:    demuxers,
		closed:      false,
		stream:      make(chan *av.Packet, playerChanSize),
	}, nil
}

func (p *Player) Start() {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	for idx := range p.demuxers {
		go p.playDemuxer(int8(idx))
	}
}

func (p *Player) Codecs() []av.CodecData {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	var codecs []av.CodecData
	for _, demuxer := range p.demuxers {
		codecs = append(codecs, demuxer.Codec())
	}
	return codecs
}

func (p *Player) Close() error {
	if !p.closed {
		p.mutex.Lock()
		defer p.mutex.Unlock()
		for _, demuxer := range p.demuxers {
			if demuxer != nil {
				demuxer.Close()
			}
		}
		p.closed = true
		close(p.stream)
	}
	return nil
}

func (p *Player) IsClosed() bool {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	return p.closed
}

func (p *Player) Stream() chan *av.Packet {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	return p.stream
}

func (p *Player) Codec(idx int8) av.CodecData {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	return p.demuxers[idx].Codec()
}

func (p *Player) Direction() int8 {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	return p.direction
}

func (p *Player) SeekToTime(seekMicroseconds int64) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	for _, demuxer := range p.demuxers {
		if demuxer != nil {
			packet, err := demuxer.SeekToTime(seekMicroseconds)
			if err != nil {
				return err
			}
			p.stream <- packet
			return nil
		}
	}
	return nil
}

func (p *Player) readPacket(idx int8) (*av.Packet, error) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	return p.demuxers[idx].ReadPacket(p.direction)
}

func (p *Player) setCurrentTime(t *time.Time) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.currentTime = t
}

func (p *Player) GetCurrentTime() *time.Time {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	return p.currentTime
}

func (p *Player) getRate() float32 {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	if p.paused {
		return 1
	} else {
		return p.rate
	}
}

func (p *Player) isPaused() bool {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	return p.paused
}

func (p *Player) isPlaying() bool {
	return !p.isPaused()
}

func (p *Player) SetRate(rate float32) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.rate = rate
}

func (p *Player) Play() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.paused = false
}

func (p *Player) Pause() {
	p.mutex.Lock()
	p.paused = true
	p.mutex.Unlock()
	p.SeekToTime(p.currentTime.Unix())
}

func (p *Player) close(idx int8) (err error) {
	p.mutex.RLock()
	demuxer := p.demuxers[idx]
	p.mutex.RUnlock()
	if demuxer != nil {
		err = demuxer.Close()
		p.mutex.Lock()
		p.demuxers[idx] = nil
		p.mutex.Unlock()
		if demuxer.Codec().Type().IsAudio() {
			p.mutex.RLock()
			hasVideo := false
			for _, demuxer := range p.demuxers {
				if demuxer != nil && demuxer.Codec().Type().IsVideo() {
					hasVideo = true
					break
				}
			}
			p.mutex.RUnlock()
			if hasVideo {
				return
			}
		}
		if err != nil {
			log.Println("Player demuxer close error", err)
		}
	}
	err = p.Close()
	return
}

func (p *Player) playDemuxer(idx int8) {
	isVideo := p.Codec(idx).Type().IsVideo()
	direction := p.direction
	started := false
	var currentKeyframe *av.Packet
	var currentKeyframeTime time.Time
	for {
		packet := currentKeyframe
		packetTime := currentKeyframeTime
		if p.isPlaying() || packet == nil {
			var err error
			if packet, err = p.readPacket(idx); err != nil {
				slog.Error("codec failed to read packet", "codec", idx, "error", err)
				break
			}
			packetTime = format.GetPacketTime(packet)
			if isVideo {
				p.setCurrentTime(&packetTime)
				if !started {
					if packet.IsKeyFrame {
						started = true
					} else {
						continue
					}
				}
				if packet.IsKeyFrame {
					currentKeyframe = packet
					currentKeyframeTime = packetTime
				}
			} else {
				currentKeyframe = packet
				currentKeyframe.Data = nil
				currentKeyframeTime = packetTime
			}
		}
		var sleepTime time.Duration
		if direction == PlaybackForward {
			sleepTime = packetTime.Sub(*p.GetCurrentTime()) - packet.Duration
		} else if direction == PlaybackBackward {
			sleepTime = p.GetCurrentTime().Sub(packetTime) - packet.Duration
		}
		p.stream <- packet
		if sleepTime > 0 {
			sleepTime = time.Duration(float32(sleepTime) / p.getRate())
			time.Sleep(sleepTime)
		}
		if p.IsClosed() {
			break
		}
	}
	if isVideo {
		p.Close()
	} else {
		p.close(idx)
	}
}

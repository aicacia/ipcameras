package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"log/slog"

	"github.com/aicacia/go-peer"
	"github.com/deepch/vdk/av"
	"github.com/deepch/vdk/codec/h264parser"
	"github.com/pion/webrtc/v4"
	"github.com/pion/webrtc/v4/pkg/media"
)

var (
	ErrorNotFound          = errors.New("WebRTC Stream Not Found")
	ErrorCodecNotSupported = errors.New("WebRTC Codec Not Supported")
	ErrorClientOffline     = errors.New("WebRTC Client Offline")
	ErrorNotTrackAvailable = errors.New("WebRTC Not Track Available")
	ErrorIgnoreAudioTrack  = errors.New("WebRTC Ignore Audio Track codec not supported WebRTC support only PCM_ALAW or PCM_MULAW")
)

func CodecsToStrings(codecs []av.CodecData) []string {
	var out []string
	for _, codec := range codecs {
		if codec.Type() != av.H264 && codec.Type() != av.PCM_ALAW && codec.Type() != av.PCM_MULAW && codec.Type() != av.OPUS {
			log.Println("Codec Not Supported WebRTC ignore this track", codec.Type())
			continue
		}
		if codec.Type().IsVideo() {
			out = append(out, "video")
		} else {
			out = append(out, "audio")
		}
	}
	return out
}

func InitRTSP() {
	PeersOnConnect.Append(func(id, kind string, p *peer.Peer) {
		if kind == "webrtcrtsp" {
			var onData peer.OnData
			onData = func(message webrtc.DataChannelMessage) {
				p.OffData(onData)
				var data map[string]interface{}
				err := json.Unmarshal(message.Data, &data)
				if err != nil {
					slog.Error("error parsing messsage: %v", "error", err)
					return
				}
				if kind, ok := data["type"].(string); ok {
					switch kind {
					case "rtsp":
						if rtspUrl, ok := data["rtspUrl"].(string); ok {
							go streamRTSPURL(p, rtspUrl)
						} else {
							slog.Error("invalid message: %v", "data", data)
						}
					}
				} else {
					slog.Error("invalid message: %v", "data", data)
				}
			}
			p.OnData(onData)
		}
	})
}

type stream struct {
	track *webrtc.TrackLocalStaticSample
	codec av.CodecData
}

func streamRTSPURL(peer *peer.Peer, rtspUrl string) {
	err := RunIfNotRunning(rtspUrl)
	if err != nil {
		slog.Error("error streaming rtsp: %v", "error", err)
		return
	}
	codecs := GetClientCodecs(rtspUrl)
	var streams []stream
	for _, codec := range codecs {
		var track *webrtc.TrackLocalStaticSample
		if codec.Type().IsVideo() {
			if codec.Type() == av.H264 {
				track, err = webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{
					MimeType: webrtc.MimeTypeH264,
				}, "pion-rtsp-video", "pion-video")
				if err != nil {
					slog.Error("error creating webrtc track: %v", "error", err)
					break
				}
				if rtpSender, err := peer.AddTrack(track); err != nil {
					slog.Error("error adding webrtc track: %v", "error", err)
					break
				} else {
					go func() {
						rtcpBuf := make([]byte, 1500)
						for {
							if _, _, rtcpErr := rtpSender.Read(rtcpBuf); rtcpErr != nil {
								return
							}
						}
					}()
					slog.Info("Added H.264 video track")
					streams = append(streams, stream{
						track: track,
						codec: codec,
					})
				}
			}
		} else if codec.Type().IsAudio() {
			AudioCodecString := webrtc.MimeTypePCMA
			switch codec.Type() {
			case av.PCM_ALAW:
				AudioCodecString = webrtc.MimeTypePCMA
			case av.PCM_MULAW:
				AudioCodecString = webrtc.MimeTypePCMU
			case av.OPUS:
				AudioCodecString = webrtc.MimeTypeOpus
			default:
				continue
			}
			track, err = webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{
				MimeType:  AudioCodecString,
				Channels:  uint16(codec.(av.AudioCodecData).ChannelLayout().Count()),
				ClockRate: uint32(codec.(av.AudioCodecData).SampleRate()),
			}, "pion-rtsp-audio", "pion-rtsp-audio")
			if err != nil {
				slog.Error("error creating webrtc track: %v", "error", err)
				break
			}
			if rtpSender, err := peer.AddTrack(track); err != nil {
				slog.Error("error adding webrtc track: %v", "error", err)
				break
			} else {
				go func() {
					rtcpBuf := make([]byte, 1500)
					for {
						if _, _, rtcpErr := rtpSender.Read(rtcpBuf); rtcpErr != nil {
							return
						}
					}
				}()
				slog.Info("Added %s audio track", "AudioCodecString", AudioCodecString)
				streams = append(streams, stream{
					track: track,
					codec: codec,
				})
			}
		}
	}
	peerId := peer.Id()
	viewer := AddViewer(rtspUrl, peerId)
	if err != nil {
		slog.Error("error marshalling codecs: %v", "error", err)
		return
	}
	go func() {
		defer DeleteViewer(rtspUrl, viewer.Id)
		closeSignal := peer.CloseSignal()

		for {
			select {
			case <-closeSignal:
				slog.Debug("Peer closed")
				return
			case packet := <-viewer.Socket:
				if err := writePacket(streams, packet); err != nil {
					slog.Error("error writing packet: %v", "error", err)
					return
				}
			}
		}
	}()
}

func writePacket(streams []stream, packet *av.Packet) error {
	wrotePacket := false
	for _, stream := range streams {
		if len(packet.Data) < 5 {
			return nil
		}
		switch stream.codec.Type() {
		case av.H264:
			nalus, _ := h264parser.SplitNALUs(packet.Data)
			for _, nalu := range nalus {
				naltype := nalu[0] & 0x1f
				var err error
				if naltype == 5 {
					codec := stream.codec.(h264parser.CodecData)
					err = stream.track.WriteSample(media.Sample{Data: append([]byte{0, 0, 0, 1}, bytes.Join([][]byte{codec.SPS(), codec.PPS(), nalu}, []byte{0, 0, 0, 1})...), Duration: packet.Duration})

				} else if naltype == 1 {
					err = stream.track.WriteSample(media.Sample{Data: append([]byte{0, 0, 0, 1}, nalu...), Duration: packet.Duration})
				}
				if err != nil {
					return err
				}
				wrotePacket = true
			}
			return nil
		case av.PCM_ALAW:
		case av.OPUS:
		case av.PCM_MULAW:
		case av.AAC:
			return ErrorCodecNotSupported
		case av.PCM:
			return ErrorCodecNotSupported
		default:
			return ErrorCodecNotSupported
		}
		err := stream.track.WriteSample(media.Sample{Data: packet.Data, Duration: packet.Duration})
		if err == nil {
			wrotePacket = true
		}
	}
	if !wrotePacket {
		slog.Debug("No packet written for stream")
	}
	return nil
}

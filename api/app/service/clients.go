package service

import (
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"sync/atomic"
	"time"

	"github.com/aicacia/go-cmap"
	"github.com/aicacia/go-cslice"
	"github.com/aicacia/go-pubsub"
	"github.com/aicacia/ipcameras/api/app/config"
	"github.com/deepch/vdk/av"
	"github.com/deepch/vdk/codec/h264parser"
	"github.com/deepch/vdk/format/rtspv2"
	"github.com/google/uuid"
)

var (
	ErrClientExitDisconnect = errors.New("client exit disconnect")
	ErrClientNoClient       = errors.New("no client exists")
	ErrUnsupportedScheme    = errors.New("unsupported scheme")
)

var clients = cmap.New[string, *clientST]()

const (
	clientCloseEvent = iota
	clientReadyEvent
)

type clientEvent interface {
	Kind() int
}

type clientEventCloseST struct{}

func (c *clientEventCloseST) Kind() int {
	return clientCloseEvent
}

type clientEventReadyST struct{}

func (c *clientEventReadyST) Kind() int {
	return clientReadyEvent
}

type clientST struct {
	url     url.URL
	ready   atomic.Bool
	running atomic.Bool
	closed  atomic.Bool
	codecs  cslice.CSlice[av.CodecData]
	viewers cmap.CMap[string, *ViewerST]
	pubsub  *pubsub.PubSub[clientEvent]
}

func RunIfNotRunning(clientUrlString string) error {
	client, ok := clients.Get(clientUrlString)
	if !ok || client == nil {
		pubsub := pubsub.NewPubSub[clientEvent]()
		clientUrl, err := url.Parse(clientUrlString)
		if err != nil {
			return err
		}
		client = &clientST{
			url:     *clientUrl,
			pubsub:  &pubsub,
			viewers: cmap.New[string, *ViewerST](),
		}
		clients.Set(clientUrlString, client)
	}
	if client != nil {
		if client.url.Scheme == "rtsp" {
			go rtspWorkerLoop(clientUrlString)
		} else {
			return ErrUnsupportedScheme
		}
	}
	return nil
}

func clientSendClose(url string) {
	if client, ok := clients.Get(url); ok && client != nil && !client.closed.Load() {
		slog.Debug("Sending close signal", "url", url)
		var event clientEvent = &clientEventCloseST{}
		client.pubsub.Publish(event)
	}
}

func clientClose(url string) {
	if client, ok := clients.Get(url); ok && client != nil && client.closed.CompareAndSwap(false, true) {
		var event clientEvent = &clientEventCloseST{}
		client.pubsub.Publish(event)
	}
}

func clientDelete(url string) {
	clientSendClose(url)
	clients.Remove(url)
	slog.Debug("Deleted RTSP", "url", url)
}

func clientSetReady(url string) {
	if client, ok := clients.Get(url); ok && client != nil && client.ready.CompareAndSwap(false, true) {
		var event clientEvent = &clientEventReadyST{}
		client.pubsub.Publish(event)
	}
}

func clientIsClosed(url string) bool {
	if client, ok := clients.Get(url); ok && client != nil {
		return client.closed.Load()
	} else {
		return true
	}
}

func clientCodecsSet(url string, codecs []av.CodecData) {
	if client, ok := clients.Get(url); ok && client != nil {
		client.codecs.Overwrite(codecs)
		clientSetReady(url)
	}
}

func waitForClientReady(url string) bool {
	if client, ok := clients.Get(url); ok && client != nil {
		if client.ready.Load() {
			return true
		} else {
			slog.Debug("Waiting for client to be ready", "url", url)
			s := client.pubsub.Subscribe()
			defer s.Close()
			for e := range s.C {
				switch e.Kind() {
				case clientCloseEvent:
					slog.Debug("Client closed", "url", url)
					return false
				case clientReadyEvent:
					slog.Debug("Client ready", "url", url)
					return true
				}
			}
		}
	}
	return false
}

func GetClientCodecs(url string) []av.CodecData {
	if waitForClientReady(url) {
		if client, ok := clients.Get(url); ok && client != nil {
			validCodecs := client.codecs.Len() > 0
			for codec := range client.codecs.Iter() {
				if codec.Type() == av.H264 {
					codecVideo, ok := codec.(h264parser.CodecData)
					if !ok || codecVideo.SPS() == nil || codecVideo.PPS() == nil || len(codecVideo.SPS()) <= 0 || len(codecVideo.PPS()) <= 0 {
						slog.Debug("Bad Video Codec SPS or PPS Wait", "url", url)
						validCodecs = false
						break
					}
				}
			}
			if validCodecs {
				return client.codecs.Slice()
			}
		}
	}
	slog.Debug("No client to get Codecs", "url", url)
	return nil
}

func GetClientCurrentCodecs(url string) []av.CodecData {
	if client, ok := clients.Get(url); ok && client != nil {
		return client.codecs.Slice()
	}
	return nil
}

const viewerChanSize = 1024

type ViewerST struct {
	Id     string
	Socket chan *av.Packet
}

func AddViewer(url, id string) *ViewerST {
	if client, ok := clients.Get(url); ok && client != nil {
		viewer := &ViewerST{
			Id:     fmt.Sprintf("%s-%s", uuid.New().String(), id),
			Socket: make(chan *av.Packet, viewerChanSize),
		}
		client.viewers.Set(viewer.Id, viewer)
		return viewer
	}
	return nil
}

func DeleteViewer(url, id string) {
	if client, ok := clients.Get(url); ok && client != nil {
		client.viewers.Remove(id)
		if client.viewers.Count() == 0 {
			slog.Debug("No more viewers closing", "url", url)
			clientSendClose(url)
		}
		slog.Debug("Closed RTSP viewer", "url", url, "id", id)
	}
}

func (v *ViewerST) cast(packet *av.Packet) {
	if len(v.Socket) < cap(v.Socket) {
		v.Socket <- packet
	}
}

func cast(url string, packet *av.Packet) {
	if client, ok := clients.Get(url); ok && client != nil {
		for viewer := range client.viewers.Values() {
			viewer.cast(packet)
		}
	}
}

func rtspWorkerLoop(url string) {
	if client, ok := clients.Get(url); ok && client != nil {
		if !client.running.CompareAndSwap(false, true) {
			slog.Debug("Already running RTSP worker loop", "url", url)
			return
		}
	} else {
		slog.Debug("No client to create RTSP worker loop", "url", url)
		return
	}
	defer clientDelete(url)
	for {
		closed, err := rtspWorker(url)
		if err != nil {
			slog.Error("RTSP worker loop Error", "url", url, "error", err)
		}
		if closed || clientIsClosed(url) {
			slog.Debug("RTSP worker loop closed", "url", url)
			break
		}
	}
}

func rtspWorker(url string) (bool, error) {
	rtsp_client, err := rtspv2.Dial(rtspv2.RTSPClientOptions{
		URL:              url,
		DisableAudio:     false,
		DialTimeout:      time.Duration(config.Get().RTSP.ConnectTimeoutSeconds) * time.Second,
		ReadWriteTimeout: time.Duration(config.Get().RTSP.IOTimeoutSeconds) * time.Second,
		Debug:            config.Get().RTSP.Debug,
	})
	if err != nil {
		return true, err
	}
	defer rtsp_client.Close()

	if rtsp_client.CodecData != nil {
		slog.Debug("RTSP worker Codecs", "url", url, "codecs", len(rtsp_client.CodecData))
		clientCodecsSet(url, rtsp_client.CodecData)
	}
	client, ok := clients.Get(url)
	if !ok || client == nil || client.closed.Load() {
		return true, ErrClientNoClient
	}
	s := client.pubsub.Subscribe()
	defer s.Close()
	for {
		select {
		case e := <-s.C:
			switch e.Kind() {
			case clientCloseEvent:
				slog.Debug("RTSP worker Camera close signal", "url", url)
				clientClose(url)
				return true, nil
			}
		case signals := <-rtsp_client.Signals:
			switch signals {
			case rtspv2.SignalCodecUpdate:
				slog.Debug("RTSP worker rtspv2.SignalCodecUpdate Codecs", "url", url, "codecs", len(rtsp_client.CodecData))
				clientCodecsSet(url, rtsp_client.CodecData)
			case rtspv2.SignalStreamRTPStop:
				slog.Debug("RTSP worker rtspv2.SignalClientRTPStop", "url", url)
				return false, ErrClientExitDisconnect
			}
		case packetAV := <-rtsp_client.OutgoingPacketQueue:
			packetAV.Time = time.Duration(time.Now().UTC().UnixNano())
			cast(url, packetAV)
		}
	}
}

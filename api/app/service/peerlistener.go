package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/aicacia/go-cmap"
	"github.com/aicacia/go-cslice"
	"github.com/aicacia/go-peer"
	"github.com/aicacia/ipcameras/api/app/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pion/webrtc/v4"
	"golang.org/x/net/websocket"
)

type PeerOnConnect = func(id, kind string, p *peer.Peer)

var peers = cmap.New[string, *peer.Peer]()
var PeersOnConnect = cslice.New[PeerOnConnect]()

func InitPeerListener() error {
	ctx := context.Background()
	go peerListener(ctx)
	return nil
}

func peerListener(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			slog.Error("peerListener panic", "error", err)
			go peerListener(ctx)
		}
	}()
	for {
		jwtToken, err := authenticate()
		if err != nil {
			slog.Error("failed to create token", "error", err)
			time.Sleep(time.Second * 30)
			continue
		}
		conn, err := websocket.Dial(fmt.Sprintf("%s/server/websocket?token=%s", config.Get().GetP2PWSURL(), url.QueryEscape(jwtToken)), "", config.Get().GetP2PWSURL())
		if err != nil {
			slog.Error("failed to connect WS", "error", err)
			time.Sleep(time.Second * 5)
			continue
		}
		for {
			var msg map[string]interface{}
			err := websocket.JSON.Receive(conn, &msg)
			if err != nil {
				slog.Error("failed to receive message", "error", err)
				break
			}
			switch msg["type"].(string) {
			case "join":
				slog.Debug("joined", "from", msg["from"])
			case "leave":
				slog.Debug("left", "from", msg["from"])
			case "message":
				var from string
				var ok bool
				if from, ok = msg["from"].(string); !ok {
					slog.Debug("invalid message")
					continue
				}
				var payload map[string]interface{}
				if payload, ok = msg["payload"].(map[string]interface{}); !ok {
					slog.Debug("%s: invalid payload", "from", from)
					continue
				}
				var peerId string
				if peerId, ok = payload["peerId"].(string); !ok {
					slog.Debug("%s: invalid payload peerId", "from", from)
					continue
				}
				var payloadType string
				if payloadType, ok = payload["type"].(string); !ok {
					slog.Debug("%s: invalid payload type", "from", from)
					continue
				}
				switch payloadType {
				case "init":
					slog.Debug("init", "from", from, "peer", peerId)
					var peerType string
					if peerType, ok = payload["peerType"].(string); !ok {
						slog.Debug("%s: invalid payload peerType", "from", from)
						continue
					}
					var p *peer.Peer
					ordered := true
					p = peer.NewPeer(peer.PeerOptions{
						Id: peerId,
						ChannelConfig: &webrtc.DataChannelInit{
							Ordered: &ordered,
						},
						Config: &webrtc.Configuration{
							ICEServers: []webrtc.ICEServer{
								{
									URLs: []string{"stun:stun.l.google.com:19302"},
								},
							},
						},
						OnSignal: func(message map[string]interface{}) error {
							return websocket.JSON.Send(conn, map[string]interface{}{
								"to": from,
								"payload": map[string]interface{}{
									"type":    "signal",
									"peerId":  peerId,
									"message": message,
								},
							})
						},
						OnConnect: func() {
							for fn := range PeersOnConnect.Iter() {
								fn(peerId, peerType, p)
							}
						},
						OnClose: func() {
							slog.Debug("%s: closed", "peerId", peerId)
						},
					})
					peers.Set(peerId, p)
					if err := p.Init(); err != nil {
						slog.Error("failed to init peer", "peer", from, "error", err)
						continue
					}
				case "signal":
					var p *peer.Peer
					if p, ok = peers.Get(peerId); !ok {
						slog.Debug("%s: peer not found", "from", from)
						continue
					}
					var message map[string]interface{}
					if message, ok = payload["message"].(map[string]interface{}); !ok {
						slog.Debug("%s: invalid message", "from", from)
						continue
					}
					if err := p.Signal(message); err != nil {
						slog.Error("failed to signal peer", "peer", from, "error", err)
						continue
					}
				}
			}
		}
	}
}

func authenticate() (string, error) {
	requestBody := map[string]string{
		"password": config.Get().P2P.Password,
	}
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", config.Get().GetP2PURL()+"/server", bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().Unix() + 60,
		"sub": config.Get().P2P.Id,
	})
	jwtTokenString, err := jwtToken.SignedString([]byte(config.Get().P2P.JWTSecret))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+jwtTokenString)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", errors.New("failed to authenticate")
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(responseBody), nil
}

package service

import (
	"io"
	"log/slog"
	"net/http"

	"github.com/aicacia/go-peer"
	webrtchttp "github.com/aicacia/go-webrtchttp"
	"github.com/gofiber/fiber/v2"
)

func InitWebRTCServer(fiberApp *fiber.App) {
	servers := map[string]*webrtchttp.WebRTCServerST{}
	PeersOnConnect.Append(func(id, kind string, peer *peer.Peer) {
		if kind == "webrtchttp" {
			peer.OnClose(func() {
				delete(servers, id)
			})
			servers[id] = webrtchttp.NewServer(peer.Channel(), func(w http.ResponseWriter, r *http.Request) {
				webrtcHandler(fiberApp, w, r)
			})
		}
	})
}

func webrtcHandler(fiberApp *fiber.App, responseWriter http.ResponseWriter, webrtcRequest *http.Request) {
	req, err := http.NewRequest(webrtcRequest.Method, webrtcRequest.URL.String(), webrtcRequest.Body)
	if err != nil {
		slog.Error("failed to create request", "err", err)
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("{\"errors\":{\"internal\":[{\"error\":\"application\",\"parameters\":[]}]}}"))
		return
	}
	for name, values := range webrtcRequest.Header {
		for _, value := range values {
			req.Header.Add(name, value)
		}
	}
	res, err := fiberApp.Test(req)
	if err != nil {
		slog.Error("failed to call handler", "err", err)
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("{\"errors\":{\"internal\":[{\"error\":\"application\",\"parameters\":[]}]}}"))
		return
	}
	header := responseWriter.Header()
	for name, values := range res.Header {
		for _, value := range values {
			header.Add(name, value)
		}
	}
	_, err = io.Copy(responseWriter, res.Body)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("{\"errors\":{\"internal\":[{\"error\":\"application\",\"parameters\":[]}]}}"))
		return
	}
	responseWriter.WriteHeader(res.StatusCode)
}

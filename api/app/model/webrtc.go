package model

import "github.com/pion/webrtc/v4"

type OfferBodyST struct {
	OfferBase64 string `json:"offer_base64" validate:"required"`
} // @name OfferBody

type AnswerST struct {
	AnswerBase64 string `json:"answer_base64" validate:"required"`
} // @name Answer

type ICEServerST struct {
	URLs           []string    `json:"urls"`
	Username       string      `json:"username,omitempty"`
	Credential     interface{} `json:"credential,omitempty"`
	CredentialType string      `json:"credentialType,omitempty"`
} // @name ICEServer

func FromWebRTCIceServer(s webrtc.ICEServer) ICEServerST {
	return ICEServerST{
		URLs:           s.URLs,
		Username:       s.Username,
		Credential:     s.Credential,
		CredentialType: s.CredentialType.String(),
	}
}

package model

type LicenseRequest struct {
	Request   string `json:"request"`
	Signature string `json:"signature"`
	Signer    string `json:"signer"`
}

type RequestMessage struct {
	Payload           string `json:"payload"`
	ContentId         string `json:"content_id"`
	Provider          string `json:"provider"`
	AllowedTrackTypes string `json:"allowed_track_types"`
}

type PolicyOverrides struct {
	CanPlay bool `json:"can_play,omitempty"`
}

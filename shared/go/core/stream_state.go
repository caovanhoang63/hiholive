package core

type StreamState struct {
	State     string `json:"state"`
	Uid       string `json:"id"`
	StreamKey string `json:"stream_key,omitempty"`
}

package model

type Metadata map[string]interface{}

type NotifyReq struct {
	DeviceToke string   `json:"device_toke" validate:"required, string"`
	Title      string   `json:"title" validate:"required, string"`
	Message    string   `json:"message"  validate:"required, string"`
	Metadata   Metadata `json:"metadata"`
}


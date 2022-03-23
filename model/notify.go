package model

type Metadata map[string]interface{}

type NotifyReq struct {
	DeviceToke string   `json:"device_toke" validate:"required, string"`
	Title      string   `json:"title" validate:"required, string"`
	Message    string   `json:"message"  validate:"required, string"`
	Metadata   Metadata `json:"metadata"`
}

type MultiNotifyReq struct {
	DeviceTokes []string `json:"device_tokes" validate:"required"`
	Title      string   `json:"title" validate:"required, string"`
	Message    string   `json:"message"  validate:"required, string"`
	Metadata   Metadata `json:"metadata"`
}

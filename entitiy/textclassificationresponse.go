package entitiy

import "github.com/hupe1980/go-huggingface"

type TextClassificationResponse struct {
	Data       huggingface.TextClassificationResponse `json:"data"`
	ServerTime string                                 `json:"server_time"`
	Status     string                                 `json:"status"`
	Error      string                                 `json:"error"`
}

package model

import "encoding/json"

type PromptResponse struct {
	NodeErrors struct{} `json:"node_errors"`
	Number     int      `json:"number"`
	PromptID   string   `json:"prompt_id"`
}

type PromptRequest struct {
	ClientID  string `json:"client_id"`
	ExtraData struct {
		ExtraPngInfo struct {
			WorkFlow json.RawMessage `json:"work_flow"`
		} `json:"extra_pnginfo"`
	} `json:"extra_data"`
	Prompt json.RawMessage `json:"prompt"`
}

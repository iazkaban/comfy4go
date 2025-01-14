package model

import "encoding/json"

type PromptResponse struct {
	NodeErrors struct{}   `json:"node_errors"`
	Number     int        `json:"number"`
	PromptID   string     `json:"prompt_id"`
	Error      *NodeError `json:"error"`
}

type NodeError struct {
	Type      string   `json:"type"`
	Message   string   `json:"message"`
	Details   string   `json:"details"`
	ExtraInfo struct{} `json:"extra_info"`
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

package websocket_message_model

// Progress 是流程处理过程中的进度信息
type Progress struct {
	Value    int    `json:"value"`
	Max      int    `json:"max"`
	PromptID string `json:"prompt_id"`
	Node     string `json:"node"`
}

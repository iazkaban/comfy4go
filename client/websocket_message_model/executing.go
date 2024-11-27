package websocket_message_model

// Executing 是流程处理过程中的跳转信息
type Executing struct {
	Node        string `json:"node"`
	DisplayNode string `json:"display_node"`
	PromptID    string `json:"prompt_id"`
	Timestamp   int64  `json:"timestamp"`
}

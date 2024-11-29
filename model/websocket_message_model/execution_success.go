package websocket_message_model

// ExecutionSuccess 流程处理成功的通知
type ExecutionSuccess struct {
	PromptID  string `json:"prompt_id"`
	Timestamp int64  `json:"timestamp"`
}

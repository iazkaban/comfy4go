package websocket_message_model

// ExecutionStart 一个新的流程请求开始处理的通知
type ExecutionStart struct {
	PromptID  string `json:"prompt_id"`
	Timestamp int64  `json:"timestamp"`
}

package websocket_message_model

// ExecutionCached
type ExecutionCached struct {
	Nodes     []string `json:"nodes"`
	PromptID  string   `json:"prompt_id"`
	Timestamp int64    `json:"timestamp"`
}

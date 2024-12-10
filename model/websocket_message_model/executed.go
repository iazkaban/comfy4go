package websocket_message_model

// Executed 是执行结束之后的通知结构
type Executed struct {
	Node        string  `json:"node"`
	DisplayNode string  `json:"display_node"`
	Output      *Output `json:"output"`
	PromptID    string  `json:"prompt_id"`
}

type Output struct {
	Images []*Image `json:"images,omitempty"`
	Text   []string `json:"text,omitempty"`
}

type Image struct {
	Filename  string `json:"filename"`
	SubFolder string `json:"subfolder"`
	Type      string `json:"type"`
}

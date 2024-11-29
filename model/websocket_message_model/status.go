package websocket_message_model

// Status 是ComfyUI流程队列状态改变的信息
type Status struct {
	SID    string `json:"sid"`
	Status struct {
		ExecInfo struct {
			QueueRemaining int `json:"queue_remaining"`
		} `json:"exec_info"`
	} `json:"status"`
}

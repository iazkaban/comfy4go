package client

import (
	"sync"
)

var (
	websocketMessageProcessorList = make(map[string]func(message *BaseWebsocketMessage) error)
	websocketMessageProcessorLock = sync.RWMutex{}
)

func (client *Client) RegisterMessageProcessor(messageType string, function func(message *BaseWebsocketMessage) error) {
	websocketMessageProcessorLock.Lock()
	defer websocketMessageProcessorLock.Unlock()

	if _, ok := websocketMessageProcessorList[messageType]; ok {
		client.log.Error(ErrorMessageTypeProcessorIsExists)
		return
	}

	websocketMessageProcessorList[messageType] = function
}

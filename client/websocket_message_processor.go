package client

import (
	"errors"
	"sync"
)

var (
	websocketMessageProcessorList = make(map[string]func(message *BaseWebsocketMessage) error)
	websocketMessageProcessorLock = sync.RWMutex{}
)

var (
	ErrorMessageTypeProcessorIsExists = errors.New("this message type processor is Exists")
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

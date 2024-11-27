package client

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"net/url"
)

type BaseWebsocketMessage struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

func newWebsocket(host string, clientID string) (*websocket.Conn, error) {
	websocketUri := "/ws"

	websocketUrl := url.URL{
		Scheme:   "ws",
		Host:     host,
		Path:     websocketUri,
		RawQuery: "clientId=" + clientID,
	}

	conn, resp, err := websocket.DefaultDialer.Dial(websocketUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (client *Client) CloseWebsocket() error {
	return client.WebSocketConnection.Close()
}

func (client *Client) listenWebsocket() {
	defer client.Wg.Done()
	tmp := &BaseWebsocketMessage{}

	for {
		messageType, message, err := client.WebSocketConnection.ReadMessage()

		if messageType == websocket.CloseMessage || messageType <= 0 {
			client.log.Info("server closed connection")
			return
		}
		if err != nil {
			client.log.Error(err.Error() + "messageType:" + fmt.Sprint(messageType))
			return
		}

		err = json.Unmarshal(message, tmp)
		if err != nil {
			client.log.Error(err)
			continue
		}

		go func() {
			websocketMessageProcessorLock.RLock()
			defer websocketMessageProcessorLock.RUnlock()

			if function, ok := websocketMessageProcessorList[tmp.Type]; ok {
				err = function(tmp)
				if err != nil {
					client.log.Error(err)
				}
			} else {
				client.log.Info("type [" + tmp.Type + "]dose not found processor function")
			}
		}()
	}
}

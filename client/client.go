package client

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/iazkaban/comfy4go/log"
)

type Client struct {
	WebSocketConnection *websocket.Conn
	Wg                  *sync.WaitGroup
	ClientID            string
	client              *http.Client
	host                string
	UserID              string
	log                 log.Logger
}

type ClientOption struct {
	Host     string
	Port     int
	Logger   log.Logger
	Wg       *sync.WaitGroup
	ClientID string
}

func NewClient(opt *ClientOption) (*Client, error) {
	host := fmt.Sprintf("%s:%d", opt.Host, opt.Port)
	wsConn, err := newWebsocket(host, opt.ClientID)
	if err != nil {
		return nil, err
	}

	c := &Client{
		client:              http.DefaultClient,
		host:                host,
		WebSocketConnection: wsConn,
		Wg:                  opt.Wg,
		ClientID:            opt.ClientID,
	}

	if opt.Logger != nil {
		c.log = opt.Logger
	} else {
		c.log = &log.SimpleLog{}
	}

	if c.Wg == nil {
		c.Wg = &sync.WaitGroup{}
	}
	c.Wg.Add(1)

	go c.listenWebsocket()

	return c, nil
}

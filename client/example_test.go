package client

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
)

// ExampleNewClient is a sample function to show how to use NewClient
func ExampleNewClient() {
	opt := &ClientOption{
		Host:     "127.0.0.1",
		Port:     8818,
		Wg:       &sync.WaitGroup{},
		ClientID: uuid.New().String(),
	}

	c, err := NewClient(opt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(c.ClientID)
	// Output: 61d52256-add9-4573-8b31-d27c92f6d8bb
}

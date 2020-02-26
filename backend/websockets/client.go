package websockets

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"

	"github.com/jakebjorke/sezzle-interview/models"
)

//Client is a websocket client
type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		message := models.Message{Type: messageType, Body: string(p)}
		c.Pool.Broadcast <- message
		fmt.Printf("message received: %+v\n", message)
	}
}

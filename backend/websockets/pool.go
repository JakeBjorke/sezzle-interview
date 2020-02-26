package websockets

import (
	"log"

	"github.com/jakebjorke/sezzle-interview/history"
	"github.com/jakebjorke/sezzle-interview/models"
)

//Pool is used to distribute values
type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan models.Message
	MsgLog     *history.MessageLog
}

//NewPool creates a new pool
func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan models.Message),
		MsgLog:     history.NewMessageLog(10),
	}
}

//Start is used to run the pool
func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			log.Println("Size of connection pool: ", len(pool.Clients))
			//todo the requirements do not say that we have to populate it initally but could easily do it here...
			// history := pool.MsgLog.GetLog()
			// client.Conn.WriteJSON(history)
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			log.Println("Size of connection pool:  ", len(pool.Clients))
			break
		case message := <-pool.Broadcast:
			log.Println("Sending expression history to all clients in pool")
			pool.MsgLog.Push(message)
			history := pool.MsgLog.GetLog()
			for client := range pool.Clients {
				if err := client.Conn.WriteJSON(history); err != nil {
					log.Println(err)
					return
				}
			}
		}
	}
}

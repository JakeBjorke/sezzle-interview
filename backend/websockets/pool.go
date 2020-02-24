package websockets

import (
	"log"
)

//Pool is used to distribute values
type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

//NewPool creates a new pool
func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

//Start is used to run the pool
func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			log.Println("Size of connection pool: ", len(pool.Clients))
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			log.Println("Size of connection pool:  ", len(pool.Clients))
			break
		case message := <-pool.Broadcast:
			//todo come back and make sure that we limit this to 10 and deal with tossing out the old ones.
			log.Println("Sending message to all clients in pool")
			for client, _ := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					log.Println(err)
					return
				}
			}
		}
	}
}

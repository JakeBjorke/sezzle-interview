package history

import (
	"sync"

	"github.com/jakebjorke/sezzle-interview/models"
)

/*
Really only using a mutex here in case of something 
not being thread safe.  It is intended to be used in the 
web socket pool which uses channels so it _should_ be thread
safe but doing it this way is safer....
*/

//MessageLog is an object used to track messages
type MessageLog struct {
	log []models.Message
	mutex sync.RWMutex
	capacity int
}

//NewMessageLog creates a new message log
func NewMessageLog(capacity int) *MessageLog {
	return &MessageLog{
		log: make([]models.Message, capacity),	
		capacity: capacity,
	}
}


//Push adds a new element 
func (m *MessageLog) Push(msg models.Message) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	/* 
		this is the least efficient way to do this
		but it is only ever 10 items so I am not too concerned 
		about memory usage....
	*/
	m.log = append([]models.Message{msg}, m.log[0:(m.capacity-1)]...)
}

//GetLog outs the log
func(m *MessageLog) GetLog() []models.Message {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.log
}

package history

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jakebjorke/sezzle-interview/models"
)

func Test_Message_LogCreate(t *testing.T) {
	assert := assert.New(t)

	m := NewMessageLog(1)

	assert.NotNil(m)
}

func Test_MessageLog_Append(t *testing.T) {
	assert := assert.New(t)
	cap := 10
	iter := 15
	m := NewMessageLog(cap)
	for i := 0; i < iter; i++ {
		m.Push(models.Message{Type: i, Body: "nope"})
	}

	out := m.GetLog()
	fmt.Printf("%v", out)
	assert.Equal(cap, len(out))
	for i := 0; i < cap; i++ {
		assert.Equal((iter-1)-i, out[i].Type)
	}
}

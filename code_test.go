package eventbus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Foo struct {
	EventBus
}

func TestEventBus(t *testing.T) {
	assert := assert.New(t)
	foo := &Foo{}
	token := "banana"
	received := ""
	foo.On("quack", func(args ...any) {
		received = args[0].(string)
	})
	foo.Trigger("quack", token)
	assert.Equal(received, token)
}

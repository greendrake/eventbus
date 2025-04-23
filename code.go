package eventbus

import (
	"sync"
)

type LF func(args ...any)

type EventBusInterface interface {
	On(n string, l LF)
	Trigger(n string, args ...any)
}

type EventBus struct {
	listeners sync.Map
}

func (eb *EventBus) On(n string, l LF) {
	ls, ok := eb.listeners.Load(n)
	if ok {
		ls = append(ls.([]LF), l)
	} else {
		ls = []LF{l}
	}
	eb.listeners.Store(n, ls)
}

func (eb *EventBus) Trigger(n string, args ...any) {
	ls, ok := eb.listeners.Load(n)
	if ok {
		for _, listener := range ls.([]LF) {
			listener(args...)
		}
	}
}

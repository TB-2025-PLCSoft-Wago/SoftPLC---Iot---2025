package bus

import (
	"sync"
)

type EventBus struct {
	subscribers []chan string
	mu          sync.Mutex
}

func NewEventBus() *EventBus {
	return &EventBus{}
}

func (e *EventBus) Subscribe() <-chan string {
	e.mu.Lock()
	defer e.mu.Unlock()
	ch := make(chan string, 1)
	e.subscribers = append(e.subscribers, ch)
	return ch
}

func (e *EventBus) Publish(msg string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	for _, ch := range e.subscribers {
		select {
		case ch <- msg:
		default:
			// Ne bloque pas si le canal est plein
		}
	}
}

func (b *EventBus) Reset() {
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, ch := range b.subscribers {
		close(ch)
	}
	b.subscribers = nil
}

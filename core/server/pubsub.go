package server

import (
	"github.com/alabianca/codernames/core/mongo/models"
	"sync"
)

type PubSub struct {
	mu sync.RWMutex
	topics map[string][]chan models.Game
	closed bool
}

func NewPubSub() *PubSub {
	return &PubSub{
		topics: make(map[string][]chan models.Game),
	}
}

func (p *PubSub) Publish(topic string, game models.Game) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.closed {
		return
	}

	for _, ch := range p.topics[topic] {
		ch <- game
	}
}

func (p *PubSub) Subscribe(topic string) <-chan models.Game {
	p.mu.Lock()
	defer p.mu.Unlock()

	out := make(chan models.Game, 1) // important to buffer the channel
	p.topics[topic] = append(p.topics[topic], out)

	return out
}

func (p *PubSub) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()

	if !p.closed {
		p.closed = true
		for _, subs := range p.topics {
			for _, ch := range subs {
				close(ch)
			}
		}
	}
}

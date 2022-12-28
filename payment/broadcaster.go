package payment

import "sync"

type Broadcaster struct {
	clients map[chan<- Payment]struct{}
	mu      sync.Mutex
}

func NewBroadcaster() *Broadcaster {
	return &Broadcaster{
		clients: make(map[chan<- Payment]struct{}),
	}
}

func (b *Broadcaster) Add(c chan<- Payment) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.clients[c] = struct{}{}
}

func (b *Broadcaster) Remove(c chan<- Payment) {
	b.mu.Lock()
	defer b.mu.Unlock()

	delete(b.clients, c)
}

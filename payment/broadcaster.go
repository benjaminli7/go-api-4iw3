package payment

type Broadcaster struct {
	payments    chan Payment
	subscribe   chan chan Payment
	unsubscribe chan chan Payment
}

func NewBroadcaster() *Broadcaster {
	return &Broadcaster{
		payments:    make(chan Payment),
		subscribe:   make(chan chan Payment),
		unsubscribe: make(chan chan Payment),
	}
}

func (b *Broadcaster) Subscribe() <-chan Payment {
	ch := make(chan Payment)
	b.subscribe <- ch
	return ch
}

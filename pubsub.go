package pubsub

// Pubsub ...
type Pubsub struct {
	channels map[string][]chan string
}

// New ...
func New() *Pubsub {
	return &Pubsub{
		channels: make(map[string][]chan string),
	}
}

// Subscribe ...
func (p *Pubsub) Subscribe(ch string) chan string {
	newch := make(chan string, 1)
	p.channels[ch] = append(p.channels[ch], newch)
	return newch
}

// Publish ...
func (p *Pubsub) Publish(ch, data string) {
	for _, c := range p.channels[ch] {
		c <- data
	}
}

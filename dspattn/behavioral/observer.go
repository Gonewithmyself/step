package dspattn

type broker struct {
	subers map[int]suber
}

type suber interface {
	onMsg(event int)
}

func (b *broker) register(event int, s suber) {
	b.subers[event] = s
}

func (b *broker) post(event int) {
	sub, ok := b.subers[event]
	if ok {
		sub.onMsg(event)
	}
}

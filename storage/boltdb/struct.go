package main

type config struct {
	id   int32
	size int32
}

func assign() {
	a := &config{id: 1}

	b := &config{id: 2}

	_ = a.id

	b = a

	_ = b
}

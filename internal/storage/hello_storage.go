package storage

import "math/rand"

type Hello interface {
	Random() string
}

type helloStorage struct {
	list []string
}

func NewHelloStorage() Hello {
	return &helloStorage{
		list: []string{"World", "Mundo", "Welt", "Mondo", "Monde", "世界"},
	}
}

func (hs *helloStorage) Random() string {
	return hs.list[rand.Intn(len(hs.list))]
}

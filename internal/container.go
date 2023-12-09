package internal

import "github.com/edmarfelipe/go-template/internal/storage"

type Container struct {
	HelloStorage storage.Hello
}

func (ct *Container) Close() error {
	return nil
}

func NewContainer() (Container, error) {
	return Container{
		HelloStorage: storage.NewHelloStorage(),
	}, nil
}

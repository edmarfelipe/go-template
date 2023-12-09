package handler

import (
	"net/http"

	"github.com/edmarfelipe/go-template/internal"
)

type HelloHandler struct {
	ct internal.Container
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello " + h.ct.HelloStorage.Random()))
}

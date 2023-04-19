package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Handler struct {
	l       *log.Logger
	message string
}

func NewHandler(l *log.Logger, message string) *Handler {
	return &Handler{l, message}
}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")
	_, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Hello %s\n", h.message)
}

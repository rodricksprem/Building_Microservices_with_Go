package handlers

import (
	"log"
	"net/http"
)

type GoodByeHandler struct {
	l *log.Logger
}

func NewGoodByeHandler(l *log.Logger) *GoodByeHandler {
	return &GoodByeHandler{l}
}

func (n *GoodByeHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Bye...."))
}

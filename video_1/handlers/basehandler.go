package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type BaseHandler struct {
	l *log.Logger
}

func NewBaseHandler(l *log.Logger) *BaseHandler {
	return &BaseHandler{l}
}

func (B *BaseHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Print("Welcome to Microservice ")
	d, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Oops Error occured", 400)
	}
	fmt.Fprintf(res, "Hello %s", d)

}

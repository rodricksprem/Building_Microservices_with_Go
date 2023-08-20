package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"example.com/video1/data"
)

type ProductHandler struct {
	l *log.Logger
}

func NewProductHandler(l *log.Logger) *ProductHandler {
	return &ProductHandler{l}
}

func (ph *ProductHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	products := data.GetProducts()
	data, err := json.Marshal(products)
	if err != nil {
		http.Error(res, "Failed to parse data ", http.StatusInternalServerError)
	}
	res.Write(data)
}

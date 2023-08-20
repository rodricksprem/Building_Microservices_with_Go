package handlers

import (
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
	if req.Method == http.MethodGet {
		ph.GetProducts(res, req)
		return
	}
	if req.Method == http.MethodPost {
		ph.l.Println("Handle POST Method")
		ph.AddProducts(res, req)
		return
	}
	res.WriteHeader(http.StatusMethodNotAllowed)
}

func (ph *ProductHandler) GetProducts(res http.ResponseWriter, req *http.Request) {
	ps := data.GetProducts()
	err := ps.ToJSON(res)
	if err != nil {
		http.Error(res, "Failed to parse data ", http.StatusInternalServerError)
	}
}

func (ph *ProductHandler) AddProducts(res http.ResponseWriter, req *http.Request) {
	ph.l.Println("Handle POST Product")
	prod := &data.Product{}
	err := prod.FromJSON(req.Body)
	if err != nil {
		http.Error(res, "Failed to unmarshal data ", http.StatusBadRequest)
	}
	data.AddProducts(prod)
	ph.l.Printf("Prod %#v", prod)
}

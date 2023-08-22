package handlers

import (
	"log"
	"net/http"
	"strconv"

	"example.com/video1/data"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	l *log.Logger
}

func NewProductHandler(l *log.Logger) *ProductHandler {
	return &ProductHandler{l}
}

func (ph *ProductHandler) GetProducts(res http.ResponseWriter, req *http.Request) {
	ps := data.GetProducts()
	err := ps.ToJSON(res)
	if err != nil {
		http.Error(res, "Failed to parse data ", http.StatusInternalServerError)
	}
}

func (ph *ProductHandler) AddProduct(res http.ResponseWriter, req *http.Request) {
	ph.l.Println("Handle POST Product")
	prod := &data.Product{}
	err := prod.FromJSON(req.Body)
	if err != nil {
		http.Error(res, "Failed to unmarshal data ", http.StatusBadRequest)
	}
	data.AddProducts(prod)
	ph.l.Printf("Prod %#v", prod)
}

func (ph *ProductHandler) UpdateProducts(res http.ResponseWriter, req *http.Request) {
	ph.l.Println("Handle Update Product")
	vars := mux.Vars(req)
	idstr := vars["id"]
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(res, "Failed to convert id into int data ", http.StatusBadRequest)
	}
	ph.l.Println(" got id ", id)
	prod := &data.Product{}
	errnew := prod.FromJSON(req.Body)
	if errnew != nil {
		http.Error(res, "Failed to unmarshal data ", http.StatusBadRequest)
	}
	errnew1 := data.UpdateProducts(id, prod)
	if errnew1 != nil {
		http.Error(res, "unable to update product", http.StatusNotModified)
	}

	ph.l.Printf("Prod %#v", prod)
}

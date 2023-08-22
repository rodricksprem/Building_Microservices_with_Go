package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

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
		ph.l.Println("Handle GET Method")
		ph.GetProducts(res, req)
		return
	}
	if req.Method == http.MethodPost {
		ph.l.Println("Handle POST Method")
		ph.AddProducts(res, req)
		return
	}
	if req.Method == http.MethodPut {
		ph.l.Println("Handle PUT Method")
		ph.UpdateProducts(res, req)
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

func (ph *ProductHandler) UpdateProducts(res http.ResponseWriter, req *http.Request) {
	ph.l.Println("Handle Update Product")
	compiler := regexp.MustCompile("/([0-9]*)")
	g := compiler.FindAllStringSubmatch(req.URL.Path, -1)

	if len(g) == 1 {
		fmt.Println(len(g[0]))
		if len(g[0]) == 2 {
			idstring := g[0][1]
			id, err := strconv.Atoi(idstring)
			if err != nil {
				http.Error(res, "Invalid URI", http.StatusBadRequest)
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
	}

}

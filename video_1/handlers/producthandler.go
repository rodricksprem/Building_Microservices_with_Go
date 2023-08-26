package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"example.com/video1/data"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	l *log.Logger
}

type KeyProduct struct{}

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
	prod := req.Context().Value(KeyProduct{}).(data.Product)
	data.AddProducts(&prod)
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
	prod := req.Context().Value(KeyProduct{}).(data.Product)
	errval := prod.Validate()
	if errval != nil {
		ph.l.Println("[Error] validating product ", err)
		http.Error(res, fmt.Sprintf("Data Validation Failed %s", err), http.StatusBadRequest)
	}
	errnew1 := data.UpdateProducts(id, &prod)
	if errnew1 != nil {
		http.Error(res, "unable to update product", http.StatusNotModified)
	}

	ph.l.Printf("Prod %#v", prod)
}
func (ph *ProductHandler) MiddleWareForPayloadValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		ph.l.Println("Handle the middleware ")

		prod := data.Product{}
		errnew := prod.FromJSON(req.Body)
		if errnew != nil {

			http.Error(res, "Failed to unmarshal data ", http.StatusBadRequest)
		}
		err := prod.Validate()
		if err != nil {
			ph.l.Println("[Error] validating product ", err)
			http.Error(res, fmt.Sprintf("Data Validation Failed %s", err), http.StatusBadRequest)
		}

		ctx := context.WithValue(req.Context(), KeyProduct{}, prod)

		r := req.WithContext(ctx)
		next.ServeHTTP(res, r)

	})

}

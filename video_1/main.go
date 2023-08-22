// https://www.youtube.com/watch?v=VzBGi_n65iU&list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"example.com/video1/handlers"
	"github.com/gorilla/mux"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	servermux := mux.NewRouter()
	getProductsRouter := servermux.Methods(http.MethodGet).Subrouter()
	addProductsRouter := servermux.Methods(http.MethodPost).Subrouter()
	putProductRouter := servermux.Methods(http.MethodPut).Subrouter()
	productHandler := handlers.NewProductHandler(l)
	getProductsRouter.HandleFunc("/getproducts", productHandler.GetProducts)
	addProductsRouter.HandleFunc("/addproduct", productHandler.AddProduct)
	putProductRouter.HandleFunc("/updateproduct/{id:[0-9]+}", productHandler.UpdateProducts)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      servermux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	signalchannel := make(chan os.Signal)
	signal.Notify(signalchannel, os.Interrupt)
	signal.Notify(signalchannel, os.Kill)
	sig := <-signalchannel
	l.Println(sig, " signal received and gracefully shuting down the server ... ")
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second) // wait for graceful shutdown
	s.Shutdown(tc)

}

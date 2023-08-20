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
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	servermux := http.NewServeMux()

	baseHandler := handlers.NewBaseHandler(l)
	goodbyeHandler := handlers.NewGoodByeHandler(l)
	servermux.Handle("/", baseHandler)
	servermux.Handle("/goodbye", goodbyeHandler)

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

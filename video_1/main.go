// https://www.youtube.com/watch?v=VzBGi_n65iU&list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_
package main

import (
	"log"
	"net/http"
	"os"

	"example.com/video1/handlers"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	servermux := http.NewServeMux()

	baseHandler := handlers.NewBaseHandler(l)
	goodbyeHandler := handlers.NewGoodByeHandler(l)
	servermux.Handle("/", baseHandler)
	servermux.Handle("/goodbye", goodbyeHandler)

	http.ListenAndServe(":9000", servermux)
}

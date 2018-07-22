package main

import (
	"flag"
	"github.com/hmoniaga/go-friend/pkg/user"
	"log"
	"net/http"
)

func main() {
	var addr string
	flag.StringVar(&addr, "addr", ":8080", "binding address to listen to")
	flag.Parse()

	log.Println("starting go-friend")
	http.ListenAndServe(addr, user.NewHTTPHandler())
}

package main

import (
	"context"
	"log"
	"net/http"
)

func g() {
	log.Println("g(): context has been canceled")
}

func serveHTTP(_ http.ResponseWriter, r *http.Request) {
	log.Println("[BEGIN] Serving HTTP")
	defer log.Println("[DONE] Serving HTTP")
	context.AfterFunc(r.Context(), g)
}

func main() {
	http.Handle("/", http.HandlerFunc(serveHTTP))
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatalln(err)
	}
}

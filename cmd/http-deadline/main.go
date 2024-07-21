package main

import (
	"log"
	"net/http"
)

func serveHTTP(_ http.ResponseWriter, r *http.Request) {
	log.Println("[BEGIN] Serving HTTP")
	defer log.Println("[DONE] Serving HTTP")
	deadline, ok := r.Context().Deadline()
	log.Println("deadline:", deadline, "ok:", ok)
}

func main() {
	http.Handle("/", http.HandlerFunc(serveHTTP))
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatalln(err)
	}
}

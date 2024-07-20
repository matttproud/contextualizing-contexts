package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func f(ctx context.Context) {
	log.Println("[BEGIN] f")
	defer log.Println("[END] f")
	// Flimsily make it improbable for this function to continue while the
	// handler is serving.
	time.Sleep(time.Second)
	select {
	case <-time.After(5 * time.Second):
		log.Println("5s")
	case <-ctx.Done():
		log.Println("canceled")
	}
}

func serveHTTP(_ http.ResponseWriter, r *http.Request) {
	log.Println("[BEGIN] Serving HTTP")
	defer log.Println("[DONE] Serving HTTP")
	go f(r.Context())
}

func main() {
	http.Handle("/", http.HandlerFunc(serveHTTP))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}

}

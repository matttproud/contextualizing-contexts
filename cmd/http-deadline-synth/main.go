package main

import (
	"context"
	"log"
	"net/http"
)

func wrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if time, err := http.ParseTime(r.Header.Get("X-MTP-Deadline")); err == nil {
			ctx, cancel := context.WithDeadline(r.Context(), time)
			defer cancel()
			r = r.WithContext(ctx)
		}
		h.ServeHTTP(w, r)
	})
}

func serveHTTP(_ http.ResponseWriter, r *http.Request) {
	log.Println("[BEGIN] Serving HTTP")
	defer log.Println("[DONE] Serving HTTP")
	deadline, ok := r.Context().Deadline()
	log.Println("deadline:", deadline, "ok:", ok)
}

func main() {
	http.Handle("/", wrap(http.HandlerFunc(serveHTTP)))
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatalln(err)
	}
}

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

func g() {
	log.Println("g(): context has been canceled")
}

func serveHTTP(_ http.ResponseWriter, r *http.Request) {
	log.Println("[BEGIN] Serving HTTP")
	defer log.Println("[DONE] Serving HTTP")
	ctx := r.Context()
	{
		// See http-branch.
		go f(ctx)
	}
	{
		// See http-notify.
		context.AfterFunc(ctx, g)
	}
	{
		// See http-deadline.
		deadline, ok := ctx.Deadline()
		log.Println("deadline:", deadline, "ok:", ok)
	}
}

// See http-deadline-synth.
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

func main() {
	http.Handle("/", wrap(http.HandlerFunc(serveHTTP)))
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatalln(err)
	}
}

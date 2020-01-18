package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const requestIDKey = 42

func main() {
	http.HandleFunc("/", Decorate(handler))
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	Println(ctx, "handler started")
	defer Println(ctx, "handler ended")

	select {
	case <-time.After(time.Second * 5):
		fmt.Fprintf(w, "Hello")
	case <-ctx.Done():
		err := ctx.Err()
		Println(ctx, "handler ended")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Println(ctx context.Context, msg string) {
	id, ok := ctx.Value(requestIDKey).(int64)

	if !ok {
		log.Println("could not find request ID in context")
		return
	}
	log.Printf("[%d] %s", id, msg)
}

//Decorate middleware to logging request ID
func Decorate(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := rand.Int63()
		ctx = context.WithValue(ctx, requestIDKey, id)
		f(w, r.WithContext(ctx))
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Handler struct {
	Name string
}

func (ptr *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Name:", ptr.Name, "URL:", r.URL.String())

}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "url:", r.URL.String())
}

func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:         ":4001",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	mux.HandleFunc("/", handler2)

	log.Println("starting server in http://localhost:4000")
	server.ListenAndServe()

}

package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("RequestID", "d41d8cd98f00b204")
	w.Header().Add("RequestID", "dkurmant")
	w.Header().Del("RequestID")
	data := w.Header().Get("RequestID")
	fmt.Fprintln(w, data)

}
func main() {

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.HandleFunc("/", handler)
	http.HandleFunc("/page",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "single page:", r.URL.String())
		})
	http.HandleFunc("/pages/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Multiple pages:", r.URL.String())

		})

	log.Println("starting server at: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}

}

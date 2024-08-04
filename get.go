package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler3)
	log.Println("server starting in http://localhost:3000")
	http.ListenAndServe(":3000", nil)

}

func handler3(w http.ResponseWriter, r *http.Request) {
	myParam := r.URL.Query().Get("param")
	if myParam != "" {
		fmt.Fprintln(w, "myparam is", myParam)
	}
	key := r.FormValue("key")
	if key != "" {
		fmt.Fprintln(w, "key is:", key)

	}
}

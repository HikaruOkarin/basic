package main

import (
	"fmt"
	"log"
	"net/http"
)

var loginFormTmpl = []byte(
	`
	<html>
	<body>
	<form action="/" method="post">
	Login: <input type="text" name="login">
	Password: <input type="password" name="password">
	<input type="submit" value="Login">
	</form>
	</body>
	</html>
	`)

func MainPage(w http.ResponseWriter, r *http.Request) {
	if http.MethodPost == r.Method {
		data := r.FormValue("login")
		fmt.Fprintln(w, "login:", data)
	} else {
		w.Write(loginFormTmpl)
	}
}
func main() {
	http.HandleFunc("/", MainPage)
	log.Println("http://localhost:2000")
	http.ListenAndServe(":2000", nil)
}

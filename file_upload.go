package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
)

var html = []byte(`
<html>
<body>
<form action="/upload" method="post"
enctype="multipart/form-data">
Image: <input type="file" name="my_file">
<input type="submit" value="upload">

</form>

</body>
</html>
`)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(html)

}

func Controller(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(5 * 1024 * 1025)
	file, handler, err := r.FormFile("my_file")
	if err != nil {
		return
	}
	defer file.Close()
	fmt.Fprintln(w, handler.Filename)
	fmt.Fprintln(w, handler.Header)
	hasher := md5.New()
	io.Copy(hasher, file)
	fmt.Fprintln(w, string(hasher.Sum(nil)))
}
func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/upload", Controller)
	log.Println("http://localhost:1000")
	http.ListenAndServe(":1000", nil)

}

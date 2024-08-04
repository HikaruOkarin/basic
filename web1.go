package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	
)

func main() {
	SendForm()
}

func MakeGetRequest() {
	client := http.Client{}

	request, err := http.NewRequest("GET", "httpbin.org/get", nil)
	// resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close() // first in last out
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Println(result)
}

func MakePostRequest() {
	message := map[string]interface{}{
		"hello": "world",
		"life":  42,
		"embedded": map[string]string{
			"yes": "of course!",
		},
	}
	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}
	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)
	log.Println(res)
	log.Println(res["data"])
}

func SendForm() {
	formData := url.Values{ // map[string] string
		"name": {"masnun"},
	}
	resp, err := http.PostForm("https://httpbin.org/post", formData)
	if err != nil {
		log.Fatalln(err)
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Println(result)
}


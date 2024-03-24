package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/pflag"
)

func getParams() (url, path, queryType string) {

	// Когда указывают с именами: go run main.go -u=localhost:10001 -p=/path
	pflag.StringVarP(&url, "url", "u", "localhost:8080", "server url")
	pflag.StringVarP(&path, "path", "p", "/hello", "resource path")
	pflag.StringVarP(&queryType, "queryType", "q", "GET", "get or post query")

	// Когда указывают: go run main.go -u -p - подставить значения по умолчанию
	pflag.Lookup("url").NoOptDefVal = "localhost:8080"
	pflag.Lookup("path").NoOptDefVal = "/hello"
	pflag.Lookup("queryType").NoOptDefVal = "GET"

	pflag.Parse()

	// Когда указывают: go run main.go - подставляем переменные окружения
	// -o будет равен значению из переменной окружения если его вообще никак не укажут
	if url == "" && path == "" && queryType == "" {
		log.Fatal("Ошибка")
	}

	return
}

func main() {
	url, path, queryType := getParams()

	resource := "http://"+url+path

	if queryType == "GET" {
		req, err := http.Get(resource)
		if err != nil {
			log.Println(err)
		}
	
		defer req.Body.Close()
	
		bodyBytes, err := io.ReadAll(req.Body)
		fmt.Println(string(bodyBytes))
	} else if queryType == "POST"{
		data := []byte(`{"foo":"bar"}`)
		r := bytes.NewReader(data)
		req, err := http.Post(resource, "application/json", r)
		if err != nil {
			log.Println(err)
		}
	
		defer req.Body.Close()
	
		bodyBytes, err := io.ReadAll(req.Body)
		fmt.Println(string(bodyBytes))
	}
}

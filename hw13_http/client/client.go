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
	pflag.StringVarP(&url, "url", "u", "http://localhost:8080", "server url")
	pflag.StringVarP(&path, "path", "p", "/", "resource path")
	pflag.StringVarP(&queryType, "queryType", "q", "GET", "get or post query")

	// Когда указывают: go run main.go -u -p - подставить значения по умолчанию
	pflag.Lookup("url").NoOptDefVal = "http://localhost:8080"
	pflag.Lookup("path").NoOptDefVal = "/"
	pflag.Lookup("queryType").NoOptDefVal = "GET"

	pflag.Parse()

	// Когда указывают: go run main.go - подставляем переменные окружения
	// -o будет равен значению из переменной окружения если его вообще никак не укажут
	if url == "" && path == "" && queryType == "" {
		log.Fatal("Ошибка")
	}

	return
}

func get(url, path string) (body string, err error) {
	resource := fmt.Sprintf("%s%s", url, path)
	req, err := http.Get(resource) //nolint
	if err != nil {
		log.Println(err)
	}

	defer req.Body.Close()

	bodyBytes, err := io.ReadAll(req.Body)
	return string(bodyBytes), err
}

func post(url, path string) (body string, err error) {
	resource := fmt.Sprintf("%s%s", url, path)
	data := []byte(`{"foo":"bar"}`)
	r := bytes.NewReader(data)
	req, err := http.Post(resource, "application/json", r) //nolint
	if err != nil {
		log.Println(err)
	}

	defer req.Body.Close()

	bodyBytes, err := io.ReadAll(req.Body)
	return string(bodyBytes), err
}

func main() {
	url, path, queryType := getParams()

	if queryType == "GET" {
		body, err := get(url, path)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(body)
	} else if queryType == "POST" {
		body, err := post(url, path)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(body)
	}
}

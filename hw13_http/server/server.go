package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/pflag"
)

type Person struct {
	Name    string
	Surname string
}

func getParams() (address string, port int16) {
	// Когда указывают с именами: go run main.go -u=localhost:10001 -p=/path
	pflag.StringVarP(&address, "address", "u", "localhost", "server address")
	pflag.Int16VarP(&port, "port", "p", 8080, "service port")

	pflag.Parse()

	// Когда указывают: go run main.go - подставляем переменные окружения
	// -o будет равен значению из переменной окружения если его вообще никак не укажут
	if address == "" && port == 0 {
		log.Fatal("Ошибка")
	}

	return
}

func hello(w http.ResponseWriter, r *http.Request) {
	answer := "Hello!"

	switch r.Method {
	case "GET":
		nameQueryParam := r.URL.Query().Get("name")
		surnameQueryParam := r.URL.Query().Get("surname")
		log.Println("GET " + answer + " " + nameQueryParam + " " + surnameQueryParam)
		fmt.Fprintf(w, "GET "+answer+" "+nameQueryParam+" "+surnameQueryParam)
	case "POST":
		var p Person

		_ = json.NewDecoder(r.Body).Decode(&p)

		log.Println("POST " + answer + " " + p.Name + " " + p.Surname)
		fmt.Fprintf(w, "POST "+answer+" "+p.Name+" "+p.Surname)
	default:
		// fmt.Fprintf(w, "Only GET and POST methods are supported.")
		code := 405
		http.Error(w, "Only GET and POST methods are supported.", code)
	}
}

func main() {
	address, port := getParams()
	host := fmt.Sprintf("%s:%d", address, port)

	http.HandleFunc("/", hello)
	http.ListenAndServe(host, nil) //nolint
}

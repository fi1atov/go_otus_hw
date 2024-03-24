package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/pflag"
)

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
		log.Println("GET"+answer)
		fmt.Fprintf(w, "GET "+answer)
	case "POST":
		log.Println("POST"+answer)
		fmt.Fprintf(w, "POST "+answer)
	default:
		fmt.Fprintf(w, "Only GET and POST methods are supported.")
	}
}

func main() {
	address, port := getParams()
	host := fmt.Sprintf("%s:%d", address, port)

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(host, nil)
}

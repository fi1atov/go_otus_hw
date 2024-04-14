package main

import (
	"fmt"
	"log"
	_ "net/http"

	"github.com/fi1atov/go_otus_hw/hw15_go_sql/postgres"
	"github.com/fi1atov/go_otus_hw/hw15_go_sql/server"
	"github.com/spf13/pflag"
)

func getServerParams() (address string, port int16) {
	// Когда указывают с именами: go run main.go -u=localhost -p=8080
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

func main() {
	address, port := getServerParams()
	host := fmt.Sprintf("%s:%d", address, port)
	dbPool, ctx, err := postgres.OpenPool()
	if err != nil {
		log.Fatalf("cannot open database pool: %v", err)
	}
	srv := server.NewServer(ctx, dbPool)
	srv.Run(host)
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/pflag"
)

type Person struct {
	Name    string
	Surname string
}

func getDatabaseConn() (ctx context.Context, conn *pgxpool.Pool) {
	ctx = context.Background()
	dsn := "postgres://postgres:postgres@localhost:5432/test_db?search_path=test_schema&sslmode=disable&pool_max_conns=20"

	pgCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal("We couldn't find any correct DSN")
	}

	// conn, err := pgxpool.New(ctx, dsn)
	conn, err = pgxpool.NewWithConfig(ctx, pgCfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	if err = conn.Ping(ctx); err != nil {
		log.Fatal("We cannot connect to database")
	}
	return
}

func getParams() (address string, port int16) {
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

func getProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Вывод всех продуктов
		ctx, conn := getDatabaseConn()
		defer conn.Close()
		products, err := GetProducts(ctx, conn)
		if err != nil {
			log.Fatal(err)
		}
		b, err := json.Marshal(products)
		if err != nil {
			fmt.Println(err)
			return
		}
		result := string(b)
		log.Println(result)
		fmt.Fprintf(w, result)
	default:
		code := 405
		http.Error(w, "Only GET method supported.", code)
	}
}

func main() {
	address, port := getParams()
	host := fmt.Sprintf("%s:%d", address, port)

	http.HandleFunc("/get_products", getProducts)
	http.ListenAndServe(host, nil) //nolint
}

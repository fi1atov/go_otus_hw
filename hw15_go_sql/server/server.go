package server

import (
	"log"
	"net/http"
	"time"

	"github.com/fi1atov/go_otus_hw/hw15_go_sql/postgres"
	"github.com/fi1atov/go_otus_hw/hw15_go_sql/structs"
)

type Server struct {
	server         *http.Server
	productService structs.ProductService
	// ...
}

func NewServer(dbpool *postgres.DBPool) *Server {
	s := Server{
		server: &http.Server{
			WriteTimeout: 5 * time.Second,
			ReadTimeout:  5 * time.Second,
			IdleTimeout:  5 * time.Second,
		},
	}

	http.HandleFunc("/get_products", s.getProducts)
	// ...

	s.productService = postgres.NewProductService(dbpool)
	// ...

	return &s
}

func (s *Server) Run(host string) error {
	s.server.Addr = host
	log.Printf("server starting on address: %s", host)
	return s.server.ListenAndServe()
}

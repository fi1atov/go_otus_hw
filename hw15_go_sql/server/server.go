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
	userService    structs.UserService
	orderService   structs.OrderService
	// ...
}

func NewServer(db *postgres.DB) *Server {
	s := Server{
		server: &http.Server{
			WriteTimeout: 5 * time.Second,
			ReadTimeout:  5 * time.Second,
			IdleTimeout:  5 * time.Second,
		},
	}

	http.HandleFunc("GET /products", s.getProducts)
	http.HandleFunc("POST /product", s.createProduct)
	http.HandleFunc("PUT /product/{id}", s.updateProduct)
	http.HandleFunc("DELETE /product/{id}", s.deleteProduct)

	http.HandleFunc("GET /users", s.getUsers)
	http.HandleFunc("POST /user", s.createUser)
	http.HandleFunc("PUT /user/{id}", s.updateUser)
	http.HandleFunc("DELETE /user/{id}", s.deleteUser)
	http.HandleFunc("GET /user_stat/{id}", s.getUserStat)

	http.HandleFunc("POST /order", s.createOrder)
	http.HandleFunc("DELETE /order/{id}", s.deleteOrder)
	http.HandleFunc("GET /order/{id}", s.getOrdersByUser)
	// ...

	s.productService = postgres.NewProductService(db)
	s.userService = postgres.NewUserService(db)
	s.orderService = postgres.NewOrderService(db)
	// ...

	return &s
}

func (s *Server) Run(host string) error {
	s.server.Addr = host
	log.Printf("server starting on address: %s", host)
	return s.server.ListenAndServe()
}

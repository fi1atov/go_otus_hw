package server

import (
	"log"
	"net/http"
)

func (s *Server) getProducts(w http.ResponseWriter, r *http.Request) {
	// Вывод всех продуктов
	products, err := s.productService.GetProducts()
	if err != nil {
		log.Fatal(err)
	}

	writeJSON(w, http.StatusOK, products)
}

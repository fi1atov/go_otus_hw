package server

import (
	"log"
	"net/http"
)

func (s *Server) getProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Вывод всех продуктов
		products, err := s.productService.GetProducts()
		if err != nil {
			log.Fatal(err)
		}

		writeJSON(w, http.StatusOK, products)
	default:
		code := 405
		http.Error(w, "error", code)
	}
}

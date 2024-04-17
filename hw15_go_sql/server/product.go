package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fi1atov/go_otus_hw/hw15_go_sql/structs"
)

func (s *Server) getProducts(w http.ResponseWriter, r *http.Request) {
	// Вывод всех продуктов
	products, err := s.productService.GetProducts()
	if err != nil {
		log.Fatal(err)
	}

	writeJSON(w, http.StatusOK, products)
}

func (s *Server) createProduct(w http.ResponseWriter, r *http.Request) {
	// Получение json-данных в структуру
	var product structs.ProductPatch
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Создание продукта
	if err = s.productService.CreateProduct(&product); err != nil {
		serverError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, M{"product": product})
}

// func (s *Server) updateProduct(w http.ResponseWriter, r *http.Request) {
// 	// Обновление продукта
// 	products, err := s.productService.UpdateProduct()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	writeJSON(w, http.StatusOK, products)
// }

// func (s *Server) deleteProduct(w http.ResponseWriter, r *http.Request) {
// 	// Удаление продукта
// 	product_id := r.PathValue("id")
// 	products, err := s.productService.DeleteProduct(uint(product_id))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	writeJSON(w, http.StatusOK, products)
// }

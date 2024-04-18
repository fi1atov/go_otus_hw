package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/fi1atov/go_otus_hw/hw15_go_sql/structs"
)

func (s *Server) getProducts(w http.ResponseWriter, _ *http.Request) {
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

func (s *Server) updateProduct(w http.ResponseWriter, r *http.Request) {
	var product structs.Product
	// Получение json-данных в структуру
	var productpatch structs.ProductPatch
	err := json.NewDecoder(r.Body).Decode(&productpatch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Получаем ID продукта из URL и конвертируем в int
	productID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Fatal(err)
	}
	// Обновление продукта
	err = s.productService.UpdateProduct(productID, &product, productpatch)
	if err != nil {
		log.Fatal(err)
	}

	writeJSON(w, http.StatusAccepted, productpatch)
}

func (s *Server) deleteProduct(w http.ResponseWriter, r *http.Request) {
	// Получаем ID продукта из URL и конвертируем в int
	productID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Fatal(err)
	}
	// Удаление продукта
	err = s.productService.DeleteProduct(productID)
	if err != nil {
		log.Fatal(err)
	}

	writeJSON(w, http.StatusOK, M{"productID": productID})
}

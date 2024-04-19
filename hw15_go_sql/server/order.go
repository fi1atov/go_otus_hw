package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/fi1atov/go_otus_hw/hw15_go_sql/structs"
)

func (s *Server) createOrder(w http.ResponseWriter, r *http.Request) {
	// Получение json-данных в структуру
	var order structs.OrderPatch
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Создание продукта
	if err = s.orderService.CreateOrder(&order); err != nil {
		serverError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, M{"order": order})
}

func (s *Server) deleteOrder(w http.ResponseWriter, r *http.Request) {
	// Получаем ID продукта из URL и конвертируем в int
	orderID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Fatal(err)
	}
	// Удаление продукта
	err = s.orderService.DeleteOrder(orderID)
	if err != nil {
		log.Fatal(err)
	}

	writeJSON(w, http.StatusOK, M{"orderID": orderID})
}

func (s *Server) getOrdersByUser(w http.ResponseWriter, r *http.Request) {
	// Получаем ID пользователя из URL и конвертируем в int
	userID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Fatal(err)
	}
	// выборка заказов по пользователю
	products, err := s.orderService.GetOrdersByUser(userID)
	if err != nil {
		log.Fatal(err)
	}

	writeJSON(w, http.StatusOK, products)
}

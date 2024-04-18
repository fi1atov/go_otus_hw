package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/fi1atov/go_otus_hw/hw15_go_sql/structs"
)

func (s *Server) getUsers(w http.ResponseWriter, _ *http.Request) {
	// Вывод всех продуктов
	users, err := s.userService.GetUsers()
	if err != nil {
		log.Fatal(err)
	}

	writeJSON(w, http.StatusOK, users)
}

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	// Получение json-данных в структуру
	var user structs.UserPatch
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Создание продукта
	if err = s.userService.CreateUser(&user); err != nil {
		serverError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, M{"user": user})
}

func (s *Server) updateUser(w http.ResponseWriter, r *http.Request) {
	var user structs.User
	// Получение json-данных в структуру
	var userpatch structs.UserPatch
	err := json.NewDecoder(r.Body).Decode(&userpatch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Получаем ID продукта из URL и конвертируем в int
	userID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Fatal(err)
	}
	// Обновление продукта
	err = s.userService.UpdateUser(userID, &user, userpatch)
	if err != nil {
		log.Fatal(err)
	}

	writeJSON(w, http.StatusAccepted, userpatch)
}

func (s *Server) deleteUser(w http.ResponseWriter, r *http.Request) {
	// Получаем ID продукта из URL и конвертируем в int
	userID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Fatal(err)
	}
	// Удаление продукта
	err = s.userService.DeleteUser(userID)
	if err != nil {
		log.Fatal(err)
	}

	writeJSON(w, http.StatusOK, M{"userID": userID})
}

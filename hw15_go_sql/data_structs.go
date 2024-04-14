package main

import (
	"time"
)

// Структура для хранения данных о продукте.
type Product struct {
	ID   int
	Name string
}

// Структура для хранения данных о пользователе.
type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

// Структура для хранения данных о заказах.
type Order struct {
	ID          int
	UserID      int
	OrderDate   time.Time
	TotalAmount float32
}

// Структура для хранения статистики по пользователю.
type UserStat struct {
	UserName    string
	TotalAmount float32
	AvgPrice    float32
}
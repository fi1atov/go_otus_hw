package structs

import (
	"time"
)

// Структура для хранения данных о заказах.
type Order struct {
	ID          int
	UserID      int
	OrderDate   time.Time
	TotalAmount float32
}

type OrderPatch struct {
	UserID      *int     `json:"userId"`
	TotalAmount *float32 `json:"totalAmount"`
}

type OrderService interface {
	CreateOrder(*OrderPatch) error

	DeleteOrder(int) error

	GetOrdersByUser(int) ([]Order, error)
}

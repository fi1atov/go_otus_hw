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

// Структура продукта вложенная в заказ.
type ProductRequest struct {
	ID *int `json:"id"`
}

// Структура заказа с вложенными в заказ продуктами.
type CreateOrderRequest struct {
	UserID   *int              `json:"userId"`
	Products *[]ProductRequest `json:"products"`
}

type OrderPatch struct {
	UserID      *int     `json:"userId"`
	TotalAmount *float32 `json:"totalAmount"`
}

type OrderService interface {
	CreateOrder(*OrderPatch) error

	CreateOrderV2(*CreateOrderRequest) error

	DeleteOrder(int) error

	GetOrdersByUser(int) ([]Order, error)
}

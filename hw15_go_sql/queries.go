package main

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
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

// Получение списка всех продуктов.
func GetProducts(ctx context.Context, db *pgxpool.Pool) ([]Product, error) {
	rows, err := db.Query(ctx, `SELECT s.id, s.name FROM products s`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err = rows.Scan(&product.ID, &product.Name)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

// Получение списка всех пользователей.
func GetUsers(ctx context.Context, db *pgxpool.Pool) ([]User, error) {
	rows, err := db.Query(ctx, `SELECT u.id, u.name, u.email, u.password FROM users u`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// Получение списка всех пользователей.
func GetUserOrders(ctx context.Context, db *pgxpool.Pool) ([]Order, error) {
	rows, err := db.Query(ctx, `
		SELECT o.* FROM orders o 
		join users u ON o.user_id = u.id
		where u.name=$1`, "Дима",
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		err = rows.Scan(&order.ID, &order.UserID, &order.OrderDate, &order.TotalAmount)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

// Получение статистики по пользователю.
func GetUserStat(ctx context.Context, db *pgxpool.Pool) ([]UserStat, error) {
	rows, err := db.Query(ctx, `
		SELECT u.name, SUM(o.total_amount) AS total_amount, AVG(p.price) as avg_price
		from orders o
		join order_products op ON o.id = op.order_id
		join products p ON op.product_id = p.id
		join users u ON o.user_id = u.id
		WHERE u.name = $1
		group by u.name`,
		"Дима",
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var userStats []UserStat
	for rows.Next() {
		var userStat UserStat
		err = rows.Scan(&userStat.UserName, &userStat.TotalAmount, &userStat.AvgPrice)
		if err != nil {
			return nil, err
		}
		userStats = append(userStats, userStat)
	}

	return userStats, nil
}

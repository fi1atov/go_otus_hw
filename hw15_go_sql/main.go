package main

import (
	"context"
	"fmt"
	"log"
	"os"
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

// Структура для хранения данных о заказх.
type Order struct {
	ID          int
	UserId      int
	OrderDate   time.Time
	TotalAmount float32
}

func main() {
	ctx := context.Background()
	dsn := "postgres://postgres:postgres@localhost:5432/test_db?search_path=test_schema&sslmode=disable&pool_max_conns=20"

	pgCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal("We couldn't find any correct DSN")
	}

	// conn, err := pgxpool.New(ctx, dsn)
	conn, err := pgxpool.NewWithConfig(ctx, pgCfg)
	defer conn.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	if err := conn.Ping(ctx); err != nil {
		log.Fatal("We cannot connect to database")
	}

	// Вывод всех продуктов
	products, err := GetProducts(ctx, conn)
	if err != nil {
		log.Fatal(err)
	}
	for _, product := range products {
		fmt.Printf("ID: %d, Имя: %s\n", product.ID, product.Name)
	}

	// Вывод всех пользователей
	users, err := GetUsers(ctx, conn)
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		fmt.Printf("ID: %d, Имя: %s, Email: %s, Пароль: %s\n", user.ID, user.Name, user.Email, user.Password)
	}

	// Вывод всех заказов
	orders, err := GetUserOrders(ctx, conn)
	if err != nil {
		log.Fatal(err)
	}
	for _, order := range orders {
		fmt.Printf(
			"ID: %d, ID пользователя: %d, Дата заказа: %s, Сумма: %f\n",
			order.ID, order.UserId, order.OrderDate, order.TotalAmount,
		)
	}
}

// Получение списка всех продуктов.
func GetProducts(ctx context.Context, db *pgxpool.Pool) ([]Product, error) {
	rows, err := db.Query(ctx, `SELECT s.id, s.name FROM products s`)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

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

	defer rows.Close()

	if err != nil {
		return nil, err
	}

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

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var orders []Order
	for rows.Next() {
		var order Order
		err = rows.Scan(&order.ID, &order.UserId, &order.OrderDate, &order.TotalAmount)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

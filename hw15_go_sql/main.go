package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()
	dsn := "postgres://postgres:postgres@localhost:5432/test_db?search_path=test_schema&sslmode=disable&pool_max_conns=20"

	pgCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal("We couldn't find any correct DSN")
	}

	// conn, err := pgxpool.New(ctx, dsn)
	conn, err := pgxpool.NewWithConfig(ctx, pgCfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close()

	if err = conn.Ping(ctx); err != nil {
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
			order.ID, order.UserID, order.OrderDate, order.TotalAmount,
		)
	}

	// Вывод статистики по пользователю
	userStats, err := GetUserStat(ctx, conn)
	if err != nil {
		log.Fatal(err)
	}
	for _, userStat := range userStats {
		fmt.Printf(
			"Имя: %s, Сумма заказов: %f, Средняя цена товара: %f\n",
			userStat.UserName, userStat.TotalAmount, userStat.AvgPrice,
		)
	}
}

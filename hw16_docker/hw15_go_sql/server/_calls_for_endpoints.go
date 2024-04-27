package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/jackc/pgx/v5/pgxpool"
// )

// func calls_for_endpoints() {
// 	ctx := context.Background()
// 	dsn := "postgres://postgres:postgres@localhost:5432/test_db
// ?search_path=test_schema&sslmode=disable&pool_max_conns=20"

// 	pgCfg, err := pgxpool.ParseConfig(dsn)
// 	if err != nil {
// 		log.Fatal("We couldn't find any correct DSN")
// 	}

// 	// conn, err := pgxpool.New(ctx, dsn)
// 	conn, err := pgxpool.NewWithConfig(ctx, pgCfg)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
// 		os.Exit(1)
// 	}

// 	defer conn.Close()

// 	if err = conn.Ping(ctx); err != nil {
// 		log.Fatal("We cannot connect to database")
// 	}

// 	// Вывод всех пользователей
// 	users, err := GetUsers(ctx, conn)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for _, user := range users {
// 		fmt.Printf("ID: %d, Имя: %s, Email: %s, Пароль: %s\n", user.ID, user.Name, user.Email, user.Password)
// 	}

// 	// Вывод всех заказов
// 	orders, err := GetUserOrders(ctx, conn)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for _, order := range orders {
// 		fmt.Printf(
// 			"ID: %d, ID пользователя: %d, Дата заказа: %s, Сумма: %f\n",
// 			order.ID, order.UserID, order.OrderDate, order.TotalAmount,
// 		)
// 	}

// 	// Вывод статистики по пользователю
// 	userStats, err := GetUserStat(ctx, conn)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for _, userStat := range userStats {
// 		fmt.Printf(
// 			"Имя: %s, Сумма заказов: %f, Средняя цена товара: %f\n",
// 			userStat.UserName, userStat.TotalAmount, userStat.AvgPrice,
// 		)
// 	}

// 	// Добавление пользователя
// 	name := "Vladislav"
// 	email := "Vladislav@mail.ru"
// 	password := "Vladislav1234"
// 	err = AddUser(ctx, conn, name, email, password)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Пользователь добавлен успешно.")

// 	// Обновление пользователя
// 	id := 4
// 	name = "Борис"
// 	email = "boris@mail.ru"
// 	password = "dasf"
// 	err = UpdateUser(ctx, conn, id, name, email, password)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Пользователь обновлен успешно.")

// 	// Удаление пользователя
// 	id = 6
// 	err = DeleteUser(ctx, conn, id)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Пользователь удален успешно.")

// 	// Добавление продукта
// 	name = "Сок"
// 	price := float32(90.0)
// 	err = AddProduct(ctx, conn, name, price)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Продукт добавлен успешно.")

// 	// Обновление продукта
// 	id = 2
// 	name = "Йогурт"
// 	price = 64.0
// 	err = UpdateProduct(ctx, conn, id, name, price)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Продукт обновлен успешно.")

// 	// Удаление продукта
// 	id = 5
// 	err = DeleteProduct(ctx, conn, id)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Продукт удален успешно.")

// 	// Добавление заказа
// 	userID := 3
// 	totalAmount := float32(500.0)
// 	err = AddOrder(ctx, conn, userID, totalAmount)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Заказ добавлен успешно.")

// 	// Добавление продуктов заказа
// 	orderID := 2
// 	productID := 2
// 	err = AddOrderProduct(ctx, conn, orderID, productID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Продукты заказа добавлены успешно.")

// 	// Удаление заказа
// 	id = 1
// 	err = DeleteOrder(ctx, conn, id)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Заказ удален успешно.")
// }

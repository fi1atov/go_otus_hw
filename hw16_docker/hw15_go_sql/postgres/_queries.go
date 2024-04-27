package postgres

// import (
// 	"context"
// 	"time"

// 	"github.com/jackc/pgx/v5/pgxpool"
// )

// // Получение списка всех продуктов.
// func GetProducts(ctx context.Context, db *pgxpool.Pool) ([]Product, error) {
// 	rows, err := db.Query(ctx, `SELECT s.id, s.name FROM products s`)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer rows.Close()

// 	var products []Product
// 	for rows.Next() {
// 		var product Product
// 		err = rows.Scan(&product.ID, &product.Name)
// 		if err != nil {
// 			return nil, err
// 		}
// 		products = append(products, product)
// 	}

// 	return products, nil
// }

// // Получение списка всех пользователей.
// func GetUsers(ctx context.Context, db *pgxpool.Pool) ([]User, error) {
// 	rows, err := db.Query(ctx, `SELECT u.id, u.name, u.email, u.password FROM users u`)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer rows.Close()

// 	var users []User
// 	for rows.Next() {
// 		var user User
// 		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
// 		if err != nil {
// 			return nil, err
// 		}
// 		users = append(users, user)
// 	}

// 	return users, nil
// }

// // Получение списка всех пользователей.
// func GetUserOrders(ctx context.Context, db *pgxpool.Pool) ([]Order, error) {
// 	rows, err := db.Query(ctx, `
// 		SELECT o.* FROM orders o
// 		join users u ON o.user_id = u.id
// 		where u.name=$1`, "Дима",
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer rows.Close()

// 	var orders []Order
// 	for rows.Next() {
// 		var order Order
// 		err = rows.Scan(&order.ID, &order.UserID, &order.OrderDate, &order.TotalAmount)
// 		if err != nil {
// 			return nil, err
// 		}
// 		orders = append(orders, order)
// 	}

// 	return orders, nil
// }

// // Получение статистики по пользователю.
// func GetUserStat(ctx context.Context, db *pgxpool.Pool) ([]UserStat, error) {
// 	rows, err := db.Query(ctx, `
// 		SELECT u.name, SUM(o.total_amount) AS total_amount, AVG(p.price) as avg_price
// 		from orders o
// 		join order_products op ON o.id = op.order_id
// 		join products p ON op.product_id = p.id
// 		join users u ON o.user_id = u.id
// 		WHERE u.name = $1
// 		group by u.name`,
// 		"Дима",
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer rows.Close()

// 	var userStats []UserStat
// 	for rows.Next() {
// 		var userStat UserStat
// 		err = rows.Scan(&userStat.UserName, &userStat.TotalAmount, &userStat.AvgPrice)
// 		if err != nil {
// 			return nil, err
// 		}
// 		userStats = append(userStats, userStat)
// 	}

// 	return userStats, nil
// }

// // Добавление пользователя.
// func AddUser(ctx context.Context, db *pgxpool.Pool, name, email, password string) error {
// 	_, err := db.Exec(
// 		ctx,
// 		`INSERT INTO users (name, email, password)
// 		VALUES ($1,$2,$3)`, name, email, password,
// 	)
// 	return err
// }

// // Добавление пользователя.
// func UpdateUser(ctx context.Context, db *pgxpool.Pool, id int, name, email, password string) error {
// 	_, err := db.Exec(
// 		ctx,
// 		`UPDATE users
// 		SET
// 		name = $2,
// 		email = $3,
// 		password = $4
// 		WHERE id = $1`, id, name, email, password,
// 	)
// 	return err
// }

// // Удаление пользователя.
// func DeleteUser(ctx context.Context, db *pgxpool.Pool, id int) error {
// 	_, err := db.Exec(
// 		ctx,
// 		`DELETE FROM users
// 		WHERE id = $1`, id,
// 	)
// 	return err
// }

// // Добавление продукта.
// func AddProduct(ctx context.Context, db *pgxpool.Pool, name string, price float32) error {
// 	_, err := db.Exec(
// 		ctx,
// 		`INSERT INTO products (name, price)
// 		VALUES ($1,$2)`, name, price,
// 	)
// 	return err
// }

// // Добавление продукта.
// func UpdateProduct(ctx context.Context, db *pgxpool.Pool, id int, name string, price float32) error {
// 	_, err := db.Exec(
// 		ctx,
// 		`UPDATE products
// 		SET
// 		name = $2,
// 		price = $3
// 		WHERE id = $1`, id, name, price,
// 	)
// 	return err
// }

// // Удаление продукта.
// func DeleteProduct(ctx context.Context, db *pgxpool.Pool, id int) error {
// 	_, err := db.Exec(
// 		ctx,
// 		`DELETE FROM products
// 		WHERE id = $1`, id,
// 	)
// 	return err
// }

// // Добавление заказа.
// func AddOrder(ctx context.Context, db *pgxpool.Pool, userID int, totalAmount float32) error {
// 	_, err := db.Exec(
// 		ctx,
// 		`INSERT INTO orders (user_id, order_date, total_amount)
// 		VALUES ($1,$2,$3)`, userID, time.Now(), totalAmount,
// 	)
// 	return err
// }

// // Добавление продуктов заказа.
// func AddOrderProduct(ctx context.Context, db *pgxpool.Pool, orderID, productID int) error {
// 	_, err := db.Exec(
// 		ctx,
// 		`INSERT INTO order_products (order_id, product_id)
// 		VALUES ($1,$2)`, orderID, productID,
// 	)
// 	return err
// }

// // Удаление продукта.
// func DeleteOrder(ctx context.Context, db *pgxpool.Pool, id int) error {
// 	_, err := db.Exec(
// 		ctx,
// 		`DELETE FROM orders
// 		WHERE id = $1`, id,
// 	)
// 	return err
// }

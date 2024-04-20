package postgres

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/fi1atov/go_otus_hw/hw15_go_sql/structs"
)

type OrderService struct {
	db *DB
}

func NewOrderService(db *DB) *OrderService {
	return &OrderService{db}
}

// Создание заказа.
func (ps *OrderService) CreateOrder(order *structs.OrderPatch) error {
	tx, err := ps.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.Exec(
		"INSERT INTO orders (user_id, order_date, total_amount) VALUES ($1, $2, $3)",
		order.UserID, time.Now(), order.TotalAmount,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	// Заказ создан, фиксируем транзакцию
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Заказ успешно создан.")
	return nil
}

// Создание заказа с продуктами.
func (ps *OrderService) CreateOrderV2(order *structs.CreateOrderRequest) error {
	var lastInsertID int
	var totalAmount int
	var productPrice int

	tx, err := ps.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}

	defer tx.Rollback()

	// найти общую стоимость продуктов для получения totalAmount - он понадобится при создании заказа
	for _, product := range *order.Products {
		rows, erro := ps.db.Query(
			`SELECT p.price
			FROM products p
			WHERE p.id = $1`,
			product.ID,
		)
		if erro != nil {
			return erro
		}
		for rows.Next() {
			err = rows.Scan(&productPrice)
			if err != nil {
				return err
			}
			totalAmount += productPrice
		}
	}

	// этот инсерт должен сработать таким образом, чтобы вытащился ИД вставленного заказа
	// он потребуется для вставки в order_products
	err = tx.QueryRow(
		"INSERT INTO orders (user_id, order_date, total_amount) VALUES ($1, $2, $3) RETURNING id",
		order.UserID, time.Now(), totalAmount,
	).Scan(&lastInsertID)
	if err != nil {
		log.Println(err)
		return err
	}

	// создаем записи в order_products
	for _, product := range *order.Products {
		_, err = tx.Exec(
			"INSERT INTO order_products (order_id, product_id) VALUES ($1, $2)",
			lastInsertID, product.ID,
		)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	// Заказ создан, продукты по заказу созданы - фиксируем транзакцию
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Заказ успешно создан.")
	return nil
}

// Удаление заказа.
func (ps *OrderService) DeleteOrder(orderID int) error {
	tx, err := ps.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM orders WHERE id=$1", orderID)
	if err != nil {
		log.Println(err)
		return err
	}

	// Заказ удален, фиксируем транзакцию
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Заказ успешно удален.")
	return nil
}

// Получение всех заказов по пользователю.
func (ps *OrderService) GetOrdersByUser(userID int) ([]structs.Order, error) {
	rows, err := ps.db.Query(
		`SELECT o.id, o.user_id, o.order_date, o.total_amount
		FROM orders o
		join users u ON o.user_id = u.id
		WHERE u.id = $1`,
		userID,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orders []structs.Order
	for rows.Next() {
		var order structs.Order
		err = rows.Scan(&order.ID, &order.UserID, &order.OrderDate, &order.TotalAmount)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

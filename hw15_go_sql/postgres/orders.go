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

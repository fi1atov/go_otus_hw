package postgres

import (
	"context"
	"database/sql"
	"log"

	"github.com/fi1atov/go_otus_hw/hw15_go_sql/structs"
)

type ProductService struct {
	db *DB
}

func NewProductService(db *DB) *ProductService {
	return &ProductService{db}
}

// Получение списка всех продуктов.
func (ps *ProductService) GetProducts() ([]structs.Product, error) {
	rows, err := ps.db.Query(`SELECT s.id, s.name FROM products s`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []structs.Product
	for rows.Next() {
		var product structs.Product
		err = rows.Scan(&product.ID, &product.Name)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

// Создание продукта.
func (ps *ProductService) CreateProduct(product *structs.ProductPatch) error {
	tx, err := ps.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO products (name, price) VALUES ($1, $2)", product.Name, product.Price)
	if err != nil {
		log.Println(err)
		return err
	}

	// Продукт создан, фиксируем транзакцию
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Продукт успешно создан.")
	return nil
}

// Обновление продукта.
func (ps *ProductService) UpdateProduct(productID int, product *structs.Product, patch structs.ProductPatch) error {
	tx, err := ps.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if v := patch.Name; v != nil {
		product.Name = *v
	}

	if v := patch.Price; v != nil {
		product.Price = *v
	}

	_, err = tx.Exec("UPDATE products SET name=$2, price=$3 WHERE id=$1", productID, product.Name, product.Price)
	if err != nil {
		log.Println(err)
		return err
	}

	// Продукт обновлен, фиксируем транзакцию
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Продукт успешно обновлен.")
	return nil
}

// Удаление продукта.
func (ps *ProductService) DeleteProduct(productID int) error {
	tx, err := ps.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM products WHERE id=$1", productID)
	if err != nil {
		log.Println(err)
		return err
	}

	// Продукт создан, фиксируем транзакцию
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Продукт успешно удален.")
	return nil
}

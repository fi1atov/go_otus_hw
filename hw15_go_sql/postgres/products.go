package postgres

import (
	"context"

	"github.com/fi1atov/go_otus_hw/hw15_go_sql/structs"
)

type ProductService struct {
	dbpool *DBPool
}

func NewProductService(dbpool *DBPool) *ProductService {
	return &ProductService{dbpool}
}

// Получение списка всех продуктов.
func (ps *ProductService) GetProducts() ([]structs.Product, error) {
	rows, err := ps.dbpool.Query(context.TODO(), `SELECT s.id, s.name FROM products s`)
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

package postgres

import (
	"context"

	"github.com/fi1atov/go_otus_hw/hw15_go_sql/structs"
)

type ProductService struct {
	dbpool *DBPool
	ctx    context.Context
}

func NewProductService(ctx context.Context, dbpool *DBPool) *ProductService {
	return &ProductService{dbpool, ctx}
}

// Получение списка всех продуктов.
func (ps *ProductService) GetProducts() ([]structs.Product, error) {
	rows, err := ps.dbpool.Query(ps.ctx, `SELECT s.id, s.name FROM products s`)
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

package structs

// Структура для хранения данных о продукте.
type Product struct {
	ID    int
	Name  string
	Price float32
}

type ProductPatch struct {
	Name  *string  `json:"name"`
	Price *float32 `json:"price"`
}

type ProductService interface {
	GetProducts() ([]Product, error)

	CreateProduct(*ProductPatch) error

	UpdateProduct(int, *Product, ProductPatch) error

	DeleteProduct(int) error
}

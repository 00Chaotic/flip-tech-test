package postgres

import (
	"fmt"

	"github.com/00Chaotic/flip-tech-test/backend/internal/model"
	"github.com/jmoiron/sqlx"
)

type ProductDAO struct {
	dbx *sqlx.DB
}

func NewProductDAO(dbx *sqlx.DB) *ProductDAO {
	return &ProductDAO{dbx: dbx}
}

// GetProducts retrieves all records in the Product table.
// Ideally pagination or some other mechanism would be implemented because we shouldn't
// be getting ALL records in a table, but it's fine for the purpose of this exercise.
func (d *ProductDAO) GetProducts() ([]*model.Product, error) {
	var products []*model.Product

	query := `SELECT * FROM Product`

	err := d.dbx.Select(&products, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}

	return products, nil
}

// GetProductBySKU retrieves a Product by its unique SKU.
func (d *ProductDAO) GetProductBySKU(sku string) (*model.Product, error) {
	var product model.Product

	query := `SELECT * FROM Product WHERE SKU = $1`

	err := d.dbx.Get(&product, query, sku)
	if err != nil {
		return nil, fmt.Errorf("failed to get product by sku: %w", err)
	}

	return &product, nil
}

// UpdateProductInventory updates a Product inventory quantity by adding the difference parameter.
// A negative difference will subtract from the Product inventory.
func (d *ProductDAO) UpdateProductInventory(sku string, difference int) error {
	return withTransaction(d.dbx, func(tx *sqlx.Tx) error {
		query := `UPDATE Product SET Inventory = Inventory + $1 WHERE SKU = $2`

		_, err := tx.Exec(query, difference, sku)
		if err != nil {
			return fmt.Errorf("failed to update product inventory: %w", err)
		}

		return nil
	})
}

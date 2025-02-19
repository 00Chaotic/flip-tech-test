package postgres

import (
	"context"
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
func (d *ProductDAO) GetProducts(ctx context.Context) ([]*model.Product, error) {
	var products []*model.Product

	query := `SELECT * FROM Product`

	err := d.dbx.Select(&products, query)
	if err != nil {
		return nil, err
	}

	return products, nil
}

// GetProductBySKU retrieves a Product by its unique SKU.
func (d *ProductDAO) GetProductBySKU(ctx context.Context, sku string) (*model.Product, error) {
	var product model.Product

	query := `SELECT * FROM Product WHERE SKU = $1`

	err := d.dbx.GetContext(ctx, &product, query, sku)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

// UpdateProductInventory updates a Product inventory quantity by adding the difference parameter.
// A negative difference will subtract from the Product inventory.
func (d *ProductDAO) UpdateProductInventory(ctx context.Context, sku string, difference int) error {
	return withTransaction(ctx, d.dbx, func(tx *sqlx.Tx) error {
		query := `UPDATE Product SET Inventory = Inventory + $1 WHERE SKU = $2`

		_, err := tx.ExecContext(ctx, query, difference, sku)
		if err != nil {
			return err
		}

		return nil
	})
}

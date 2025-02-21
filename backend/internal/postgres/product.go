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

	query := `SELECT * FROM Product ORDER BY SKU ASC`

	err := d.dbx.SelectContext(ctx, &products, query)
	if err != nil {
		return nil, err
	}

	return products, nil
}

// UpdateProductInventories updates Product inventory quantities by adding the difference parameter.
// A negative difference will subtract from the Product inventory.
func (d *ProductDAO) UpdateProductInventories(ctx context.Context, items []model.PurchaseItem) (float64, []*model.Product, error) {
	var products []*model.Product
	totalPrice := float64(0)

	return totalPrice, products, withTransaction(ctx, d.dbx, func(tx *sqlx.Tx) error {
		for _, item := range items {
			var product model.Product

			query := `UPDATE Product
				SET Inventory = Inventory - $1
				WHERE SKU = $2 AND Inventory - $1 >= 0
				RETURNING *`

			err := tx.GetContext(ctx, &product, query, item.Quantity, item.SKU)
			if err != nil {
				return err
			}

			totalPrice += float64(item.Quantity) * product.Price
			products = append(products, &product)
		}

		return nil
	})
}

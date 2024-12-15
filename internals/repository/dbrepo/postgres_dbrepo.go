package dbrepo

import (
	"context"
	"database/sql"
	"github.com/PAPAvision-co/beautychi-backend.git/internals/models"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

func (r *PostgresDBRepo) Connection() *sql.DB {
	return r.DB
}

func (r *PostgresDBRepo) AllProducts() ([]*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT p.id, title, slug, p.description, poster, poster_key, price, category_id, brand_id, product_stock, product_discount_price, sub_category_id, consumer_guide, contact, status,pc.name,b.name,sc.name, p.created_at, p.updated_at FROM product p JOIN category pc on p.category_id = pc.id JOIN brand b on p.brand_id = b.id JOIN sub_category sc on p.sub_category_id = sc.id;
	`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product

	for rows.Next() {
		var pro models.Product
		err := rows.Scan(
			&pro.ID,
			&pro.Title,
			&pro.Slug,
			&pro.Description,
			&pro.Poster,
			&pro.PosterKey,
			&pro.Price,
			&pro.CategoryID,
			&pro.BrandID,
			&pro.ProductStock,
			&pro.ProductDiscountPrice,
			&pro.SubCategoryID,
			&pro.ConsumerGuide,
			&pro.Contact,
			&pro.Status,
			&pro.CategoryName,
			&pro.BrandName,
			&pro.SubCategoryName,
			&pro.CreatedAt,
			&pro.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, &pro)
	}

	return products, nil
}

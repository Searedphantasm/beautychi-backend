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

func (r *PostgresDBRepo) AllCategories() ([]*models.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT id, name, description, image, image_key, created_at, updated_at FROM category;
	`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*models.Category

	for rows.Next() {
		var category models.Category
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
			&category.Image,
			&category.ImageKey,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		categories = append(categories, &category)
	}

	return categories, nil
}

func (r *PostgresDBRepo) AllBrands() ([]*models.Brand, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT id, name, description, country, logo, logo_key, website_url, created_at, updated_at FROM brand;
	`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var brands []*models.Brand

	for rows.Next() {
		var brand models.Brand
		err := rows.Scan(
			&brand.ID,
			&brand.Name,
			&brand.Description,
			&brand.Country,
			&brand.Logo,
			&brand.LogoKey,
			&brand.Website,
			&brand.CreatedAt,
			&brand.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		brands = append(brands, &brand)
	}

	return brands, nil
}

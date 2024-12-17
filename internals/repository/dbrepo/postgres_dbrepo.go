package dbrepo

import (
	"context"
	"database/sql"
	"errors"
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

func (r *PostgresDBRepo) InsertProduct(product models.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	transaction, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO product (title, slug, description, poster, poster_key, price, category_id, brand_id, product_stock, product_discount_price, sub_category_id, consumer_guide, contact, status) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14) returning id;`

	var newProductID int

	err = transaction.QueryRowContext(ctx, stmt,
		product.Title,
		product.Slug,
		product.Description,
		product.Poster,
		product.PosterKey,
		product.Price,
		product.CategoryID,
		product.BrandID,
		product.ProductStock,
		product.ProductDiscountPrice,
		product.SubCategoryID,
		product.ConsumerGuide,
		product.Contact,
		product.Status,
	).Scan(&newProductID)

	if err != nil {
		transaction.Rollback()
		return errors.New("failed to insert product")
	}

	for _, ps := range product.ProductSpecs {
		stmt = `INSERT INTO product_specifications (product_id, specs_title, specs_description) VALUES ($1,$2,$3)`
		_, err := transaction.ExecContext(ctx, stmt,
			newProductID,
			ps.SpecsTitle,
			ps.SpecsDescription,
		)
		if err != nil {
			transaction.Rollback()
			return errors.New("failed to insert product specifications")
		}
	}

	for _, pi := range product.ProductImages {
		stmt = `INSERT INTO product_image (product_id, url,alt_text) VALUES ($1,$2,$3)`
		_, err := transaction.ExecContext(ctx, stmt,
			newProductID,
			pi.Url,
			pi.AltText,
		)
		if err != nil {
			transaction.Rollback()
			return errors.New("failed to insert product images")
		}
	}

	err = transaction.Commit()
	if err != nil {
		transaction.Rollback()
		return errors.New("failed to commit transaction")
	}

	return nil
}

func (r *PostgresDBRepo) ProductByID(id int) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT p.id, title, slug, p.description, poster, poster_key, price, category_id, brand_id, product_stock, product_discount_price, sub_category_id, consumer_guide, contact, status,pc.name,b.name,sc.name, p.created_at, p.updated_at FROM product p JOIN category pc on p.category_id = pc.id JOIN brand b on p.brand_id = b.id JOIN sub_category sc on p.sub_category_id = sc.id WHERE p.id = $1;`

	row := r.DB.QueryRowContext(ctx, query, id)
	var pro models.Product
	err := row.Scan(
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

	query = `SELECT id, product_id, specs_title, specs_description FROM product_specifications WHERE product_id = $1;`
	rows, err := r.DB.QueryContext(ctx, query, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	defer rows.Close()

	var productSpecs []*models.ProductSpecification

	for rows.Next() {
		var productSpec models.ProductSpecification
		err := rows.Scan(
			&productSpec.ID,
			&productSpec.ProductID,
			&productSpec.SpecsTitle,
			&productSpec.SpecsDescription,
		)
		if err != nil {
			return nil, err
		}

		productSpecs = append(productSpecs, &productSpec)
	}

	query = `SELECT id, product_id, url, alt_text, created_at FROM product_image WHERE product_id = $1;`
	rows, err = r.DB.QueryContext(ctx, query, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	defer rows.Close()

	var productImages []*models.ProductImage

	for rows.Next() {
		var proImage models.ProductImage
		err := rows.Scan(
			&proImage.ID,
			&proImage.ProductID,
			&proImage.Url,
			&proImage.AltText,
			&proImage.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		productImages = append(productImages, &proImage)
	}

	pro.ProductSpecs = productSpecs
	pro.ProductImages = productImages

	return &pro, nil
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

func (r *PostgresDBRepo) AllSubCategories() ([]*models.SubCategory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT id, parent_category_id, name, description, image, image_key, created_at, updated_at FROM sub_category;
	`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subCategories []*models.SubCategory
	for rows.Next() {
		var subCategory models.SubCategory
		err := rows.Scan(
			&subCategory.ID,
			&subCategory.ParentCategoryID,
			&subCategory.Name,
			&subCategory.Description,
			&subCategory.Image,
			&subCategory.ImageKey,
			&subCategory.CreatedAt,
			&subCategory.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		subCategories = append(subCategories, &subCategory)
	}

	return subCategories, nil
}

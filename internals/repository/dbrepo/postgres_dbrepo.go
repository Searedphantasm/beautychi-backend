package dbrepo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/PAPAvision-co/beautychi-backend.git/internals/models"
	validation "github.com/go-ozzo/ozzo-validation"
	"log"
	"reflect"
	"strings"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

func (r *PostgresDBRepo) Connection() *sql.DB {
	return r.DB
}

// insert is a helper function to put a value on specific index in slice
func insert(slice []any, index int, value any) []any {
	// make room for the new element by creating a new slice with one additional element
	result := make([]any, len(slice)+1)

	// copy the part of the slice before the insertion point
	copy(result, slice[:index])

	// insert the new element
	result[index] = value

	// merging both part of the slices
	copy(result[index+1:], slice[index:])

	return result
}

func (r *PostgresDBRepo) AllProducts(limit, offset int, optionalParams models.OptionalQueryParams) ([]*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var query string
	var args []any
	// pre-append limit and offset
	args = append(args, limit, offset)

	// Base query
	baseQuery := `
    SELECT p.id, title, p.slug, p.description, poster, poster_key, price, category_id, brand_id, product_stock, coalesce(product_discount_price,0), sub_category_id, consumer_guide, contact, status,pc.slug as category_slug,pc.name,b.name,sc.name, p.created_at, p.updated_at 
    FROM product p 
    JOIN category pc ON p.category_id = pc.id 
    JOIN brand b ON p.brand_id = b.id 
    JOIN sub_category sc ON p.sub_category_id = sc.id
    `

	// Conditions for WHERE clause
	var conditions []string
	var conditionIndex int = 1

	if !validation.IsEmpty(optionalParams.Search) {
		conditions = append(conditions, fmt.Sprintf("p.slug ILIKE $%d", conditionIndex))
		args = append([]any{"%" + optionalParams.Search + "%"}, args...)
		conditionIndex++

		log.Println(args)
		log.Println(conditions)
	}

	if !validation.IsEmpty(optionalParams.ProductCategory) {
		conditions = append(conditions, fmt.Sprintf("pc.slug = $%d", conditionIndex))
		if len(args) > 2 {
			args = insert(args, 1, optionalParams.ProductCategory)
		} else {
			args = append([]any{optionalParams.ProductCategory}, args...)
		}
		conditionIndex++

		log.Println(args)
		log.Println(conditions)
	}

	// FIXME: The order is not correct for this params,
	// NOTE: NOT READY
	if !validation.IsEmpty(optionalParams.Filter) {

		conditions = append(conditions, fmt.Sprintf("p.created_at::date = $%d", conditionIndex))
		args = append([]any{optionalParams.Filter}, args...)
		conditionIndex++
	}

	// Construct WHERE clause if there are any conditions
	if len(conditions) > 0 {
		query = baseQuery + " WHERE " + strings.Join(conditions, " AND ") + fmt.Sprintf(" LIMIT $%d OFFSET $%d;", conditionIndex, conditionIndex+1)
		log.Println(query)
	} else {
		query = baseQuery + " LIMIT $1 OFFSET $2;"
	}

	rows, err := r.DB.QueryContext(ctx, query, args...)
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
			&pro.CategorySlug,
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
		stmt = `INSERT INTO product_image (product_id, url,url_key,alt_text) VALUES ($1,$2,$3,$4)`
		_, err := transaction.ExecContext(ctx, stmt,
			newProductID,
			pi.Url,
			pi.UrlKey,
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

func (r *PostgresDBRepo) UpdateProduct(product models.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// getting existing product
	transaction, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var existingProduct models.Product

	query := `SELECT id, title, slug, description, poster, poster_key, price, category_id, brand_id, product_stock, coalesce(product_discount_price,0), sub_category_id, consumer_guide, contact, status,created_at, updated_at FROM product WHERE id = $1;`

	row := transaction.QueryRowContext(ctx, query, product.ID)
	err = row.Scan(
		&existingProduct.ID,
		&existingProduct.Title,
		&existingProduct.Slug,
		&existingProduct.Description,
		&existingProduct.Poster,
		&existingProduct.PosterKey,
		&existingProduct.Price,
		&existingProduct.CategoryID,
		&existingProduct.BrandID,
		&existingProduct.ProductStock,
		&existingProduct.ProductDiscountPrice,
		&existingProduct.SubCategoryID,
		&existingProduct.ConsumerGuide,
		&existingProduct.Contact,
		&existingProduct.Status,
		&existingProduct.CreatedAt,
		&existingProduct.UpdatedAt,
	)

	if err != nil {
		transaction.Rollback()
		return err
	}
	// product

	dbValue := reflect.ValueOf(&existingProduct).Elem()
	updateValue := reflect.ValueOf(&product).Elem()

	for i := 0; i < dbValue.NumField(); i++ {
		dbField := dbValue.Field(i)
		updateField := updateValue.Field(i)

		if updateField.IsValid() && dbField.CanSet() {
			dbField.Set(updateField)
		}
	}

	existingProduct.UpdatedAt = time.Now()

	stmt := `UPDATE product SET title = $1, slug = $2, description = $3, poster = $4 , poster_key = $5, 
                   price = $6 , category_id = $7 , brand_id = $8 , product_stock = $9, product_discount_price = $10 , sub_category_id = $11, consumer_guide = $12 , contact = $13 , status = $14 , updated_at = $15 WHERE id = $16;`

	_, err = transaction.ExecContext(ctx, stmt,
		existingProduct.Title,
		existingProduct.Slug,
		existingProduct.Description,
		existingProduct.Poster,
		existingProduct.PosterKey,
		existingProduct.Price,
		existingProduct.CategoryID,
		existingProduct.BrandID,
		existingProduct.ProductStock,
		existingProduct.ProductDiscountPrice,
		existingProduct.SubCategoryID,
		existingProduct.ConsumerGuide,
		existingProduct.Contact,
		existingProduct.Status,
		existingProduct.UpdatedAt,
		existingProduct.ID,
	)
	if err != nil {
		transaction.Rollback()
		return errors.New("failed to update product ")
	}

	err = transaction.Commit()
	if err != nil {
		transaction.Rollback()
		return errors.New("failed to commit update product transaction")
	}

	return nil
}

func (r *PostgresDBRepo) UpdateProductImages(productID int, productImages []models.ProductImage) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	transaction, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	query := `SELECT id, title FROM product WHERE id = $1;`

	var checkProduct struct {
		Id    int
		Title string
	}

	err = transaction.QueryRowContext(ctx, query, productID).Scan(
		&checkProduct.Id,
		&checkProduct.Title,
	)
	if err != nil {
		transaction.Rollback()
		return err
	}

	for _, pi := range productImages {
		stmt := `INSERT INTO product_image (product_id, url,url_key,alt_text) VALUES ($1,$2,$3,$4)`
		_, err := transaction.ExecContext(ctx, stmt,
			productID,
			pi.Url,
			pi.UrlKey,
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

func (r *PostgresDBRepo) ProductByIDOrSlug(params models.OneParams) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	id := params.ID
	slug := params.Slug
	var query string
	var identifier any

	if id == 0 && len(slug) > 0 {
		query = `SELECT p.id, title, p.slug, p.description, poster, poster_key, price, category_id, brand_id, product_stock, product_discount_price, sub_category_id, consumer_guide, contact, status,pc.name,b.name,sc.name, p.created_at, p.updated_at FROM product p JOIN category pc on p.category_id = pc.id JOIN brand b on p.brand_id = b.id JOIN sub_category sc on p.sub_category_id = sc.id WHERE p.slug = $1;`

		identifier = slug
	} else if id != 0 {
		query = `SELECT p.id, title, p.slug, p.description, poster, poster_key, price, category_id, brand_id, product_stock, product_discount_price, sub_category_id, consumer_guide, contact, status,pc.name,b.name,sc.name, p.created_at, p.updated_at FROM product p JOIN category pc on p.category_id = pc.id JOIN brand b on p.brand_id = b.id JOIN sub_category sc on p.sub_category_id = sc.id WHERE p.id = $1;`

		identifier = id
	}

	row := r.DB.QueryRowContext(ctx, query, identifier)

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
	rows, err := r.DB.QueryContext(ctx, query, pro.ID)
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

	query = `SELECT id, product_id, url,url_key, alt_text, created_at FROM product_image WHERE product_id = $1;`
	rows, err = r.DB.QueryContext(ctx, query, pro.ID)
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
			&proImage.UrlKey,
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

func (r *PostgresDBRepo) DeleteProductByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `DELETE FROM product WHERE id = $1;`

	_, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresDBRepo) AllCategories(optionalParams models.OptionalQueryParams) ([]*models.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var query string
	var args []any
	if !validation.IsEmpty(optionalParams.Search) {
		query = ` SELECT id, name,slug, description, image, image_key, created_at, updated_at FROM category WHERE slug ILIKE '%' || $1 || '%' ; `

		args = append(args, optionalParams.Search)
	} else {
		query = `
		SELECT id, name,slug, description, image, image_key, created_at, updated_at FROM category;
		`
	}

	rows, err := r.DB.QueryContext(ctx, query, args...)
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
			&category.Slug,
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

func (r *PostgresDBRepo) InsertCategory(category models.Category) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `INSERT INTO category (name,slug, description, image, image_key) VALUES ($1,$2, $3, $4,$5);`

	_, err := r.DB.ExecContext(ctx, stmt,
		category.Name,
		category.Slug,
		category.Description,
		category.Image,
		category.ImageKey,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresDBRepo) AllBrands() ([]*models.Brand, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT id, name,slug, description, country, logo, logo_key, website_url, created_at, updated_at FROM brand;
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
			&brand.Slug,
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

func (r *PostgresDBRepo) InsertBrand(brand models.Brand) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `INSERT INTO brand (name, slug, description, country, logo, logo_key, website_url) VALUES ($1,$2,$3,$4,$5,$6,$7);`
	_, err := r.DB.ExecContext(ctx, stmt,
		brand.Name,
		brand.Slug,
		brand.Description,
		brand.Country,
		brand.Logo,
		brand.LogoKey,
		brand.Website,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresDBRepo) DeleteBrandByID(brandID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// check if brand exists
	params := models.OneParams{
		ID:   brandID,
		Slug: "",
	}
	existing, err := r.GetBrandByIDOrSlug(params)
	if err != nil {
		return err
	}
	if existing == nil {
		return nil
	}

	query := `DELETE FROM brand WHERE id = $1;`
	_, err = r.DB.ExecContext(ctx, query, brandID)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresDBRepo) UpdateBrand(brand models.Brand) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	params := models.OneParams{
		ID:   brand.ID,
		Slug: "",
	}
	existingBrand, err := r.GetBrandByIDOrSlug(params)
	if err != nil {
		return err
	}

	if existingBrand == nil {
		return nil
	}

	dbValue := reflect.ValueOf(existingBrand).Elem()
	updateValue := reflect.ValueOf(&brand).Elem()

	for i := 0; i < dbValue.NumField(); i++ {
		dbField := dbValue.Field(i)
		updateField := updateValue.Field(i)

		if updateField.IsValid() && dbField.CanSet() {
			dbField.Set(updateField)
		}
	}

	existingBrand.UpdatedAt = time.Now()

	stmt := `UPDATE brand SET name = $1, slug = $2, description = $3 , logo = $4, logo_key = $5, updated_at = $6,country = $7,website_url = $8 WHERE id = $9;`

	_, err = r.DB.ExecContext(ctx, stmt,
		existingBrand.Name,
		existingBrand.Slug,
		existingBrand.Description,
		existingBrand.Logo,
		existingBrand.LogoKey,
		existingBrand.UpdatedAt,
		existingBrand.Country,
		existingBrand.Website,
		existingBrand.ID,
	)
	if err != nil {
		return err
	}

	return nil

}

func (r *PostgresDBRepo) GetBrandByIDOrSlug(params models.OneParams) (*models.Brand, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	brandID := params.ID
	slug := params.Slug
	var brand models.Brand
	var identifier any
	var query string

	if brandID == 0 && len(slug) > 0 {
		query = `SELECT id, name, slug, description, country, logo, logo_key, website_url, created_at, updated_at FROM brand WHERE slug = $1;`

		identifier = slug
	} else if brandID != 0 {
		query = `SELECT id, name, slug, description, country, logo, logo_key, website_url, created_at, updated_at FROM brand WHERE id = $1;`

		identifier = brandID
	}

	err := r.DB.QueryRowContext(ctx, query, identifier).Scan(
		&brand.ID,
		&brand.Name,
		&brand.Slug,
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

	return &brand, nil
}

func (r *PostgresDBRepo) AllSubCategories() ([]*models.SubCategory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT sc.id, parent_category_id,c.name as parent_category_name, sc.name,sc.slug, sc.description, sc.image, sc.image_key, sc.created_at, sc.updated_at FROM sub_category sc JOIN category c on sc.parent_category_id = c.id;`

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
			&subCategory.ParentCategoryName,
			&subCategory.Name,
			&subCategory.Slug,
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

func (r *PostgresDBRepo) InsertSubCategory(subCategory models.SubCategory) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `INSERT INTO sub_category (parent_category_id, name, slug, description, image, image_key) VALUES ($1, $2, $3, $4, $5, $6);`
	_, err := r.DB.ExecContext(ctx, stmt,
		subCategory.ParentCategoryID,
		subCategory.Name,
		subCategory.Slug,
		subCategory.Description,
		subCategory.Image,
		subCategory.ImageKey,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresDBRepo) DeleteSubCategoryByID(subCategoryID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `DELETE FROM sub_category WHERE id = $1;`
	_, err := r.DB.ExecContext(ctx, stmt, subCategoryID)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresDBRepo) GetSubCategoryByIDOrSlug(params models.OneParams) (*models.SubCategory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	subCategoryID := params.ID
	slug := params.Slug

	var subCategory models.SubCategory
	var identifier any
	var query string

	if subCategoryID == 0 && len(slug) > 0 {
		query = `
		SELECT sc.id, parent_category_id,c.name as parent_category_name, sc.name,sc.slug, sc.description, sc.image, sc.image_key, sc.created_at, sc.updated_at FROM sub_category sc JOIN category c on sc.parent_category_id = c.id WHERE sc.slug = $1;`

		identifier = slug
	} else if subCategoryID != 0 {
		query = `
		SELECT sc.id, parent_category_id,c.name as parent_category_name, sc.name,sc.slug, sc.description, sc.image, sc.image_key, sc.created_at, sc.updated_at FROM sub_category sc JOIN category c on sc.parent_category_id = c.id WHERE sc.id = $1;`

		identifier = subCategoryID
	}

	err := r.DB.QueryRowContext(ctx, query, identifier).Scan(
		&subCategory.ID,
		&subCategory.ParentCategoryID,
		&subCategory.ParentCategoryName,
		&subCategory.Name,
		&subCategory.Slug,
		&subCategory.Description,
		&subCategory.Image,
		&subCategory.ImageKey,
		&subCategory.CreatedAt,
		&subCategory.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &subCategory, nil
}

func (r *PostgresDBRepo) UpdateSubCategory(subCategory models.SubCategory) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// get existing subcategory
	params := models.OneParams{
		ID:   subCategory.ID,
		Slug: "",
	}
	existingSubCategory, err := r.GetSubCategoryByIDOrSlug(params)
	if err != nil {
		return err
	}

	if existingSubCategory == nil {
		return nil
	}

	dbValue := reflect.ValueOf(existingSubCategory).Elem()
	updateValue := reflect.ValueOf(&subCategory).Elem()

	for i := 0; i < dbValue.NumField(); i++ {
		dbField := dbValue.Field(i)
		updateField := updateValue.Field(i)

		if updateField.IsValid() && dbField.CanSet() {
			dbField.Set(updateField)
		}
	}

	existingSubCategory.UpdatedAt = time.Now()

	stmt := `UPDATE sub_category SET name = $1,slug = $2, description = $3, image = $4 , image_key = $5, parent_category_id = $6 WHERE id = $7;`

	_, err = r.DB.ExecContext(ctx, stmt,
		existingSubCategory.Name,
		existingSubCategory.Slug,
		existingSubCategory.Description,
		existingSubCategory.Image,
		existingSubCategory.ImageKey,
		existingSubCategory.ParentCategoryID,
		existingSubCategory.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

// AllSubCategoriesByParentCategoryID gets all subcategories by parent_category_id
func (r *PostgresDBRepo) AllSubCategoriesByParentCategoryID(parentCategoryID int) ([]*models.SubCategory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT sc.id, parent_category_id,c.name as parent_category_name, sc.name,sc.slug, sc.description, sc.image, sc.image_key, sc.created_at, sc.updated_at FROM sub_category sc JOIN category c on sc.parent_category_id = c.id WHERE parent_category_id = $1;`

	rows, err := r.DB.QueryContext(ctx, query, parentCategoryID)
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
			&subCategory.ParentCategoryName,
			&subCategory.Name,
			&subCategory.Slug,
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

func (r *PostgresDBRepo) UpdateCategoryByID(category models.Category) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// check if the category exists
	var existingCategory models.Category

	query := `SELECT id, name, slug, description, image, image_key, created_at, updated_at FROM category WHERE id = $1;`

	row := r.DB.QueryRowContext(ctx, query, category.ID)
	err := row.Scan(
		&existingCategory.ID,
		&existingCategory.Name,
		&existingCategory.Slug,
		&existingCategory.Description,
		&existingCategory.Image,
		&existingCategory.ImageKey,
		&existingCategory.CreatedAt,
		&existingCategory.UpdatedAt,
	)
	if err != nil {
		return err
	}

	dbValue := reflect.ValueOf(&existingCategory).Elem()
	updateValue := reflect.ValueOf(&category).Elem()

	for i := 0; i < dbValue.NumField(); i++ {
		dbField := dbValue.Field(i)
		updateField := updateValue.Field(i)

		if updateField.IsValid() && dbField.CanSet() {
			dbField.Set(updateField)
		}
	}

	existingCategory.UpdatedAt = time.Now()

	stmt := `UPDATE category SET name = $1, slug = $2, description = $3 , image = $4, image_key = $5, updated_at = $6 WHERE id = $7;`

	_, err = r.DB.ExecContext(ctx, stmt,
		existingCategory.Name,
		existingCategory.Slug,
		existingCategory.Description,
		existingCategory.Image,
		existingCategory.ImageKey,
		existingCategory.UpdatedAt,
		existingCategory.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresDBRepo) GetCategoryByIDOrSlug(params models.OneParams) (*models.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	id := params.ID
	slug := params.Slug

	var query string
	var identifier any

	if id == 0 && len(slug) > 0 {
		query = `SELECT id, name, slug, description, image, image_key, created_at, updated_at FROM category WHERE slug = $1;`
		identifier = slug
	} else if id != 0 {
		query = `SELECT id, name, slug, description, image, image_key, created_at, updated_at FROM category WHERE id = $1;`
		identifier = id
	}

	var category models.Category

	row := r.DB.QueryRowContext(ctx, query, identifier)
	err := row.Scan(
		&category.ID,
		&category.Name,
		&category.Slug,
		&category.Description,
		&category.Image,
		&category.ImageKey,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *PostgresDBRepo) DeleteCategoryByID(categoryID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// check if category exists
	params := models.OneParams{
		ID:   categoryID,
		Slug: "",
	}
	exsitingCategory, err := r.GetCategoryByIDOrSlug(params)
	if err != nil {
		return err
	}

	if exsitingCategory.ID == 0 {
		return errors.New("category does not exist")
	}

	query := `DELETE FROM category WHERE id = $1;`

	_, err = r.DB.ExecContext(ctx, query, categoryID)
	if err != nil {
		return err
	}

	return nil
}

// CUSTOMERS

func (r *PostgresDBRepo) AllCustomers(limit, offset int, optionalParams models.OptionalQueryParams) ([]*models.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var query string
	var args []any
	// pre-append limit and offset
	args = append(args, limit, offset)

	// Base query
	baseQuery := `SELECT id, username, first_name, last_name, email, phone, created_at, updated_at FROM customer`

	var conditions []string
	var conditionIndex int = 1

	if !validation.IsEmpty(optionalParams.Search) {
		conditions = append(conditions, fmt.Sprintf("phone ILIKE $%d", conditionIndex))
		args = append([]any{"%" + optionalParams.Search + "%"}, args...)
		conditionIndex++
	}

	if len(conditions) > 0 {
		query = baseQuery + " WHERE " + strings.Join(conditions, " AND ") + fmt.Sprintf(" LIMIT $%d OFFSET $%d;", conditionIndex, conditionIndex+1)
		log.Println(query)
	} else {
		query = baseQuery + " LIMIT $1 OFFSET $2;"
		log.Println(query)
	}

	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []*models.Customer

	for rows.Next() {
		var customer models.Customer
		err := rows.Scan(
			&customer.ID,
			&customer.Username,
			&customer.FirstName,
			&customer.LastName,
			&customer.Email,
			&customer.Phone,
			&customer.CreatedAt,
			&customer.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		customers = append(customers, &customer)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}

func (r *PostgresDBRepo) OneCustomerByID(id string) (*models.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT id, username, first_name, last_name, email, phone, created_at, updated_at FROM customer WHERE id = $1;`

	var customer models.Customer

	row := r.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&customer.ID,
		&customer.Username,
		&customer.FirstName,
		&customer.LastName,
		&customer.Email,
		&customer.Phone,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	query = `SELECT id, customer_id, city, state, address, postal_code, created_at, updated_at FROM customer_address
			WHERE customer_id = $1;`

	rows, err := r.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addresses []*models.CustomerAddress

	for rows.Next() {
		var address models.CustomerAddress
		err := rows.Scan(
			&address.ID,
			&address.CustomerID,
			&address.City,
			&address.State,
			&address.Address,
			&address.PostalCode,
			&address.CreatedAt,
			&address.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		addresses = append(addresses, &address)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	customer.Addresses = addresses
	return &customer, nil
}

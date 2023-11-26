package products

import (
	"context"
	"database/sql"
	"errors"

	"github.com/aldogayaladh/go-web-1598/internal/domain"
)

var (
	ErrPrepareStatement = errors.New("error prepare statement")
	ErrExecStatement    = errors.New("error exec statement")
	ErrLastInsertedId   = errors.New("error last inserted id")
)

type repositorymysql struct {
	db *sql.DB
}

// NewMemoryRepository ....
func NewMySqlRepository(db *sql.DB) Repository {
	return &repositorymysql{db: db}
}

// Create is a method that creates a new product.
func (r *repositorymysql) Create(ctx context.Context, producto domain.Producto) (domain.Producto, error) {
	statement, err := r.db.Prepare(QueryInsertProduct)
	if err != nil {
		return domain.Producto{}, ErrPrepareStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		producto.Name,
		producto.Quantity,
		producto.CodeValue,
		producto.IsPublished,
		producto.Expiration,
		producto.Price,
	)

	if err != nil {
		return domain.Producto{}, ErrExecStatement
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Producto{}, ErrLastInsertedId
	}

	producto.Id = int(lastId)

	return producto, nil

}

// GetAll is a method that returns all products.
func (r *repositorymysql) GetAll(ctx context.Context) ([]domain.Producto, error) {
	rows, err := r.db.Query(QueryGetAllProducts)
	if err != nil {
		return []domain.Producto{}, err
	}

	defer rows.Close()

	var productos []domain.Producto

	for rows.Next() {
		var producto domain.Producto
		err := rows.Scan(
			&producto.Id,
			&producto.Name,
			&producto.Quantity,
			&producto.CodeValue,
			&producto.IsPublished,
			&producto.Expiration,
			&producto.Price,
		)
		if err != nil {
			return []domain.Producto{}, err
		}

		productos = append(productos, producto)
	}

	if err := rows.Err(); err != nil {
		return []domain.Producto{}, err
	}

	return productos, nil
}

// GetByID is a method that returns a product by ID.
func (r *repositorymysql) GetByID(ctx context.Context, id int) (domain.Producto, error) {
	row := r.db.QueryRow(QueryGetProductById, id)

	var producto domain.Producto
	err := row.Scan(
		&producto.Id,
		&producto.Name,
		&producto.Quantity,
		&producto.CodeValue,
		&producto.IsPublished,
		&producto.Expiration,
		&producto.Price,
	)

	if err != nil {
		return domain.Producto{}, err
	}

	return producto, nil
}

// Update is a method that updates a product by ID.
func (r *repositorymysql) Update(
	ctx context.Context,
	producto domain.Producto,
	id int) (domain.Producto, error) {
	statement, err := r.db.Prepare(QueryUpdateProduct)
	if err != nil {
		return domain.Producto{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		producto.Name,
		producto.Quantity,
		producto.CodeValue,
		producto.IsPublished,
		producto.Expiration,
		producto.Price,
		producto.Id,
	)

	if err != nil {
		return domain.Producto{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.Producto{}, err
	}

	producto.Id = id

	return producto, nil

}

// Delete is a method that deletes a product by ID.
func (r *repositorymysql) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(QueryDeleteProduct, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return ErrNotFound
	}

	return nil
}

// Patch is a method that updates a product by ID.
func (r *repositorymysql) Patch(
	ctx context.Context,
	producto domain.Producto,
	id int) (domain.Producto, error) {
	statement, err := r.db.Prepare(QueryUpdateProduct)
	if err != nil {
		return domain.Producto{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		producto.Name,
		producto.Quantity,
		producto.CodeValue,
		producto.IsPublished,
		producto.Expiration,
		producto.Price,
		producto.Id,
	)

	if err != nil {
		return domain.Producto{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.Producto{}, err
	}

	return producto, nil
}

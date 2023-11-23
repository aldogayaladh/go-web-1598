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

// Create ....
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

// GetAll...
func (r *repositorymysql) GetAll(ctx context.Context) ([]domain.Producto, error) {
	panic("Implement me")
}

// GetByID .....
func (r *repositorymysql) GetByID(ctx context.Context, id int) (domain.Producto, error) {
	panic("Implement me")
}

// Update ...
func (r *repositorymysql) Update(
	ctx context.Context,
	producto domain.Producto,
	id int) (domain.Producto, error) {

	panic("Implement me")

}

// Delete ...
func (r *repositorymysql) Delete(ctx context.Context, id int) error {
	panic("Implement me")
}

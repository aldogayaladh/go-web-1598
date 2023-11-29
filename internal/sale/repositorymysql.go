package sale

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

type repositorymysqlsale struct {
	db *sql.DB
}

// NewMemoryRepository ....
func NewMySqlRepositorySale(db *sql.DB) Repository {
	return &repositorymysqlsale{db: db}
}

// Create is a method that creates a new sale.
func (r *repositorymysqlsale) Create(ctx context.Context, sale domain.Sale) (domain.Sale, error) {
	statement, err := r.db.Prepare(QueryInsertSale)
	if err != nil {
		return domain.Sale{}, ErrPrepareStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		sale.Description,
		sale.IdSeller,
		sale.IdProduct,
	)

	if err != nil {
		return domain.Sale{}, ErrExecStatement
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Sale{}, ErrLastInsertedId
	}

	sale.Id = int(lastId)

	return sale, nil

}

// GetAll is a method that returns all sales.
func (r *repositorymysqlsale) GetAll(ctx context.Context) ([]domain.Sale, error) {
	panic("implement me")
}

// GetByID is a method that returns a sale by ID.
func (r *repositorymysqlsale) GetByID(ctx context.Context, id int) (domain.Sale, error) {
	panic("implement me")
}

// Update is a method that updates a sale by ID.
func (r *repositorymysqlsale) Update(
	ctx context.Context,
	producto domain.Sale,
	id int) (domain.Sale, error) {
	panic("implement me")

}

// Delete is a method that deletes a sale by ID.
func (r *repositorymysqlsale) Delete(ctx context.Context, id int) error {
	panic("implement me")
}

// Patch is a method that updates a sale by ID.
func (r *repositorymysqlsale) Patch(
	ctx context.Context,
	producto domain.Sale,
	id int) (domain.Sale, error) {
	panic("implement me")
}

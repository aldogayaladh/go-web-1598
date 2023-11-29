package seller

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

type repositorymysqlseller struct {
	db *sql.DB
}

// NewMemoryRepository ....
func NewMySqlRepositorySeller(db *sql.DB) Repository {
	return &repositorymysqlseller{db: db}
}

// Create is a method that creates a new seller.
func (r *repositorymysqlseller) Create(ctx context.Context, seller domain.Seller) (domain.Seller, error) {
	statement, err := r.db.Prepare(QueryInsertSeller)
	if err != nil {
		return domain.Seller{}, ErrPrepareStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		seller.Description,
		seller.CodSeller,
		seller.IsAuthorization,
	)

	if err != nil {
		return domain.Seller{}, ErrExecStatement
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Seller{}, ErrLastInsertedId
	}

	seller.Id = int(lastId)

	return seller, nil

}

// GetAll is a method that returns all sellers.
func (r *repositorymysqlseller) GetAll(ctx context.Context) ([]domain.Seller, error) {
	panic("implement me")
}

// GetByID is a method that returns a seller by ID.
func (r *repositorymysqlseller) GetByID(ctx context.Context, id int) (domain.Seller, error) {
	panic("implement me")
}

// Update is a method that updates a seller by ID.
func (r *repositorymysqlseller) Update(
	ctx context.Context,
	producto domain.Seller,
	id int) (domain.Seller, error) {
	panic("implement me")

}

// Delete is a method that deletes a seller by ID.
func (r *repositorymysqlseller) Delete(ctx context.Context, id int) error {
	panic("implement me")
}

// Patch is a method that updates a seller by ID.
func (r *repositorymysqlseller) Patch(
	ctx context.Context,
	producto domain.Seller,
	id int) (domain.Seller, error) {
	panic("implement me")
}

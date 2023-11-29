package sale

import (
	"context"
	"log"

	"github.com/aldogayaladh/go-web-1598/internal/domain"
)

type Service interface {
	Create(ctx context.Context, sale domain.Sale) (domain.Sale, error)
	GetAll(ctx context.Context) ([]domain.Sale, error)
	GetByID(ctx context.Context, id int) (domain.Sale, error)
	Update(ctx context.Context, sale domain.Sale, id int) (domain.Sale, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, sale domain.Sale, id int) (domain.Sale, error)
}

type service struct {
	repository Repository
}

func NewServiceSale(repository Repository) Service {
	return &service{repository: repository}
}

// Create is a method that creates a new sale.
func (s *service) Create(ctx context.Context, sale domain.Sale) (domain.Sale, error) {
	sale, err := s.repository.Create(ctx, sale)
	if err != nil {
		log.Println("[ProductsService][Create] error creating sale", err)
		return domain.Sale{}, err
	}

	return sale, nil
}

// GetAll is a method that returns all sales.
func (s *service) GetAll(ctx context.Context) ([]domain.Sale, error) {
	panic("implement me")
}

// GetByID is a method that returns a sale by ID.
func (s *service) GetByID(ctx context.Context, id int) (domain.Sale, error) {
	panic("implement me")
}

// Update is a method that updates a sale by ID.
func (s *service) Update(ctx context.Context, sale domain.Sale, id int) (domain.Sale, error) {
	panic("implement me")
}

// Delete is a method that deletes a sale by ID.
func (s *service) Delete(ctx context.Context, id int) error {
	panic("implement me")
}

// Patch is a method that updates a sale by ID.
func (s *service) Patch(ctx context.Context, sale domain.Sale, id int) (domain.Sale, error) {
	panic("implement me")
}

// validatePatch is a method that validates the fields to be updated.
func (s *service) validatePatch(saleStore, sale domain.Sale) (domain.Sale, error) {

	panic("implement me")

}

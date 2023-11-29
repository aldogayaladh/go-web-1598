package seller

import (
	"context"
	"log"

	"github.com/aldogayaladh/go-web-1598/internal/domain"
)

type Service interface {
	Create(ctx context.Context, seller domain.Seller) (domain.Seller, error)
	GetAll(ctx context.Context) ([]domain.Seller, error)
	GetByID(ctx context.Context, id int) (domain.Seller, error)
	Update(ctx context.Context, seller domain.Seller, id int) (domain.Seller, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, seller domain.Seller, id int) (domain.Seller, error)
}

type service struct {
	repository Repository
}

func NewServiceSeller(repository Repository) Service {
	return &service{repository: repository}
}

// Create is a method that creates a new seller.
func (s *service) Create(ctx context.Context, seller domain.Seller) (domain.Seller, error) {
	seller, err := s.repository.Create(ctx, seller)
	if err != nil {
		log.Println("[ProductsService][Create] error creating seller", err)
		return domain.Seller{}, err
	}

	return seller, nil
}

// GetAll is a method that returns all sellers.
func (s *service) GetAll(ctx context.Context) ([]domain.Seller, error) {
	panic("implement me")
}

// GetByID is a method that returns a seller by ID.
func (s *service) GetByID(ctx context.Context, id int) (domain.Seller, error) {
	panic("implement me")
}

// Update is a method that updates a seller by ID.
func (s *service) Update(ctx context.Context, seller domain.Seller, id int) (domain.Seller, error) {
	panic("implement me")
}

// Delete is a method that deletes a seller by ID.
func (s *service) Delete(ctx context.Context, id int) error {
	panic("implement me")
}

// Patch is a method that updates a seller by ID.
func (s *service) Patch(ctx context.Context, seller domain.Seller, id int) (domain.Seller, error) {
	panic("implement me")
}

// validatePatch is a method that validates the fields to be updated.
func (s *service) validatePatch(sellerStore, seller domain.Seller) (domain.Seller, error) {

	if seller.Description != "" {
		sellerStore.Description = seller.Description
	}

	if seller.CodSeller != "" {
		sellerStore.CodSeller = seller.CodSeller
	}

	if seller.IsAuthorization != false {
		sellerStore.IsAuthorization = seller.IsAuthorization
	}

	return sellerStore, nil

}

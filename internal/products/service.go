package products

import (
	"context"
	"log"
	"time"

	"github.com/aldogayaladh/go-web-1598/internal/domain"
)

type Service interface {
	Create(ctx context.Context, producto domain.Producto) (domain.Producto, error)
	GetAll(ctx context.Context) ([]domain.Producto, error)
	GetByID(ctx context.Context, id int) (domain.Producto, error)
	Update(ctx context.Context, producto domain.Producto, id int) (domain.Producto, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, producto domain.Producto, id int) (domain.Producto, error)
}

type service struct {
	repository Repository
}

func NewServiceProduct(repository Repository) Service {
	return &service{repository: repository}
}

// Create is a method that creates a new product.
func (s *service) Create(ctx context.Context, producto domain.Producto) (domain.Producto, error) {
	producto, err := s.repository.Create(ctx, producto)
	if err != nil {
		log.Println("[ProductsService][Create] error creating product", err)
		return domain.Producto{}, err
	}

	return producto, nil
}

// GetAll is a method that returns all products.
func (s *service) GetAll(ctx context.Context) ([]domain.Producto, error) {
	listProducts, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Println("[ProductsService][GetAll] error getting all products", err)
		return []domain.Producto{}, err
	}

	return listProducts, nil
}

// GetByID is a method that returns a product by ID.
func (s *service) GetByID(ctx context.Context, id int) (domain.Producto, error) {
	producto, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[ProductsService][GetByID] error getting product by ID", err)
		return domain.Producto{}, err
	}

	return producto, nil
}

// Update is a method that updates a product by ID.
func (s *service) Update(ctx context.Context, producto domain.Producto, id int) (domain.Producto, error) {
	producto, err := s.repository.Update(ctx, producto, id)
	if err != nil {
		log.Println("[ProductsService][Update] error updating product by ID", err)
		return domain.Producto{}, err
	}

	return producto, nil
}

// Delete is a method that deletes a product by ID.
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("[ProductsService][Delete] error deleting product by ID", err)
		return err
	}

	return nil
}

// Patch is a method that updates a product by ID.
func (s *service) Patch(ctx context.Context, producto domain.Producto, id int) (domain.Producto, error) {
	productoStore, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[ProductsService][Patch] error getting product by ID", err)
		return domain.Producto{}, err
	}

	productoPatch, err := s.validatePatch(productoStore, producto)
	if err != nil {
		log.Println("[ProductsService][Patch] error validating product", err)
		return domain.Producto{}, err
	}

	producto, err = s.repository.Patch(ctx, productoPatch, id)
	if err != nil {
		log.Println("[ProductsService][Patch] error patching product by ID", err)
		return domain.Producto{}, err
	}

	return producto, nil
}

// validatePatch is a method that validates the fields to be updated.
func (s *service) validatePatch(productoStore, producto domain.Producto) (domain.Producto, error) {

	if producto.Name != "" {
		productoStore.Name = producto.Name
	}

	if producto.Quantity != 0 {
		productoStore.Quantity = producto.Quantity
	}

	if producto.CodeValue != "" {
		productoStore.CodeValue = producto.CodeValue
	}

	if !producto.Expiration.Equal(time.Time{}) {
		productoStore.Expiration = producto.Expiration
	}

	if producto.Price != 0 {
		productoStore.Price = producto.Price
	}

	return productoStore, nil

}

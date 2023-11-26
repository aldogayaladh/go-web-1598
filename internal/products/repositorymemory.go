package products

import (
	"context"
	"errors"
	"log"

	"github.com/aldogayaladh/go-web-1598/internal/domain"
)

var (
	ErrEmpty    = errors.New("empty list")
	ErrNotFound = errors.New("product not found")
)

type repository struct {
	db []domain.Producto
}

// NewMemoryRepository ....
func NewMemoryRepository(db []domain.Producto) Repository {
	return &repository{db: db}
}

// Create is a method that creates a new product.
func (r *repository) Create(ctx context.Context, producto domain.Producto) (domain.Producto, error) {
	r.db = append(r.db, producto)
	return producto, nil
}

// GetAll is a method that returns all products.
func (r *repository) GetAll(ctx context.Context) ([]domain.Producto, error) {

	contenidoContext := ctx.Value("rol")

	if contenidoContext != "" {
		log.Println("El contenido del contexto es:", contenidoContext)
	}

	if len(r.db) < 1 {
		return []domain.Producto{}, ErrEmpty
	}

	return r.db, nil
}

// GetByID is a method that returns a product by ID.
func (r *repository) GetByID(ctx context.Context, id int) (domain.Producto, error) {
	var result domain.Producto
	for _, value := range r.db {
		if value.Id == id {
			result = value
			break
		}
	}

	if result.Id < 1 {
		return domain.Producto{}, ErrNotFound
	}

	return result, nil
}

// Update is a method that updates a product by ID.
func (r *repository) Update(
	ctx context.Context,
	producto domain.Producto,
	id int) (domain.Producto, error) {

	var result domain.Producto
	for key, value := range r.db {
		if value.Id == id {
			producto.Id = id
			r.db[key] = producto
			result = r.db[key]
			break
		}
	}

	if result.Id < 1 {
		return domain.Producto{}, ErrNotFound
	}

	return result, nil

}

// Delete is a method that deletes a product by ID.
func (r *repository) Delete(ctx context.Context, id int) error {
	var result domain.Producto
	for key, value := range r.db {
		if value.Id == id {
			result = r.db[key]
			r.db = append(r.db[:key], r.db[key+1:]...)
			break
		}
	}

	if result.Id < 1 {
		return ErrNotFound
	}

	return nil
}

// Patch is a method that updates a product by ID.
func (r *repository) Patch(
	ctx context.Context,
	producto domain.Producto,
	id int) (domain.Producto, error) {

	var result domain.Producto
	for key, value := range r.db {
		if value.Id == id {
			if producto.Name != "" {
				r.db[key].Name = producto.Name
			}
			if producto.CodeValue != "" {
				r.db[key].CodeValue = producto.CodeValue
			}
			if producto.Quantity > 0 {
				r.db[key].Quantity = producto.Quantity
			}
			if producto.Price > 0 {
				r.db[key].Price = producto.Price
			}
			if producto.Expiration != (domain.Producto{}).Expiration {
				r.db[key].Expiration = producto.Expiration
			}
			if producto.IsPublished {
				r.db[key].IsPublished = producto.IsPublished
			}
			result = r.db[key]
			break
		}
	}

	if result.Id < 1 {
		return domain.Producto{}, ErrNotFound
	}

	return result, nil
}

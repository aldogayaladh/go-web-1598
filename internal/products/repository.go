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

type Repository interface {
	Create(ctx context.Context, producto domain.Producto) (domain.Producto, error)
	GetAll(ctx context.Context) ([]domain.Producto, error)
	GetByID(ctx context.Context, id string) (domain.Producto, error)
	Update(ctx context.Context, producto domain.Producto, id string) (domain.Producto, error)
	Delete(ctx context.Context, id string) error
}

type repository struct {
	db []domain.Producto
}

// NewMemoryRepository ....
func NewMemoryRepository(db []domain.Producto) Repository {
	return &repository{db: db}
}

// Create ....
func (r *repository) Create(ctx context.Context, producto domain.Producto) (domain.Producto, error) {
	r.db = append(r.db, producto)
	return producto, nil
}

// GetAll...
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

// GetByID .....
func (r *repository) GetByID(ctx context.Context, id string) (domain.Producto, error) {
	var result domain.Producto
	for _, value := range r.db {
		if value.Id == id {
			result = value
			break
		}
	}

	if result.Id == "" {
		return domain.Producto{}, ErrNotFound
	}

	return result, nil
}

// Update ...
func (r *repository) Update(
	ctx context.Context,
	producto domain.Producto,
	id string) (domain.Producto, error) {

	var result domain.Producto
	for key, value := range r.db {
		if value.Id == id {
			producto.Id = id
			r.db[key] = producto
			result = r.db[key]
			break
		}
	}

	if result.Id == "" {
		return domain.Producto{}, ErrNotFound
	}

	return result, nil

}

// Delete ...
func (r *repository) Delete(ctx context.Context, id string) error {
	var result domain.Producto
	for key, value := range r.db {
		if value.Id == id {
			result = r.db[key]
			r.db = append(r.db[:key], r.db[key+1:]...)
			break
		}
	}

	if result.Id == "" {
		return ErrNotFound
	}

	return nil
}

package products

import (
	"context"

	"github.com/aldogayaladh/go-web-1598/internal/domain"
)

type Repository interface {
	Create(ctx context.Context, producto domain.Producto) (domain.Producto, error)
	GetAll(ctx context.Context) ([]domain.Producto, error)
	GetByID(ctx context.Context, id int) (domain.Producto, error)
	Update(ctx context.Context, producto domain.Producto, id int) (domain.Producto, error)
	Delete(ctx context.Context, id int) error
}

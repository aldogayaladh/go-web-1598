package seller

import (
	"context"

	"github.com/aldogayaladh/go-web-1598/internal/domain"
)

type Repository interface {
	Create(ctx context.Context, seller domain.Seller) (domain.Seller, error)
	GetAll(ctx context.Context) ([]domain.Seller, error)
	GetByID(ctx context.Context, id int) (domain.Seller, error)
	Update(ctx context.Context, seller domain.Seller, id int) (domain.Seller, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, seller domain.Seller, id int) (domain.Seller, error)
}

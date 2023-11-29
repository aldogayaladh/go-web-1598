package sale

import (
	"context"

	"github.com/aldogayaladh/go-web-1598/internal/domain"
)

type Repository interface {
	Create(ctx context.Context, sale domain.Sale) (domain.Sale, error)
	GetAll(ctx context.Context) ([]domain.Sale, error)
	GetByID(ctx context.Context, id int) (domain.Sale, error)
	Update(ctx context.Context, sale domain.Sale, id int) (domain.Sale, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, sale domain.Sale, id int) (domain.Sale, error)
}

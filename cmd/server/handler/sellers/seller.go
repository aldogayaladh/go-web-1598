package producto

import (
	"net/http"

	"github.com/aldogayaladh/go-web-1598/internal/domain"
	seller "github.com/aldogayaladh/go-web-1598/internal/seller"
	"github.com/aldogayaladh/go-web-1598/pkg/web"
	"github.com/gin-gonic/gin"
)

type Controlador struct {
	service seller.Service
}

func NewControladorSeller(service seller.Service) *Controlador {
	return &Controlador{
		service: service,
	}
}

// Producto godoc
// @Summary seller example
// @Description Create a new seller
// @Tags seller
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /seller [post]
func (c *Controlador) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request domain.Seller

		err := ctx.Bind(&request)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		seller, err := c.service.Create(ctx, request)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, seller)

	}
}

// Implement the rest of the methods in the same way as the previous one.

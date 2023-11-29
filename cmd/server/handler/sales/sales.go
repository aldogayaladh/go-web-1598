package producto

import (
	"net/http"

	"github.com/aldogayaladh/go-web-1598/internal/domain"
	sale "github.com/aldogayaladh/go-web-1598/internal/sale"
	"github.com/aldogayaladh/go-web-1598/pkg/web"
	"github.com/gin-gonic/gin"
)

type Controlador struct {
	service sale.Service
}

func NewControladorSale(service sale.Service) *Controlador {
	return &Controlador{
		service: service,
	}
}

// Producto godoc
// @Summary sale example
// @Description Create a new sale
// @Tags sale
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /sale [post]
func (c *Controlador) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request domain.Sale

		err := ctx.Bind(&request)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		sale, err := c.service.Create(ctx, request)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, sale)

	}
}

// Implement the rest of the methods in the same way as the previous one.

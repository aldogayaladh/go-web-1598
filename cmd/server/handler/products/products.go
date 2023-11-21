package products

import (
	"context"
	"net/http"

	"github.com/aldogayaladh/go-web-1598/internal/domain"
	"github.com/aldogayaladh/go-web-1598/internal/products"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service products.Service
}

func NewControllerProducts(service products.Service) *Controller {
	return &Controller{service: service}
}

// Doc ...
func (c *Controller) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var productRequest domain.Producto

		err := ctx.Bind(&productRequest)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
				"error":   err,
			})
			return
		}

		producto, err := c.service.Create(ctx, productRequest)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": producto,
		})

	}
}

// Doc ...
func (c *Controller) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		newContext := addValueToContext(ctx)
		listProducts, err := c.service.GetAll(newContext)
		if err != nil {

			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": listProducts,
		})
	}
}

// Doc ...
func (c *Controller) HandlerGetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Recuperamos el id de la request
		idParam := ctx.Param("id")

		// Llamamos al servicio
		producto, err := c.service.GetByID(ctx, idParam)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": producto,
		})
	}
}

// Doc ...
func (c *Controller) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Recuperamos el id de la request
		idParam := ctx.Param("id")

		var productRequest domain.Producto

		err := ctx.Bind(&productRequest)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
				"error":   err,
			})
			return
		}

		// Llamamos al servicio
		producto, err := c.service.Update(ctx, productRequest, idParam)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": producto,
		})
	}
}

// Doc ...
func (c *Controller) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Recuperamos el id de la request
		idParam := ctx.Param("id")

		// Llamamos al servicio
		err := c.service.Delete(ctx, idParam)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Producto eliminado",
		})
	}
}

// addValueToContext ...
func addValueToContext(ctx context.Context) context.Context {
	newContext := context.WithValue(ctx, "rol", "admin")
	return newContext
}

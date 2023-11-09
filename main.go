package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Producto es una estructura que define ...
type Producto struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Quantity    int       `json:"quantity"`
	CodeValue   string    `json:"code_value"`
	IsPublished bool      `json:"is_published"`
	Expiration  time.Time `json:"expiration"`
	Price       float64   `json:"price"`
}

// store es una base de datos en memoria
type Store struct {
	Productos []Producto
}

func main() {

	// Carga la base de datos en memoria
	store := Store{}
	store.LoadStore()

	engine := gin.Default()

	group := engine.Group("/api/v1")
	{
		group.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"mensaje": "pong",
			})
		})

		grupoProducto := group.Group("/producto")
		{
			grupoProducto.GET("", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{
					"data": store.Productos,
				})
			})

			grupoProducto.GET("/search/:parametroPrecio", func(ctx *gin.Context) {

				precioParametro := ctx.Param("parametroPrecio")

				precioCasteado, err := strconv.ParseFloat(precioParametro, 64)
				if err != nil {
					ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"mensaje": "parametro invalido",
					})
					return
				}

				var result []Producto
				for _, producto := range store.Productos {
					if producto.Price > precioCasteado {
						result = append(result, producto)
					}
				}

				ctx.JSON(http.StatusOK, gin.H{
					"data": result,
				})

			})

		}

	}

	if err := engine.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}

// LoadStore carga la base de datos en memoria
func (s *Store) LoadStore() {
	s.Productos = []Producto{
		{
			Id:          "1",
			Name:        "Coco Cola",
			Quantity:    10,
			CodeValue:   "123456789",
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       10.5,
		},
		{
			Id:          "2",
			Name:        "Pepsito",
			Quantity:    10,
			CodeValue:   "123456789",
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       8.5,
		},
		{
			Id:          "3",
			Name:        "Fantastica",
			Quantity:    10,
			CodeValue:   "123456789",
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       5.5,
		},
	}
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	handlerPing "github.com/aldogayaladh/go-web-1598/cmd/server/handler/ping"
	handlerProducto "github.com/aldogayaladh/go-web-1598/cmd/server/handler/products"
	"github.com/aldogayaladh/go-web-1598/internal/domain"
	"github.com/aldogayaladh/go-web-1598/internal/products"
	"github.com/aldogayaladh/go-web-1598/pkg/middleware"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	// Cargar las variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Carga la base de datos en memoria
	//db := LoadStore()
	db := connectDB()

	// Ping.
	controllerPing := handlerPing.NewControllerPing()

	// Products.
	//repostory := products.NewMemoryRepository(db)
	repositoryMysl := products.NewMySqlRepository(db)
	service := products.NewServiceProduct(repositoryMysl)
	controllerProduct := handlerProducto.NewControllerProducts(service)

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middleware.Logger())

	group := engine.Group("/api/v1")
	{
		group.GET("/ping", controllerPing.HandlerPing())

		grupoProducto := group.Group("/producto")
		{
			grupoProducto.POST("", middleware.Authenticate(), controllerProduct.HandlerCreate())
			grupoProducto.GET("", middleware.Authenticate(), controllerProduct.HandlerGetAll())
			grupoProducto.GET("/:id", controllerProduct.HandlerGetByID())
			grupoProducto.PUT("/:id", middleware.Authenticate(), controllerProduct.HandlerUpdate())
			grupoProducto.DELETE("/:id", middleware.Authenticate(), controllerProduct.HandlerDelete())

		}

	}

	if err := engine.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}

// LoadStore carga la base de datos en memoria
func LoadStore() []domain.Producto {
	return []domain.Producto{
		{
			Id:          1,
			Name:        "Coco Cola",
			Quantity:    10,
			CodeValue:   "123456789",
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       10.5,
		},
		{
			Id:          2,
			Name:        "Pepsito",
			Quantity:    10,
			CodeValue:   "123456789",
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       8.5,
		},
		{
			Id:          3,
			Name:        "Fantastica",
			Quantity:    10,
			CodeValue:   "123456789",
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       5.5,
		},
	}
}

func connectDB() *sql.DB {
	var dbUsername, dbPassword, dbHost, dbPort, dbName string
	dbUsername = "root"
	dbPassword = ""
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "storage"

	// string de conexion
	// "username:password@tcp(host:puerto)/base_datos"
	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}

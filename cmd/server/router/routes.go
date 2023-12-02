package routes

import (
	"database/sql"

	"github.com/aldogayaladh/go-web-1598/pkg/middleware"

	"github.com/aldogayaladh/go-web-1598/cmd/server/handler/ping"
	handlerProducto "github.com/aldogayaladh/go-web-1598/cmd/server/handler/products"
	handlerSale "github.com/aldogayaladh/go-web-1598/cmd/server/handler/sales"
	handlerSeller "github.com/aldogayaladh/go-web-1598/cmd/server/handler/sellers"

	producto "github.com/aldogayaladh/go-web-1598/internal/products"
	"github.com/aldogayaladh/go-web-1598/internal/sale"
	seller "github.com/aldogayaladh/go-web-1598/internal/seller"
	"github.com/gin-gonic/gin"
)

// Router interface defines the methods that any router must implement.
type Router interface {
	MapRoutes()
}

// router is the Gin router.
type router struct {
	engine      *gin.Engine
	routerGroup *gin.RouterGroup
	db          *sql.DB
}

// NewRouter creates a new Gin router.
func NewRouter(engine *gin.Engine, db *sql.DB) Router {
	return &router{
		engine: engine,
		db:     db,
	}
}

// MapRoutes maps all routes.
func (r *router) MapRoutes() {
	r.setGroup()
	r.buildPingRoutes()
	r.buildProductRoutes()
	r.buildSellerRoutes()
	r.buildSaleRoutes()
}

// setGroup sets the router group.
func (r *router) setGroup() {
	r.routerGroup = r.engine.Group("/api/v1")
}

// buildProductRoutes maps all routes for the product domain.
func (r *router) buildProductRoutes() {
	// Create a new product controller.
	repository := producto.NewMySqlRepository(r.db)
	service := producto.NewServiceProduct(repository)
	controlador := handlerProducto.NewControladorProducto(service)

	grupoProducto := r.routerGroup.Group("/productos")
	{
		grupoProducto.POST("", middleware.Authenticate(), controlador.HandlerCreate())
		grupoProducto.GET("", middleware.Authenticate(), controlador.HandlerGetAll())
		grupoProducto.GET("/:id", controlador.HandlerGetByID())
		grupoProducto.PUT("/:id", middleware.Authenticate(), controlador.HandlerUpdate())
		grupoProducto.DELETE("/:id", middleware.Authenticate(), controlador.HandlerDelete())
		grupoProducto.PATCH("/:id", middleware.Authenticate(), controlador.HandlerPatch())

	}

}

func (r *router) buildSellerRoutes() {
	// Create a new seller controller.
	repository := seller.NewMySqlRepositorySeller(r.db)
	service := seller.NewServiceSeller(repository)
	controlador := handlerSeller.NewControladorSeller(service)

	grupoSeller := r.routerGroup.Group("/seller")
	{
		grupoSeller.POST("", middleware.Authenticate(), controlador.HandlerCreate())
		// Add the rest of the routes.
	}
}

func (r *router) buildSaleRoutes() {
	// Create a new sale controller.
	repository := sale.NewMySqlRepositorySale(r.db)
	service := sale.NewServiceSale(repository)
	controlador := handlerSale.NewControladorSale(service)

	grupoSale := r.routerGroup.Group("/sale")
	{
		grupoSale.POST("", middleware.Authenticate(), controlador.HandlerCreate())
	}

	// Add the rest of the routes.
}

// buildPingRoutes maps all routes for the ping domain.
func (r *router) buildPingRoutes() {
	// Create a new ping controller.
	pingController := ping.NewControllerPing()
	r.routerGroup.GET("/ping", pingController.HandlerPing())

}

package routes

import (
	"database/sql"

	"github.com/aldogayaladh/go-web-1598/pkg/middleware"

	"github.com/aldogayaladh/go-web-1598/cmd/server/handler/ping"
	handlerProducto "github.com/aldogayaladh/go-web-1598/cmd/server/handler/products"

	producto "github.com/aldogayaladh/go-web-1598/internal/products"
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
	r.buildProductRoutes()
	r.buildPingRoutes()
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

	grupoProducto := r.routerGroup.Group("/producto")
	{
		grupoProducto.POST("", middleware.Authenticate(), controlador.HandlerCreate())
		grupoProducto.GET("", middleware.Authenticate(), controlador.HandlerGetAll())
		grupoProducto.GET("/:id", controlador.HandlerGetByID())
		grupoProducto.PUT("/:id", middleware.Authenticate(), controlador.HandlerUpdate())
		grupoProducto.DELETE("/:id", middleware.Authenticate(), controlador.HandlerDelete())
		grupoProducto.PATCH("/:id", middleware.Authenticate(), controlador.HandlerPatch())

	}

}

// buildPingRoutes maps all routes for the ping domain.
func (r *router) buildPingRoutes() {
	// Create a new ping controller.
	pingController := ping.NewControllerPing()
	r.routerGroup.GET("/ping", pingController.HandlerPing())

}

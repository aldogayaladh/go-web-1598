package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	routes "github.com/aldogayaladh/go-web-1598/cmd/server/router"
	"github.com/aldogayaladh/go-web-1598/pkg/eureka"
	"github.com/aldogayaladh/go-web-1598/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/aldogayaladh/go-web-1598/docs"
	_ "github.com/go-sql-driver/mysql"
)

const (
	puerto     = "9090"
	appID      = "service-product-1"
	appName    = "service-product"
	statusUp   = "UP"
	statusDown = "DOWN"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:9090
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

	// Recover from panic.
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}()

	// Load the environment variables.
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the database.
	db := connectDB()

	eureka.RegisterApp(appID, appName)
	time.Sleep(time.Second * 5)
	eureka.UpdateStatus(appID, appName, statusUp)

	task := eureka.ScheduleHeartbeat(appID, appName)

	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt)

	go func() {
		select {
		case sig := <-channel:
			_ = sig
			task.Cancel()
			eureka.UpdateStatus(appID, appName, statusDown)
			time.Sleep(time.Second * 5)
			eureka.DeleteApp(appName, appID)
			os.Exit(1)
		}

	}()

	// Create a new Gin engine.
	router := gin.New()
	router.Use(gin.Recovery())
	// Add the logger middleware.
	router.Use(middleware.Logger())

	// Add the swagger handler.
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run the application.
	runApp(db, router)

	// Close the connection.
	defer db.Close()

}

func runApp(db *sql.DB, engine *gin.Engine) {
	// Run the application.
	router := routes.NewRouter(engine, db)
	// Map all routes.
	router.MapRoutes()
	if err := engine.Run(fmt.Sprintf(":%s", puerto)); err != nil {
		panic(err)
	}

}

// connectDB connects to the database.
func connectDB() *sql.DB {
	var dbUsername, dbPassword, dbHost, dbPort, dbName string
	dbUsername = "root"
	dbPassword = ""
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "my_db"

	// Create the data source.
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)

	// Open the connection.
	db, err := sql.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}

	// Check the connection.
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	return db
}

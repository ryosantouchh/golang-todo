package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"golang.org/x/time/rate"
	"gorm.io/driver/sqlite"

	// "gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/ryosantouchh/golang-todo/auth"
	"github.com/ryosantouchh/golang-todo/todo"
)

var (
	buildcommit = "dev"
	buildtime   = time.Now().String()
)

func main() {
	// load env
	err := godotenv.Load("local.env")
	if err != nil {
		log.Printf("please consider environment variables: %s", err)
	}

	// open db connection
	// db, err := gorm.Open(mysql.Open(os.Getenv("DB_CONN")), &gorm.Config{})
	db, err := gorm.Open(sqlite.Open(os.Getenv("DB_CONN")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// auto migrate db
	db.AutoMigrate(&todo.Todo{})

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:8080",
	}
	config.AllowHeaders = []string{
		"Origin",
		"Authorization",
		"TransactionID",
	}
	r.Use(cors.New(config))

	r.GET("/healthz", func(c *gin.Context) {
		c.Status(200)
	})
	r.GET("/limitz", limitedHandler)
	r.GET("/x", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"buildcommit": buildcommit,
			"buildtime":   buildtime,
		})
	})

	// route.NewRouterHandler(r)

	// r.GET("/tokenz", auth.AccessToken) // normal function version
	r.GET("/tokenz", auth.AccessToken([]byte(os.Getenv("SIGN")))) // middleware version
	protected := r.Group("/", auth.Protect([]byte(os.Getenv("SIGN"))))
	// the signature is in the main.go -- easily read / protect

	gormStore := todo.NewGormStore(db)

	todoHandler := todo.NewTodoHandler(gormStore)
	protected.POST("/todos", todo.ConvertGinHandler(todoHandler.NewTask))
	// protected.GET("/todos", todoHandler.GetTodoList)
	// protected.DELETE("/todos/:id", todoHandler.DeleteTodo)

	r.Run()
}

var limiter = rate.NewLimiter(5, 5)

func limitedHandler(c *gin.Context) {
	if !limiter.Allow() {
		c.AbortWithStatus(http.StatusTooManyRequests)
		return
	}
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"stockexchange.com/bank"
	"stockexchange.com/config"

	"stockexchange.com/transaction"
)

var db *gorm.DB = config.SetupDatabaseConnection()


func initDB() {

	bank.DB = db
	transaction.DB = db

}

func GinMiddlewareForSocketIO() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Request.Header.Del("Origin")

		c.Next()
	}
}

func NewServer() *gin.Engine {
	router := gin.Default()

	CORSHandler := cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
	})
	router.Use(CORSHandler)
	router.Use(GinMiddlewareForSocketIO())

	superRouter := router.Group("/stock.com")
	{

		bank.GroupApi(superRouter)
		transaction.GroupApi(superRouter)
	
	}
	// socket := socket.NewSocketServer()

	return router
}

func main() {
	initDB()
	server := NewServer()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)
	defer config.CloseDatabaseConnection(db)


}




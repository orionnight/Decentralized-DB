package router

import (
	"net/http"

	"example.com/ece1770/controller/database"
	"example.com/ece1770/controller/user"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.Use(gin.Recovery())

	corsMiddleware := func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "User-Agent, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", http.MethodPost)

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
	r.Use(corsMiddleware)

	userRoute := r.Group("/user")
	{
		userRoute.POST("/create", user.CreateKeystore)
		userRoute.POST("/login", user.Login)
	}

	// Add routing group for DB operations
	dbRoute := r.Group("/db")
	{
		dbRoute.GET("/ping", database.Ping)
		dbRoute.POST("/create", database.HandleDBCreation)
		dbRoute.GET("/find", database.HandleDBFind)
		dbRoute.GET("/update", database.HandleDBUpdate)
		dbRoute.GET("/delete", database.HandleDBDeletion)
	}

	return r
}

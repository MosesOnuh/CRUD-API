package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", helloWorldhandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	_ = router.Run(":" + port)

}

func helloWorldhandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Guys",
	})
}

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/Nice", helloWorldhandler)

	_ = router.Run(":3000")
}

func helloWorldhandler(c *gin.Context){
	c.JSON(200, gin.H{
		"message":  "hello world",
	})
}

package main

import (
	"os"
	"github.com/gin-gonic/gin"
)



// multipleUsers := [User{Micheal, James, Timmy}]
// Micheal

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
	Country string `json:"country"`
}

func main() {
	router := gin.Default()
	router.GET("/getUsers", getHandler)
	router.POST("/createUsers", createHandler)
	router.PATCH("/updateUsers", updateHandler)
	router.DELETE("/delete", deleteHandler)
	

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	_ = router.Run(":" + port)

}

func getHandler(c *gin.Context) {
	var user User
	user = User{
		Name: "John Doe",
		Age: 39,
		Email: "johndoe@gmail.com",
		Country: "Cameroon",
	}
	c.JSON(200, gin.H{
		"message": "Hello Guys",
		"data": user,
	})
}

func createHandler(c *gin.Context) {
	var user User

	err := c.ShouldBindJSON(&user)
	if err != nil {
	c.JSON(400, gin.H{
			"error": "invalid request data",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Profile successfully created",
		"data": user,
	})
}

func updateHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Profile Updated",
	})
}

func deleteHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Good People",
	})
}
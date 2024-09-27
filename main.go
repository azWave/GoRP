package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type user struct {
	Id	int `json:"id"`
	Username string `json:"username"`
}

var users = []user{
	{Id:123 , Username :"Simon"},
	{Id:124 , Username :"Maxime"},
}



func getUsers(c *gin.Context){
	c.IndentedJSON(http.StatusOK, users)
}

func main () {
	router := gin.Default()
	router.GET("/users",getUsers)

	router.Run("localhost:8080")

	
}

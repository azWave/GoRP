package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Roles struct {
	Name string `json:"name"`
	Health int `json:"health"`
	Mana int `json:"mana"`
	Stength int `json:"strength"`
	Armor int `json:"armor"`
}

type player struct {
	Username string `json:"username"`
	Role Roles `json:"role"`
	Health int `json:"health"`
	Mana int `json:"mana"`
	Stength int `json:"strength"`
	Armor int `json:"armor"`	
}




var roles = []Roles{
	{Name : "Warrior",Health : 150,Mana : 50,Stength : 200,Armor:150},
	{Name : "Thief",Health : 100,Mana : 75,Stength : 100,Armor:75},
	{Name : "Mage",Health : 120,Mana : 200,Stength : 75,Armor:100},
}



func getRoles(ctx *gin.Context){
	ctx.IndentedJSON(http.StatusOK, roles)
}

func main () {
	router := gin.Default()
	router.GET("/roles",getRoles)

	router.Run("localhost:8080")

	
}

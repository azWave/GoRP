package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Roles struct {
	Name    string `json:"name"`
	Health  int    `json:"health"`
	Mana    int    `json:"mana"`
	Stength int    `json:"strength"`
	Armor   int    `json:"armor"`
}

type Player struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     Roles  `json:"role"`
	Health   int    `json:"health"`
	Mana     int    `json:"mana"`
	Stength  int    `json:"strength"`
	Armor    int    `json:"armor"`
}

var roles = []Roles{
	{Name: "Warrior", Health: 150, Mana: 50, Stength: 200, Armor: 150},
	{Name: "Thief", Health: 100, Mana: 75, Stength: 100, Armor: 75},
	{Name: "Mage", Health: 120, Mana: 200, Stength: 75, Armor: 100},
}

var players = []Player{}

func getRoles(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, roles)
}

func startGame(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Bonjour dans ce jeu choisissez un nom d'utilisateur et un role\nvoici les roles que vous pouvez choisir : \n")
	ctx.IndentedJSON(http.StatusOK, roles)
	ctx.String(http.StatusOK, "\nPour choisir un role, faites une requete POST sur /create-player avec le role choisi et votre nom d'utilisateur\n")
	response := map[string]string{
		"username": "Votre nom d'utilisateur",
		"role":     "Nom du role",
	}
	ctx.IndentedJSON(http.StatusOK, response)
}

// /start-game GET -> return infos about the game
// /create-player POST -> create a player -> return created player

func main() {
	router := gin.Default()
	router.GET("/roles", getRoles)
	router.GET("/start-game", startGame)

	router.Run("localhost:8080")

}

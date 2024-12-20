package input

import (
	"fmt"
	"gorp/core/domain/entities"
	"gorp/core/domain/interfaces"
	"gorp/core/services"
)

func MapCreationHandler(printer interfaces.Printer, service *services.MapService) {
	var width, height int

	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
	fmt.Println("Création d'une carte")
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
	fmt.Println("Entrez la largeur de la carte :")
	fmt.Scanln(&width)

	fmt.Println("Entrez la hauteur de la carte :")
	fmt.Scanln(&height)

	service.InitializeMap(width, height)
	// 10% obstacle from the total number of tiles
	obstaclesCount := int(float64(width*height) * 0.1)
	service.AddRandomObstacles(obstaclesCount)

	service.SaveMap("world-1")

	PrintMap(printer, service.GameMap, entities.Position{X: 0, Y: 0, MapName: "world-1"})
}

func PrintMap(printer interfaces.Printer, gameMap *entities.Map, playerPosition entities.Position) {
	for y := 0; y < gameMap.Height; y++ {
		for x := 0; x < gameMap.Width; x++ {
			if x == playerPosition.X && y == playerPosition.Y {
				printer.Print("P")
			} else {
				tile, _ := gameMap.GetTile(x, y)
				printer.Print(tile.DisplayChar)
			}
		}
		printer.Println()
	}
}

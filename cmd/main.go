package main

import (
	"gorp/core/services"
	"gorp/infra/input"
	"gorp/infra/output"
)

func main() {
	repo := &output.FileRepository{BasePath: "data/"}
	characterService := &services.CharacterService{Repo: repo}
	input.CharacterCreationHandler(characterService)
}
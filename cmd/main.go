package main

import (
	"gorp/core/services"
	"gorp/infra/input"
	"gorp/infra/output"
)

func main() {
	repo := &output.FileRepository{BasePath: "tmp/data/"}
	characterService := &services.CharacterService{Repo: repo}
	printer := &output.FmtPrinter{}
	input.CharacterCreationHandler(printer, characterService)
}

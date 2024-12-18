package input_test

import (
	"bytes"
	"gorp/core/domain/entities"
	"gorp/infra/input"
	"io"
	"os"
	"testing"
)

func TestPrintCharacterSummary(t *testing.T) {
	// Configuration du personnage
	character := entities.Character{
		Name:  "Aragorn",
		Class: "Guerrier",
		Stats: entities.ClassStats["Guerrier"],
	}

	// Création d'un pipe pour rediriger la sortie standard
	r, w, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = w // Redirection de la sortie standard
	defer func() {
		os.Stdout = stdout // Restauration de la sortie standard
		r.Close()
		w.Close()
	}()

	// Buffer pour capturer la sortie
	var buffer bytes.Buffer
	done := make(chan bool) // Canal pour synchronisation

	go func() {
		io.Copy(&buffer, r) // Copie de la sortie standard redirigée vers le buffer
		done <- true
	}()

	// Appel de la fonction à tester
	input.PrintCharacterSummary(character)

	w.Close() // Ferme l'écriture pour permettre à io.Copy de terminer
	<-done    // Attente que la copie soit terminée

	// Lecture de la sortie capturée
	actualOutput := buffer.String()

	// Sortie attendue
	expectedOutput := "Nom : Aragorn\n" +
		"Classe : Guerrier\n" +
		"PV : 150\n" +
		"PM : 50\n" +
		"Force : 15\n" +
		"Intelligence : 5\n" +
		"Défense : 12\n" +
		"Résistance Magique : 6\n" +
		"Agilité : 8\n" +
		"Chance : 5\n" +
		"Endurance : 10\n" +
		"Esprit : 4\n"

	// Comparaison de la sortie
	if actualOutput != expectedOutput {
		t.Errorf("PrintCharacterSummary() = %v, want %v", actualOutput, expectedOutput)
	}
}

package pkg

import (
	"bytes"
	"io"
	"os"
)

func CaptureFunctionConsoleOutput(f func()) (string, error) {
	// Création d'un pipe pour rediriger la sortie standard
	r, w, err := os.Pipe()
	if err != nil {
		return "", err
	}

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

	// Appel de la fonction à capturer
	f()

	w.Close() // Ferme l'écriture pour permettre à io.Copy de terminer
	<-done    // Attente que la copie soit terminée

	return buffer.String(), nil
}

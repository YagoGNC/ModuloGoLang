package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Nombre vacio")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"¡Hola, %v! ¡Bienvenido!",
		"¡Què bueo verte, %v!",
		"¡Saludo, %v! ¡Encantado de conocerte!",
	}
	return formats[rand.Intn(len(formats))]
}

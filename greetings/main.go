package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("¡Bienvenido %v!", name)
	return message
}

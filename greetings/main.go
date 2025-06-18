package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("Nombre vacio")
	}
	message := fmt.Sprint(randomFormat())
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	//Creamos un map con make para almacenar el mensaje y el nombre
	messages := make(map[string]string)

	//Recorremos nombre por nombre el slice
	for _, name := range names {
		//en caso de lograrlo, lo almacenamos en message
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		//agarramos el map, le pasamos como clave nombre y como valor el mensaje
		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"¡Hola, %v! ¡Bienvenido!",
		"¡Què bueo verte, %v!",
		"¡Saludo, %v! ¡Encantado de conocerte!",
	}
	return formats[rand.Intn(len(formats))]
}

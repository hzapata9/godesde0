package ejer_interfaces

import (
	"fmt"

	"github.com/hzapata9/godesde0/interfaces"
)

func HumanoRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", hu.Sexo())
}

func HumanoEstaVivo(hu interfaces.SerVivo) {
	if hu.EstaVivo() == true {
		fmt.Println("Es un/a Ser Vivo")
	}
}

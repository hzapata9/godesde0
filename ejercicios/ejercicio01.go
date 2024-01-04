package ejercicios

import (
	"log"
	"strconv"
)

func ReturnTwoValues(param string) (valor int, salida string) {

	valor, err := strconv.Atoi(param)

	if err != nil {
		log.Fatal(err)
	} else {
		if valor > 100 {
			salida = "Es Mayor a 100"
		} else {
			salida = "Es Menor a 100"
		}
	}
	return valor, salida
}

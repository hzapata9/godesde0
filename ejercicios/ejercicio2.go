package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func IngresarNumero() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese número 1 : ")
	if scanner.Scan() {
		numero1, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto " + err.Error())
			IngresarNumero()
		} else {
			for indice := 1; indice < 11; indice++ {
				fmt.Println("Tabla númerica ", indice, " : ", indice*numero1)
			}
		}
	}

}

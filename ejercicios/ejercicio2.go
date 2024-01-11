package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var err error
var texto string

func IngresarNumero() string {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Ingrese número 1 : ")
		if scanner.Scan() {
			numero1, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for indice := 1; indice < 11; indice++ {
		texto += fmt.Sprintln("Tabla númerica ", indice, " : ", indice*numero1)
	}
	return texto
}

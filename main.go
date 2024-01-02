package main

import (
	"fmt"

	"github.com/hzapata9/godesde0/variables"
)

func main() {
	variables.MuestraEntyeros()

	variables.RestoVariables()

	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
}

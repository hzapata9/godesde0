package main

import (
	"fmt"

	"github.com/hzapata9/godesde0/ejercicios"
)

/*import (
	"fmt"
	"runtime"

	"github.com/hzapata9/godesde0/variables"
)

func main() {
	variables.MuestraEntyeros()

	variables.RestoVariables()

	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)

	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Es Linux")
	} else {
		fmt.Println("Es otro: ", os)
	}
}
*/

func main() {

	numero, texto := ejercicios.ReturnTwoValues("10")
	fmt.Println(numero)
	fmt.Println(texto)
}

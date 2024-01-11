package defer_panic

import (
	"fmt"
	"log"
)

func VamosDefer() {
	fmt.Println("Primer mensaje")
	defer fmt.Println("mensaje final")

	fmt.Println("segundo mensaje..")
}

func EjemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("ocurrio un error que genero Panic \n %v", reco)
		}
	}()

	a := 1
	if a == 1 {
		panic("Se encontro el valor 1")
	}
}

package defer_panic

import (
	"fmt"
)

func VamosDefer() {
	fmt.Println("Primer mensaje")
	defer fmt.Println("mensaje final")

	fmt.Println("segundo mensjae..")
}

package files

import (
	"fmt"
	"os"

	"github.com/hzapata9/godesde0/ejercicios"
)

var filename string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.IngresarNumero()
	archivo, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error al crear el archivo " + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.IngresarNumero()
	if Append(filename, texto) == false {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(filename, os.O_WRONLY || os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el Append " + err.Error())
		return false
	}

	 , err = arch.WriteString(texto)
	 if err != nil {
		fmt.Println("Error durante el WriteString")
		return false
	 }
	 arch.Close()
	 return true
}

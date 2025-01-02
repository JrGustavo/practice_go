package main

import (
	e "github.com/JrGustavo/practice_go/ejer_interfaces"
	"github.com/JrGustavo/practice_go/modelos"
)

func main() {
	//	//estado, texto := variables.ConviertoATexto(1588)
	//	//fmt.Println(estado)
	//	//fmt.Println(texto)
	//
	//	if os := runtime.GOOS; os == "Linux" || os == "OS X." {
	//		fmt.Println("Esto no es Windows, es ", os)
	//	} else {
	//		fmt.Println("Esto es Windows")
	//	}
	//	switch os := runtime.GOOS; os {
	//	case "linux":
	//		fmt.Println("Esto es Linux")
	//	case "darwin":
	//		fmt.Println("Esto es Darwin")
	//	default:
	//		fmt.Printf("%s \n", os)
	//	}
	//numero, texto := ejercicios.ConNumerico("pablito")
	//fmt.Print(numero)
	//fmt.Print(texto)

	//teclado.IngresoNumeros()

	//fmt.Print(ejercicios.TablaMultiplicar())

	//files.GrabaTabla()

	//files.LeoArchivo()

	//funciones.Calculos()

	//funciones.LlamarClosure()

	//arreglos_slices.MuestroArreglo()

	//mapas.MostrarMapas()

	Junior := new(modelos.Hombre)
	e.HumanosRespirando(Junior)

	Maria := new(modelos.Mujer)
	e.HumanosRespirando(Maria)

}

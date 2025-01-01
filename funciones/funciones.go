package funciones

import "fmt"

//Esto se conoce como funciones anonimas

func Calculos() {

	var numero4 int = 32
	var numero3 int = 243

	calculo := func(numero1 int, numero2 int) int {
		return numero1 + numero2 + numero4 + numero3
	}

	fmt.Println(calculo(10, 25))

	calculo = func(numero1 int, numero2 int) int {
		return numero1 * numero2 * numero4

	}

	fmt.Println(calculo(10, 25))

}

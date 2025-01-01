package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaMultiplicar() string {
	scanner := bufio.NewScanner(os.Stdin)

	var numero int
	var err error
	var texto string

	for {
		fmt.Print("Ingrese un número: ")
		if !scanner.Scan() {
			fmt.Println("Entrada finalizada.")

		}
		numero, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Por favor, ingrese un número válido.")
			continue
		}
		break
	}

	if scanner.Err() != nil {
		fmt.Println("Error al leer entrada:", scanner.Err())

	}

	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d\n", numero, i, numero*i)
	}
	fmt.Println(texto)

	return texto
}

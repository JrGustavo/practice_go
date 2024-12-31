package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TabladeMultiplicar() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Ingrese un numero: ")
		if scanner.Scan() {
			numero, err := strconv.Atoi(scanner.Text())
			if err != nil {

				continue
			}
		}
	}
}

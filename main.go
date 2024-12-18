package main

import (
	"fmt"
	"github.com/JrGustavo/practice_go/variables"
)

func main() {
	estado, texto := variables.ConviertoATexto(1588)
	fmt.Print(estado)
	fmt.Print(texto)

}

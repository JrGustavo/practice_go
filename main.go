package main

import (
	"fmt"
	"github.com/JrGustavo/practice_go/variables"
	"runtime"
)

func main() {
	estado, texto := variables.ConviertoATexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
	os := runtime.GOOS
	if os == "Linux." {

	} else {
		fmt.Println("No es Linux")
	}

}

package defer_panic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("Este es el mensaje final")

	fmt.Println("Este es un segundo mensaje")

}

func EjemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error que genero Panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Se encontro el valor de 1")
	}
}

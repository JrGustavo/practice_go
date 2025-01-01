package funciones

import "fmt"

// Tabla retorna una función que genera la secuencia multiplicando por un valor
func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}

// LlamarClosure imprime la tabla del número dado usando un closure
func LlamarClosure() {
	tablaDel := 2
	MiTabla := Tabla(tablaDel)
	for i := 1; i <= 10; i++ {
		fmt.Println(MiTabla())
	}
}

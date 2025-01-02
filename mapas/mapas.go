package mapas

import "fmt"

func MostrarMapas() {

	paises := make(map[string]string)
	//fmt.Println(paises)

	paises["Colombia"] = "Bogotá"
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises)

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  40,
		"Chivas":       37,
		"Boca Juniors": 30,
	}
	fmt.Println(campeonato)

	//for equipo, puntaje := range campeonato {
	//	fmt.Printf("El equipo %s, tiene un puntaje de: %d \n", equipo, puntaje)
	//}
	//
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Juventus"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe = %t \n", puntaje, existe)
}

package arreglos_slices

import "fmt"

var tablaS []int

func MuestroSlice() {
	fmt.Println(tablaS)

	porcion := tablaS[3:]   // Slice creado desde un vector desde la posicion 3
	porcion2 := tablaS[:5]  // Slice creado desde un vector desde la posicion 0 hasta la 5
	porcion3 := tablaS[2:4] // Slice creado desde un vector desde la posicion 2 hasta la 4

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %v, Capacidad %v", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nLargo %v, Capacidad %v", len(nums), cap(nums))

}

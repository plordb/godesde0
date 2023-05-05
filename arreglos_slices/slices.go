package arreglosslices

import "fmt"

var tablaS []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 98, 21, 674, 18, 36, 78, 9}

func MuestroSlice() {
	fmt.Println(tablaS)

	porcion := arreglo[3:]   // Slice creado desde un vector desde la posicion 3
	porcion2 := arreglo[:5]  // Slice creado desde la posicion 0 hasta la posicion 5
	porcion3 := arreglo[6:7] // Slice creado desde la posicion 6 hasta la posicin 7

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d \n", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("Largo(nums) %d, Capacidad(nums) %d \n", len(nums), cap(nums))
}

package arreglosslices

import (
	"fmt"
)

var tabla [10]int

func MuestroArreglos() {
	tabla[7] = 33
	tabla[2] = 54

	fmt.Println(tabla)
}

package middleware

import (
	"fmt"
)

func sumar(a int, b int) int {
	return a + b
}

func restar(a int, b int) int {
	return a - b
}

func multiplicar(a int, b int) int {
	return a * b
}

func MiMiddleware() {
	fmt.Println("Inicio")

	result := operacionesMidd(sumar)(2, 4)
	fmt.Println(result)

	result = operacionesMidd(restar)(2, 4)
	fmt.Println(result)

	result = operacionesMidd(multiplicar)(2, 4)
	fmt.Println(result)
}

func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de Operación")
		return f(a, b)
	}
}

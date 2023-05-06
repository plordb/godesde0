package main

import (
	"fmt"

	"github.com/plordb/godesde0/goroutines"
)

// 02-25

func main() {

	//fmt.Println(ejercicios.TablaMultiplicar())
	//files.GrabaTabla()
	//files.SumaTabla()
	//files.LeoArchivo()
	//funciones.Calculos()
	//funciones.LlamarClosure()
	//funciones.Exponencia(2)

	//arreglosslices.MuestroSlice()
	//arreglosslices.Capacidad()
	//mapas.MostrarMapas()
	//users.AltaUsuario()

	/*Pedro := new(modelos.Hombre)
	e.HumanosRespirando(Pedro)

	Maria := new(modelos.Mujer)
	e.HumanosRespirando(Maria)*/

	//d.VemosDefer()
	//d.EjemploPanic()

	go goroutines.MiNombreLento("Pablo Lorenzo")

	fmt.Println("Estoy aqui")
	var x string
	fmt.Scanln(&x)
}

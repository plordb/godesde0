package main

import (
	e "github.com/plordb/godesde0/ejer_interfaces"
	"github.com/plordb/godesde0/modelos"
)

// 02-23

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

	Pedro := new(modelos.Hombre)
	e.HumanosRespirando(Pedro)

	Maria := new(modelos.Mujer)
	e.HumanosRespirando(Maria)
}

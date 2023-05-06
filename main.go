package main

import (
	"github.com/plordb/godesde0/webserver"
)

// 02-26

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

	/*canal1 := make(chan bool)
	go goroutines.MiNombreLento("Pablo Lorenzo", canal1)

	defer func() {
		<-canal1
	}()

	fmt.Println("Estoy aqui")*/

	webserver.MiWebServer()

}

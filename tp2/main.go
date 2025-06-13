package tp2

import (
	"bufio"
	"os"
	"strings"
	"tdas/diccionario"
)

func comparacion_fechas(fecha1 string,fecha2 string){
	//..
}


func agregar_archivo(ruta string){
	abb := diccionario.CrearABB[string,Vuelo](comparacion_fechas)

	archivo, _ := os.Open(ruta)
	defer archivo.Close()
	lectura := bufio.NewScanner(archivo)
	for lectura.Scan() {
		linea := lectura.Text()
		linea_sep := strings.Split(linea,",")
		vuelo := CrearVuelo(linea_sep)
		abb.Guardar(vuelo.Obtener_fecha(),vuelo)
	}
}

func main(){
	args := os.Args
	switch args[1]{
	case "agregar_archivo":
		agregar_archivo(args[2])

	// case "ver_tablero":
	// 	printVuelos(args[2])	
	}
}

// func printVuelos(ruta string){
// 	archivo, _ := os.Open(ruta)
// 	defer archivo.Close()
// 	lectura := bufio.NewScanner(archivo)
// 	for lectura.Scan() {
// 		vuelo := lectura.Text()
// 		vuelo_sep := strings.Split(vuelo,",")
// 		vuelos_juntos := strings.Join(vuelo_sep, "	")
// 		fmt.Println(vuelos_juntos)
// 	}
// }

// func print_vuelo(vuelo *vuelo){
// 	vuelos_juntos := strings.Join(vuelo_sep, "	")
// }
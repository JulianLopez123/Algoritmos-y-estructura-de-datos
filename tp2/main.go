package tp2

import (
	"bufio"
	"os"
	"strings"
	"tdas/diccionario"
)


func comparacion_fechas_ascendente(fecha1 string,fecha2 string) int{
	if fecha1 > fecha2{
		return 1
	}else if fecha1 < fecha2{
		return -1 
	}
	return 0
}


func agregar_archivo(ruta string,hash *diccionario.Diccionario){
	
	archivo, _ := os.Open(ruta)
	defer archivo.Close()
	lectura := bufio.NewScanner(archivo)
	for lectura.Scan() {
		linea := lectura.Text()
		linea_sep := strings.Split(linea,",")
		vuelo := CrearVuelo(linea_sep)
		hash.Guardar(vuelo.Obtener_numero_vuelo(),&vuelo)
	}
}

func main(){
	args := os.Args

	hash := diccionario.CrearHash[int,*Vuelo]()
	abb := diccionario.CrearABB[string,*Vuelo](comparacion_fechas_ascendente)

	switch args[1]{
	case "agregar_archivo":
		agregar_archivo(args[2],hash)

	// case "ver_tablero":
	// 	printVuelos(args[2])	
	}
}

// func printVuelos(ruta string){
// 
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
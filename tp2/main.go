package tp2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tdas/diccionario"
)


type vuelo struct {
	numero_vuelo int
	aerolinea string
	aeropuerto_origen string
	aeropuerto_destino string
	numero_cola string
	prioridad int
	fecha string
	retraso_salida int
	tiempo_vuelo int
	cancelado int
}

func crearVuelo(entrada []string) Vuelo{
	numero_vuelo,_ := strconv.Atoi(entrada[0])
	prioridad,_ := strconv.Atoi(entrada[5])
	retraso_salida,_ := strconv.Atoi(entrada[7])
	tiempo_vuelo,_ := strconv.Atoi(entrada[8])
	cancelado,_ := strconv.Atoi(entrada[9])
	
	return &vuelo{numero_vuelo: numero_vuelo,aerolinea:entrada[1],aeropuerto_origen: entrada[2],
	aeropuerto_destino: entrada[3],numero_cola: entrada[4],prioridad: prioridad,fecha: entrada[6],
	retraso_salida:retraso_salida,tiempo_vuelo: tiempo_vuelo,cancelado: cancelado,}
}

// func print_vuelo(vuelo *vuelo){
// 	vuelos_juntos := strings.Join(vuelo_sep, "	")
// }

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
		vuelo := crearVuelo(linea_sep)
		abb.Guardar(vuelo.fecha,vuelo)
	}
}

func main(){
	args := os.Args
	switch args[1]{
	case "ver_tablero":
		printVuelos(args[2])
	
	case "agregar_archivo":
		agregar_archivo(args[2])
	}

	
	
}


func printVuelos(ruta string){
	archivo, _ := os.Open(ruta)
	defer archivo.Close()
	lectura := bufio.NewScanner(archivo)
	for lectura.Scan() {
		vuelo := lectura.Text()
		vuelo_sep := strings.Split(vuelo,",")
		vuelos_juntos := strings.Join(vuelo_sep, "	")
		fmt.Println(vuelos_juntos)
	}
}
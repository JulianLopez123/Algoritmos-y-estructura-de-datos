package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tdas/diccionario"
	"tp2/TDAVuelo"
)


func comparacion_fechas_ascendente(fecha1 string,fecha2 string) int{
	if fecha1 > fecha2{
		return 1
	}else if fecha1 < fecha2{
		return -1 
	}
	return 0
}

func comparacion_fechas_descendente(fecha1,fecha2 string) int{
	return comparacion_fechas_ascendente(fecha2,fecha1)
}


// func agregar_archivo(ruta string){
// 	hash := diccionario.CrearHash[int,*TDAVuelo.Vuelo]()
// 	archivo, _ := os.Open(ruta)
// 	defer archivo.Close()
// 	lectura := bufio.NewScanner(archivo)
// 	for lectura.Scan() {
// 		linea := lectura.Text()
// 		linea_sep := strings.Split(linea,",")
// 		vuelo := TDAVuelo.CrearVuelo(linea_sep)
// 		hash.Guardar(vuelo.Obtener_numero_vuelo(),&vuelo)
// 	}
// }

func main() {
	lectura := bufio.NewScanner(os.Stdin)

	for {
		lectura.Scan()
		linea := lectura.Text()
		parametros := strings.Split(linea, " ")
		operacion := parametros[0]

		switch operacion {
		case "agregar_archivo":
			agregar_archivo(parametros[1])

		case "ver_tablero":
			cant_vuelos, _ := strconv.Atoi(parametros[1])
			ver_tablero(cant_vuelos, parametros[2], parametros[3], parametros[4])
		}
	}
	
}

func ver_tablero(cant_vuelos int,modo,desde,hasta string){
	if modo != "desc" && modo != "asc" || cant_vuelos <= 0 || comparacion_fechas_ascendente(desde,hasta) < 0{
		fmt.Println("Error en comando ver_tablero")
	}

	var abb diccionario.DiccionarioOrdenado[string, TDAVuelo.Vuelo]
	if modo == "desc"{
		abb = diccionario.CrearABB[string,TDAVuelo.Vuelo](comparacion_fechas_descendente)
	}else if modo == "asc"{
		abb = diccionario.CrearABB[string,TDAVuelo.Vuelo](comparacion_fechas_ascendente)
	}

	
	datos_almacenados, _ := os.Open("datos_almacenados.csv")
	defer datos_almacenados.Close()
	lectura := bufio.NewScanner(datos_almacenados)
	for lectura.Scan() {
		linea := lectura.Text()
		linea_sep := strings.Split(linea,",")
		vuelo := TDAVuelo.CrearVuelo(linea_sep)
		clave := fmt.Sprintf("%s - %s",vuelo.Obtener_fecha(),vuelo.Obtener_numero_vuelo())
		abb.Guardar(clave,vuelo)
	}

	// clave_desde := fmt.Sprintf("%s - 00000000",desde)//clave minima
	// clave_hasta := fmt.Sprintf("%s - 99999999",hasta)
	var contador int
	
	abb.IterarRango(&desde , &hasta, func(clave string, dato TDAVuelo.Vuelo)bool{
		if contador == cant_vuelos{
			return false
		}
		fmt.Println(clave)
		contador ++
		return true
	})
	fmt.Println("OK")
	
}


func agregar_archivo(ruta string) {
	hash := diccionario.CrearHash[int, TDAVuelo.Vuelo]()
	archivo, _ := os.Open(ruta)
	defer archivo.Close()
	lectura := bufio.NewScanner(archivo)
	for lectura.Scan() {
		linea := lectura.Text()
		linea_sep := strings.Split(linea, ",")
		vuelo := TDAVuelo.CrearVuelo(linea_sep)
		hash.Guardar(vuelo.Obtener_numero_vuelo(), vuelo)
	}
	printVuelosEnArchivo(ruta, hash)
}


func printVuelosEnArchivo(ruta string, hash diccionario.Diccionario[int, TDAVuelo.Vuelo]) {
	archivo, _ := os.Open(ruta)
	iterador := hash.Iterador()
	defer archivo.Close()
	write := bufio.NewWriter(archivo)
	for iterador.HaySiguiente() {
		_,vuelo := iterador.VerActual()
		linea := fmt.Sprintf("%d,%s,%s,%s,%s,%d,%s,%02d,%d,%d\n",
			vuelo.numero_vuelo,
			vuelo.aerolinea,
			vuelo.aeropuerto_origen,
			vuelo.aeropuerto_destino,
			vuelo.numero_cola,
			vuelo.prioridad,
			vuelo.fecha,
			vuelo.retraso_salida,
			vuelo.tiempo_vuelo,
			vuelo.cancelado,
		)
		write.WriteString(linea)
	}
	write.Flush()
}
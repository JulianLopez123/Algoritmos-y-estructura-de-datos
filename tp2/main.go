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

type tuple struct {
	fecha  string
	codigo string
}

func comparacion_fechas_ascendente(fecha1 tuple, fecha2 tuple) int {
	result := strings.Compare(fecha1.fecha, fecha2.fecha)
	if result == 0 {
		result = strings.Compare(fecha1.codigo, fecha2.codigo)
	}
	return result
}

// func comparacion_fechas_ascendente(fecha1 string, fecha2 string) int {
// 	if fecha1 > fecha2 {
// 		return 1
// 	} else if fecha1 < fecha2 {
// 		return -1
// 	}
// 	return 0
// }

// func agregar_archivo(ruta string){
// 	hash := diccionario.CrearHash[int,*TDAVuelo.Vuelo]()
// 	archivo, _ := os.Open(ruta)
// 	defer archivo.Close()
// 	lectura := bufio.NewScanner(archivo)
// 	for lectura.Scan() {
// 		linea := lectura.Text()
// 		linea_sep := strings.Split(linea,",")
// 		vuelo := TDAVuelo.CrearVuelo(linea_sep)
// 		hash.Guardar(vuelo.Numero_vuelo(),&vuelo)
// 	}
// }

func main() {
	lectura := bufio.NewScanner(os.Stdin)
	hash := diccionario.CrearHash[string, TDAVuelo.Vuelo]()
	abb := diccionario.CrearABB[tuple, TDAVuelo.Vuelo](comparacion_fechas_ascendente)
	for {
		lectura.Scan()
		linea := lectura.Text()
		parametros := strings.Split(linea, " ")
		operacion := parametros[0]

		switch operacion {
		case "agregar_archivo":
			if !agregar_archivo(parametros[1], hash, abb) {
				imprimirError(operacion)
			}
		case "ver_tablero":
			cant_vuelos, _ := strconv.Atoi(parametros[1])
			if !ver_tablero(cant_vuelos, parametros[2], parametros[3], parametros[4], abb) {
				imprimirError(operacion)
			}
		}
	}

}

func ver_tablero(cant_vuelos int, modo string, desde, hasta string, abb diccionario.DiccionarioOrdenado[tuple, TDAVuelo.Vuelo]) bool {
	if modo != "desc" && modo != "asc" || cant_vuelos <= 0 { //|| comparacion_fechas_ascendente(desde,hasta) < 0
		return false //con errores
	}

	// var abb diccionario.DiccionarioOrdenado[string, TDAVuelo.Vuelo]
	// if modo == "desc"{
	//
	//	}else if modo == "asc"{
	//		abb = diccionario.CrearABB[string,TDAVuelo.Vuelo](comparacion_fechas_ascendente)
	//	}
	// datos_almacenados, _ := os.Open("datos_almacenados.csv")
	// defer datos_almacenados.Close()
	// lectura := bufio.NewScanner(datos_almacenados)
	// for lectura.Scan() {
	// 	linea := lectura.Text()
	// 	linea_sep := strings.Split(linea, ",")
	// 	vuelo := TDAVuelo.CrearVuelo(linea_sep)
	// 	clave := fmt.Sprintf("%s - %d", vuelo.Fecha(), vuelo.Numero_vuelo())
	// 	abb.Guardar(clave, vuelo)
	// }

	// clave_desde := fmt.Sprintf("%s - 00000000",desde)//clave minima
	// clave_hasta := fmt.Sprintf("%s - 99999999",hasta)

	nose := tuple{fecha: desde, codigo: "0"}
	nose2 := tuple{fecha: hasta, codigo: "99999999999"}

	var contador int
	if modo == "desc" {
		//desde, hasta = hasta, desde
		nose, nose2 = nose2, nose
	}

	abb.IterarRango(&nose, &nose2, func(clave tuple, dato TDAVuelo.Vuelo) bool {
		if contador == cant_vuelos {
			return false
		}
		fmt.Println(clave)
		contador++
		return true
	})
	fmt.Println("OK")
	return true //sin errores
}

func agregar_archivo(ruta string, hash diccionario.Diccionario[string, TDAVuelo.Vuelo], abb diccionario.DiccionarioOrdenado[tuple, TDAVuelo.Vuelo]) bool {
	archivo, err := os.Open(ruta)
	if err != nil {
		return false
	}
	defer archivo.Close()
	lectura := bufio.NewScanner(archivo)
	for lectura.Scan() {
		linea := lectura.Text()
		linea_sep := strings.Split(linea, ",")
		vuelo := TDAVuelo.CrearVuelo(linea_sep)
		hash.Guardar(vuelo.Numero_vuelo(), vuelo)
		abb.Guardar(tuple{fecha: vuelo.Fecha(), codigo: vuelo.Numero_vuelo()}, vuelo)
	}
	fmt.Println("OK")
	return true
}

// func info_vuelo(hash diccionario.DiccionarioOrdenado[int, TDAVuelo.Vuelo], codigo_vuelo int) {
// 	if !hash.Pertenece(codigo_vuelo) {
// 		fmt.Println("Error en comando info_vuelo")
// 	}

// }

// func borrar(abb diccionario.DiccionarioOrdenado[string, TDAVuelo.Vuelo], desde, hasta string) {
// 	if comparacion_fechas_ascendente(desde, hasta) < 0 {
// 		fmt.Println("Error en el comando borrar")
// 	}

// 	// abb.IterarRango(&desde, &hasta, func(clave string, dato TDAVuelo.Vuelo) bool {

// 	// })
// 	abb.IteradorRango(&desde, &hasta)
// 	fmt.Println("OK")
// }

// func printVuelosEnArchivo(ruta string, hash diccionario.Diccionario[int, TDAVuelo.Vuelo]) {
// 	archivo, _ := os.Open(ruta)
// 	iterador := hash.Iterador()
// 	defer archivo.Close()
// 	write := bufio.NewWriter(archivo)
// 	for iterador.HaySiguiente() {
// 		_,vuelo := iterador.VerActual()
// 		linea := vuelo.Obtener_toda_info()
// 		write.WriteString(linea)
// 		iterador.Siguiente()
// 	}
// 	write.Flush()
// }

func imprimirError(comando string) {
	fmt.Fprintln(os.Stderr, "Error en comando", comando)
}

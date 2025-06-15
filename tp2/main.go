package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tdas/cola_prioridad"
	"tdas/diccionario"
	"tp2/TDAVuelo"
)

type tuple struct {
	fecha  string
	codigo string
}
type numVuelo_prioridad struct {
	numero_vuelo string
	prioridad    int
}

// func comparacion_fechas_ascendente(fecha1 string, fecha2 string) int {
// 	if fecha1 > fecha2 {
// 		return 1
// 	} else if fecha1 < fecha2 {
// 		return -1
// 	}
// 	return 0
// }

func comparacion_fechas_ascendente(fecha1 tuple, fecha2 tuple) int {
	result := strings.Compare(fecha1.fecha, fecha2.fecha)
	if result == 0 {
		result = strings.Compare(fecha1.codigo, fecha2.codigo)
	}
	return result
}

func comparar_numero_vuelo_prioridad(a, b numVuelo_prioridad) int { //heap maximos
	resultado := a.prioridad - b.prioridad
	if resultado == 0 {
		return strings.Compare(a.numero_vuelo, b.numero_vuelo)
	}
	return resultado
}

func main() {
	lectura := bufio.NewScanner(os.Stdin)
	hash := diccionario.CrearHash[string, TDAVuelo.Vuelo]()
	abb := diccionario.CrearABB[tuple, TDAVuelo.Vuelo](comparacion_fechas_ascendente)
	// hash := diccionario.CrearHash[int, TDAVuelo.Vuelo]()
	// abb := diccionario.CrearABB[string,TDAVuelo.Vuelo](comparacion_fechas_ascendente)

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
		case "info_vuelo":
			if !info_vuelo(parametros[1], hash) {
				imprimirError(operacion)
			}
		case "prioridad_vuelos":
			cant_vuelos, _ := strconv.Atoi(parametros[1])
			if !prioridad_vuelos(cant_vuelos, hash) {
				imprimirError(operacion)
			}
		case "siguiente_vuelo":
			if !siguiente_vuelo(parametros[1],parametros[2],parametros[3],abb){
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
		fmt.Println(clave.fecha, "-", clave.codigo)
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

func info_vuelo(numero_vuelo string, hash diccionario.Diccionario[string, TDAVuelo.Vuelo]) bool {
	if !hash.Pertenece(numero_vuelo) {
		return false
	}
	vuelo := hash.Obtener(numero_vuelo)
	fmt.Println(vuelo.Obtener_string())
	println("OK")
	return true
}

func prioridad_vuelos(cant_vuelos int, hash diccionario.Diccionario[string, TDAVuelo.Vuelo]) bool {
	heap := cola_prioridad.CrearHeap(comparar_numero_vuelo_prioridad) //heap maximos
	iterador := hash.Iterador()
	for iterador.HaySiguiente() {
		_, vuelo := iterador.VerActual()
		num_vuelo_prioridad := numVuelo_prioridad{numero_vuelo: vuelo.Numero_vuelo(), prioridad: vuelo.Prioridad()}
		heap.Encolar(num_vuelo_prioridad)
		iterador.Siguiente()
	}

	top := make([]numVuelo_prioridad, cant_vuelos)
	for i := 0; i < cant_vuelos; i++ {
		top[i] = heap.Desencolar()
	}
	for i := 0; i < cant_vuelos; i++ {
		numero_vuelo, prioridad := top[i].numero_vuelo, top[i].prioridad
		fmt.Println(prioridad, "-", numero_vuelo)
	}
	fmt.Println("OK")
	return true
}

func siguiente_vuelo(origen,destino,desde string, abb diccionario.DiccionarioOrdenado[tuple,TDAVuelo.Vuelo])bool{
	hallado := false
	fecha := tuple{fecha: desde, codigo: "0"}
	abb.IterarRango(&fecha, nil, func(clave tuple, vuelo TDAVuelo.Vuelo) bool {
		
		if origen == vuelo.Aeropuerto_origen() && destino == vuelo.Aeropuerto_destino() {
			fmt.Println(clave.fecha, "-", clave.codigo)
			hallado = true
			return false
		}
		return true
	})
	if !hallado{
		fmt.Println("No hay vuelo registrado desde",origen ,"hacia", destino,"desde",desde)
	}
	fmt.Println("OK")
	return true
}

func imprimirError(comando string) {
	fmt.Fprintln(os.Stderr, "Error en comando", comando)
}
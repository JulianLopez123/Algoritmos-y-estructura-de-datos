package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tdas/cola_prioridad"
	"tdas/diccionario"
	"tdas/pila"
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
		return -(strings.Compare(a.numero_vuelo, b.numero_vuelo))
	}
	return resultado
}

func main() {
	lectura := bufio.NewScanner(os.Stdin)
	hash := diccionario.CrearHash[string, TDAVuelo.Vuelo]()
	abb := diccionario.CrearABB[tuple, TDAVuelo.Vuelo](comparacion_fechas_ascendente)

	for {
		lectura.Scan()
		linea := lectura.Text()
		if linea == "" {
			break
		}
		parametros := strings.Split(linea, " ")
		operacion := parametros[0]

		switch operacion {
		case "agregar_archivo":
			if !verificarCantParametros(parametros, 2) {
				imprimirError(operacion)
				continue
			}
			if !agregar_archivo(parametros[1], hash, abb) {
				imprimirError(operacion)
			}
		case "ver_tablero":
			if !verificarCantParametros(parametros, 5) {
				imprimirError(operacion)
				continue
			}
			cant_vuelos, _ := strconv.Atoi(parametros[1])
			if !ver_tablero(cant_vuelos, parametros[2], parametros[3], parametros[4], abb) {
				imprimirError(operacion)
			}
		case "info_vuelo":
			if !verificarCantParametros(parametros, 2) {
				imprimirError(operacion)
				continue
			}
			if !info_vuelo(parametros[1], hash) {
				imprimirError(operacion)
			}
		case "prioridad_vuelos":
			if !verificarCantParametros(parametros, 2) {
				imprimirError(operacion)
				continue
			}
			cant_vuelos, _ := strconv.Atoi(parametros[1])
			if !prioridad_vuelos(cant_vuelos, hash) {
				imprimirError(operacion)
			}
		case "siguiente_vuelo":
			if !verificarCantParametros(parametros, 4) {
				imprimirError(operacion)
				continue
			}
			if !siguiente_vuelo(parametros[1], parametros[2], parametros[3], abb) {
				imprimirError(operacion)
			}
		case "borrar":
			if !verificarCantParametros(parametros, 3) {
				imprimirError(operacion)
				continue
			}
			if !borrar(parametros[1], parametros[2], abb, hash) {
				imprimirError(operacion)
			}
		default:
			imprimirError(operacion)
		}
	}
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
	}

	hash.Iterar(func(clave string, dato TDAVuelo.Vuelo) bool {
		abb.Guardar(tuple{fecha: dato.Fecha(), codigo: dato.Numero_vuelo()}, dato)
		return true
	})

	fmt.Println("OK")
	return true
}

func ver_tablero(cant_vuelos int, modo string, desde, hasta string, abb diccionario.DiccionarioOrdenado[tuple, TDAVuelo.Vuelo]) bool {
	if modo != "desc" && modo != "asc" || cant_vuelos <= 0 || strings.Compare(desde, hasta) > 0 { //
		return false //con errores
	}

	nose := tuple{fecha: desde, codigo: "0"}
	nose2 := tuple{fecha: hasta, codigo: "99999999999"}

	// iter := abb.IteradorRango(&nose, &nose2)
	// var resultados []tuple

	// for iter.HaySiguiente() {
	// 	clave, _ := iter.VerActual()
	// 	resultados = append(resultados, clave)
	// 	iter.Siguiente()
	// }

	// if modo == "asc" {
	// 	for i := 0; i < cant_vuelos; i++ {
	// 		clave := resultados[i]
	// 		fmt.Println(clave.fecha, "-", clave.codigo)
	// 	}
	// } else {
	// 	for i := len(resultados) - 1; i >= (len(resultados) - cant_vuelos); i-- {
	// 		clave := resultados[i]
	// 		fmt.Println(clave.fecha, "-", clave.codigo)
	// 	}
	// }

	if modo == "asc" {
		var contador_asc int
		abb.IterarRango(&nose, &nose2, func(clave tuple, dato TDAVuelo.Vuelo) bool {
			if contador_asc == cant_vuelos {
				return false
			}
			fmt.Println(clave.fecha, "-", clave.codigo)
			contador_asc++
			return true
		})
	}

	if modo == "desc" {
		pila := pila.CrearPilaDinamica[tuple]()
		abb.IterarRango(&nose, &nose2, func(clave tuple, dato TDAVuelo.Vuelo) bool {
			pila.Apilar(clave)
			return true
		})
		for i := 0; i < cant_vuelos; i++ {
			tope := pila.Desapilar()
			fmt.Println(tope.fecha, "-", tope.codigo)
		}
	}

	fmt.Println("OK")
	return true //sin errores
}

func info_vuelo(numero_vuelo string, hash diccionario.Diccionario[string, TDAVuelo.Vuelo]) bool {
	if !hash.Pertenece(numero_vuelo) {
		return false
	}
	vuelo := hash.Obtener(numero_vuelo)
	fmt.Println(vuelo.Obtener_string())
	fmt.Println("OK")
	return true
}

func prioridad_vuelos(cant_vuelos int, hash diccionario.Diccionario[string, TDAVuelo.Vuelo]) bool {
	if cant_vuelos < 0 {
		return false
	}
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

func siguiente_vuelo(origen, destino, desde string, abb diccionario.DiccionarioOrdenado[tuple, TDAVuelo.Vuelo]) bool {
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
	if !hallado {
		fmt.Println("No hay vuelo registrado desde", origen, "hacia", destino, "desde", desde)
	}
	fmt.Println("OK")
	return true
}

func borrar(fecha_desde, fecha_hasta string, abb diccionario.DiccionarioOrdenado[tuple, TDAVuelo.Vuelo], hash diccionario.Diccionario[string, TDAVuelo.Vuelo]) bool {
	desde := tuple{fecha: fecha_desde, codigo: "0"}
	hasta := tuple{fecha: fecha_hasta, codigo: "99999999999"}

	var claves []tuple
	abb.IterarRango(&desde, &hasta, func(clave tuple, dato TDAVuelo.Vuelo) bool {
		claves = append(claves, clave)
		return true
	})
	for i := 0; i < len(claves); i++ {
		codigo := claves[i].codigo
		vuelo := hash.Obtener(codigo)
		fmt.Println(vuelo.Obtener_string())
		hash.Borrar(codigo)
	}

	fmt.Println("OK")
	return true
}

func imprimirError(comando string) {
	fmt.Fprintln(os.Stderr, "Error en comando", comando)
}

func verificarCantParametros(parametros []string, cantidad int) bool {
	return len(parametros) == cantidad
}

package TDATabla

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

type tabla struct {
	base_datos diccionario.Diccionario[string, TDAVuelo.Vuelo]
	dicc_fecha diccionario.DiccionarioOrdenado[tuple, TDAVuelo.Vuelo]
	// dicc_prioridad cola_prioridad.ColaPrioridad[numVuelo_prioridad]
}

const PARAMETROS_AGREGAR_ARCHIVO = 2
const PARAMETROS_VER_TABLERO = 5
const PARAMETROS_INFO_VUELO = 2
const PARAMETROS_PRIORIDAD_VUELOS = 2
const PARAMETROS_SIGUIENTE_VUELO = 4
const PARAMETROS_BORRAR = 3

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

func CrearTabla() Tabla {
	return &tabla{base_datos: diccionario.CrearHash[string, TDAVuelo.Vuelo](),
		dicc_fecha: diccionario.CrearABB[tuple, TDAVuelo.Vuelo](comparacion_fechas_ascendente)}

}

func (tabla *tabla) Agregar_archivo(parametros []string) {
	if !verificarCantParametros(parametros, PARAMETROS_AGREGAR_ARCHIVO) {
		imprimirError("agregar_archivo")
		return
	}
	ruta := parametros[1]
	archivo, err := os.Open(ruta)
	if err != nil {
		imprimirError("agregar_archivo")
		return
	}
	defer archivo.Close()
	lectura := bufio.NewScanner(archivo)
	for lectura.Scan() {
		linea := lectura.Text()
		linea_sep := strings.Split(linea, ",")
		vuelo := TDAVuelo.CrearVuelo(linea_sep)
		tabla.base_datos.Guardar(vuelo.Numero_vuelo(), vuelo)
	}

	tabla.dicc_fecha = diccionario.CrearABB[tuple, TDAVuelo.Vuelo](comparacion_fechas_ascendente)
	tabla.base_datos.Iterar(func(clave string, dato TDAVuelo.Vuelo) bool {
		tabla.dicc_fecha.Guardar(tuple{fecha: dato.Fecha(), codigo: dato.Numero_vuelo()}, dato)
		return true
	})

	fmt.Println("OK")
}

func (tabla *tabla) Ver_tablero(parametros []string) {
	if !verificarCantParametros(parametros, PARAMETROS_VER_TABLERO) {
		imprimirError("ver_tablero")
		return
	}
	cant_vuelos, _ := strconv.Atoi(parametros[1])
	modo := parametros[2]
	desde := parametros[3]
	hasta := parametros[4]

	if modo != "desc" && modo != "asc" || cant_vuelos <= 0 || strings.Compare(desde, hasta) > 0 { //
		imprimirError("ver_tablero")
		return
	}

	fecha_desde := tuple{fecha: desde, codigo: "0"}
	fecha_hasta := tuple{fecha: hasta, codigo: "99999999999"}

	if modo == "asc" {
		var contador_asc int
		tabla.dicc_fecha.IterarRango(&fecha_desde, &fecha_hasta, func(clave tuple, dato TDAVuelo.Vuelo) bool {
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
		tabla.dicc_fecha.IterarRango(&fecha_desde, &fecha_hasta, func(clave tuple, dato TDAVuelo.Vuelo) bool {
			pila.Apilar(clave)
			return true
		})
		for i := 0; i < cant_vuelos && !pila.EstaVacia(); i++ {
			tope := pila.Desapilar()
			fmt.Println(tope.fecha, "-", tope.codigo)
		}
	}
	fmt.Println("OK")
}

func (tabla *tabla) Info_vuelo(parametros []string) {
	if !verificarCantParametros(parametros, PARAMETROS_INFO_VUELO) {
		imprimirError("info_vuelo")
		return
	}
	numero_vuelo := parametros[1]

	if !tabla.base_datos.Pertenece(numero_vuelo) {
		imprimirError("info_vuelo")
		return
	}
	vuelo := tabla.base_datos.Obtener(numero_vuelo)
	fmt.Println(vuelo.Obtener_string())
	fmt.Println("OK")
}

func (tabla *tabla) Siguiente_vuelo(parametros []string) {
	if !verificarCantParametros(parametros, PARAMETROS_SIGUIENTE_VUELO) {
		imprimirError("siguiente_vuelo")
		return
	}
	origen := parametros[1]
	destino := parametros[2]
	desde := parametros[3]

	hallado := false
	fecha := tuple{fecha: desde, codigo: "0"}
	tabla.dicc_fecha.IterarRango(&fecha, nil, func(clave tuple, vuelo TDAVuelo.Vuelo) bool {
		if origen == vuelo.Aeropuerto_origen() && destino == vuelo.Aeropuerto_destino() {
			fmt.Println(vuelo.Obtener_string())
			hallado = true
			return false
		}
		return true
	})
	if !hallado {
		fmt.Println("No hay vuelo registrado desde", origen, "hacia", destino, "desde", desde)
	}
	fmt.Println("OK")
}

func (tabla *tabla) Prioridad_vuelos(parametros []string) {
	if !verificarCantParametros(parametros, PARAMETROS_PRIORIDAD_VUELOS) {
		imprimirError("prioridad_vuelo")
		return
	}
	cant_vuelos, err := strconv.Atoi(parametros[1])
	if cant_vuelos <= 0 {
		imprimirError("prioridad_vuelos")
		return
	}
	if err != nil {
		imprimirError("prioridad_vuelos")
		return
	}
	heap := cola_prioridad.CrearHeap(comparar_numero_vuelo_prioridad) //heap maximos
	iterador := tabla.base_datos.Iterador()
	cant := 0
	for iterador.HaySiguiente() {
		_, vuelo := iterador.VerActual()
		num_vuelo_prioridad := numVuelo_prioridad{numero_vuelo: vuelo.Numero_vuelo(), prioridad: vuelo.Prioridad()}
		heap.Encolar(num_vuelo_prioridad)
		cant++
		iterador.Siguiente()
	}

	for i := 0; i < cant_vuelos && !heap.EstaVacia(); i++ {
		tope := heap.Desencolar()
		fmt.Println(tope.prioridad, "-", tope.numero_vuelo)
	}
	fmt.Println("OK")
}

func (tabla *tabla) Borrar(parametros []string) {
	if !verificarCantParametros(parametros, PARAMETROS_BORRAR) {
		imprimirError("borrar")
		return
	}

	fecha_desde := parametros[1]
	fecha_hasta := parametros[2]

	if strings.Compare(fecha_desde, fecha_hasta) > 0 {
		imprimirError("borrar")
		return
	}

	desde := tuple{fecha: fecha_desde, codigo: "0"}
	hasta := tuple{fecha: fecha_hasta, codigo: "99999999999"}

	var claves []tuple
	tabla.dicc_fecha.IterarRango(&desde, &hasta, func(clave tuple, dato TDAVuelo.Vuelo) bool {
		claves = append(claves, clave)
		return true
	})
	for i := 0; i < len(claves); i++ {
		codigo := claves[i].codigo
		vuelo := tabla.base_datos.Obtener(codigo)
		fmt.Println(vuelo.Obtener_string())
		tabla.base_datos.Borrar(codigo)
		tabla.dicc_fecha.Borrar(claves[i])
	}

	fmt.Println("OK")
}

func imprimirError(comando string) {
	fmt.Fprintln(os.Stderr, "Error en comando", comando)
}

func verificarCantParametros(parametros []string, cantidad int) bool {
	return len(parametros) == cantidad
}

package TDATabla

import (
	"bufio"
	"fmt"
	"os"
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
}

func comparacion_fechas_ascendente(fecha1 tuple, fecha2 tuple) int {
	result := strings.Compare(fecha1.fecha, fecha2.fecha)
	if result == 0 {
		result = strings.Compare(fecha1.codigo, fecha2.codigo)
	}
	return result
}

func comparar_numero_vuelo_prioridad(a, b numVuelo_prioridad) int {
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

func (tabla *tabla) AgregarArchivo(ruta string) {
	archivo, err := os.Open(ruta)
	if err != nil {
		tabla.ImprimirError("agregar_archivo")
		return
	}
	defer archivo.Close()
	lectura := bufio.NewScanner(archivo)
	for lectura.Scan() {
		linea := lectura.Text()
		linea_sep := strings.Split(linea, ",")
		vuelo := TDAVuelo.CrearVuelo(linea_sep)
		tabla.base_datos.Guardar(vuelo.NumeroVuelo(), vuelo)
	}

	tabla.dicc_fecha = diccionario.CrearABB[tuple, TDAVuelo.Vuelo](comparacion_fechas_ascendente)
	tabla.base_datos.Iterar(func(clave string, dato TDAVuelo.Vuelo) bool {
		tabla.dicc_fecha.Guardar(tuple{fecha: dato.Fecha(), codigo: dato.NumeroVuelo()}, dato)
		return true
	})

	fmt.Println("OK")
}

func (tabla *tabla) VerTablero(cant_vuelos int,modo,desde,hasta string) {
	if modo != "desc" && modo != "asc" || cant_vuelos <= 0 || strings.Compare(desde, hasta) > 0 { 
		tabla.ImprimirError("ver_tablero")
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

func (tabla *tabla) InfoVuelo(numero_vuelo string) {
	if !tabla.base_datos.Pertenece(numero_vuelo) {
		tabla.ImprimirError("info_vuelo")
		return
	}
	vuelo := tabla.base_datos.Obtener(numero_vuelo)
	fmt.Println(vuelo.ObtenerString())
	fmt.Println("OK")
}

func (tabla *tabla) SiguienteVuelo(origen,destino,fecha_desde string) {
	hallado := false
	fecha := tuple{fecha: fecha_desde, codigo: "0"}
	tabla.dicc_fecha.IterarRango(&fecha, nil, func(clave tuple, vuelo TDAVuelo.Vuelo) bool {
		if origen == vuelo.AeropuertoOrigen() && destino == vuelo.AeropuertoDestino() {
			fmt.Println(vuelo.ObtenerString())
			hallado = true
			return false
		}
		return true
	})
	if !hallado {
		fmt.Println("No hay vuelo registrado desde", origen, "hacia", destino, "desde", fecha_desde)
	}
	fmt.Println("OK")
}

func (tabla *tabla) PrioridadVuelos(cant_vuelos int) {
	if cant_vuelos <= 0 {
		tabla.ImprimirError("prioridad_vuelos")
		return
	}
	heap := cola_prioridad.CrearHeap(comparar_numero_vuelo_prioridad)
	iterador := tabla.base_datos.Iterador()
	cant := 0
	for iterador.HaySiguiente() {
		_, vuelo := iterador.VerActual()
		num_vuelo_prioridad := numVuelo_prioridad{numero_vuelo: vuelo.NumeroVuelo(), prioridad: vuelo.Prioridad()}
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

func (tabla *tabla) Borrar(fecha_desde,fecha_hasta string) {
	if strings.Compare(fecha_desde, fecha_hasta) > 0 {
		tabla.ImprimirError("borrar")
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
		fmt.Println(vuelo.ObtenerString())
		tabla.base_datos.Borrar(codigo)
		tabla.dicc_fecha.Borrar(claves[i])
	}

	fmt.Println("OK")
}

func (tabla *tabla) ImprimirError(comando string){
	fmt.Fprintln(os.Stderr, "Error en comando", comando)
}

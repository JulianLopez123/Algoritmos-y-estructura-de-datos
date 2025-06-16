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

type numVuelo_fecha struct {
	fecha        string
	numero_vuelo string
}
type numVuelo_prioridad struct {
	numero_vuelo string
	prioridad    int
}

type conexion struct {
	aeropuerto_origen  string
	aeropuerto_destino string
}

type conexion_fecha struct {
	fecha    string
	conexion conexion
}

type tabla struct {
	base_datos      diccionario.Diccionario[string, TDAVuelo.Vuelo]
	dicc_fecha      diccionario.DiccionarioOrdenado[numVuelo_fecha, TDAVuelo.Vuelo]
	dicc_conexiones diccionario.DiccionarioOrdenado[conexion_fecha, TDAVuelo.Vuelo]
	clave_mayor     string
}

func comparacion_fechas_ascendente(fecha1 numVuelo_fecha, fecha2 numVuelo_fecha) int {
	result := strings.Compare(fecha1.fecha, fecha2.fecha)
	if result == 0 {
		result = strings.Compare(fecha1.numero_vuelo, fecha2.numero_vuelo)
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

func comparar_conexiones(a, b conexion_fecha) int {
	resultado := strings.Compare(a.conexion.aeropuerto_origen, b.conexion.aeropuerto_origen)
	if resultado == 0 {
		resultado = strings.Compare(a.conexion.aeropuerto_destino, b.conexion.aeropuerto_destino)
	}
	if resultado == 0 {
		resultado = strings.Compare(a.fecha, b.fecha)
	}
	return resultado
}

func CrearTabla() Tabla {
	return &tabla{base_datos: diccionario.CrearHash[string, TDAVuelo.Vuelo](),
		dicc_fecha: diccionario.CrearABB[numVuelo_fecha, TDAVuelo.Vuelo](comparacion_fechas_ascendente), clave_mayor: ""}
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

	var claveMayor string
	tabla.dicc_fecha = diccionario.CrearABB[numVuelo_fecha, TDAVuelo.Vuelo](comparacion_fechas_ascendente)
	tabla.dicc_conexiones = diccionario.CrearABB[conexion_fecha, TDAVuelo.Vuelo](comparar_conexiones)
	tabla.base_datos.Iterar(func(clave string, dato TDAVuelo.Vuelo) bool {
		if strings.Compare(clave, claveMayor) == 1 {
			claveMayor = clave
		}
		tabla.dicc_fecha.Guardar(numVuelo_fecha{fecha: dato.Fecha(), numero_vuelo: dato.NumeroVuelo()}, dato)
		tabla.dicc_conexiones.Guardar(conexion_fecha{conexion: conexion{aeropuerto_destino: dato.AeropuertoDestino(), aeropuerto_origen: dato.AeropuertoOrigen()}, fecha: dato.Fecha()}, dato)
		return true
	})

	tabla.clave_mayor = claveMayor

	fmt.Println("OK")
}

func (tabla *tabla) VerTablero(cant_vuelos int, modo, desde, hasta string) {
	if modo != "desc" && modo != "asc" || cant_vuelos <= 0 || strings.Compare(desde, hasta) > 0 {
		tabla.ImprimirError("ver_tablero")
		return
	}

	fecha_desde := numVuelo_fecha{fecha: desde, numero_vuelo: "0"}
	fecha_hasta := numVuelo_fecha{fecha: hasta, numero_vuelo: tabla.clave_mayor}

	if modo == "asc" {
		var contador_asc int
		tabla.dicc_fecha.IterarRango(&fecha_desde, &fecha_hasta, func(clave numVuelo_fecha, dato TDAVuelo.Vuelo) bool {
			if contador_asc == cant_vuelos {
				return false
			}
			fmt.Println(clave.fecha, "-", clave.numero_vuelo)
			contador_asc++
			return true
		})
	}

	if modo == "desc" {
		pila := pila.CrearPilaDinamica[numVuelo_fecha]()
		tabla.dicc_fecha.IterarRango(&fecha_desde, &fecha_hasta, func(clave numVuelo_fecha, dato TDAVuelo.Vuelo) bool {
			pila.Apilar(clave)
			return true
		})
		for i := 0; i < cant_vuelos && !pila.EstaVacia(); i++ {
			tope := pila.Desapilar()
			fmt.Println(tope.fecha, "-", tope.numero_vuelo)
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

func (tabla *tabla) SiguienteVuelo(origen, destino, fecha_desde string) {
	hallado := false
	clave := conexion_fecha{fecha: fecha_desde, conexion: conexion{aeropuerto_origen: origen, aeropuerto_destino: destino}}

	tabla.dicc_conexiones.IterarRango(&clave, nil, func(clave conexion_fecha, vuelo TDAVuelo.Vuelo) bool {
		if origen == clave.conexion.aeropuerto_origen && destino == clave.conexion.aeropuerto_destino {
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
	for iterador.HaySiguiente() {
		_, vuelo := iterador.VerActual()
		num_vuelo_prioridad := numVuelo_prioridad{numero_vuelo: vuelo.NumeroVuelo(), prioridad: vuelo.Prioridad()}
		heap.Encolar(num_vuelo_prioridad)
		iterador.Siguiente()
	}

	for i := 0; i < cant_vuelos && !heap.EstaVacia(); i++ {
		tope := heap.Desencolar()
		fmt.Println(tope.prioridad, "-", tope.numero_vuelo)
	}
	fmt.Println("OK")
}

func (tabla *tabla) Borrar(fecha_desde, fecha_hasta string) {
	if strings.Compare(fecha_desde, fecha_hasta) > 0 {
		tabla.ImprimirError("borrar")
		return
	}

	desde := numVuelo_fecha{fecha: fecha_desde, numero_vuelo: "0"}
	hasta := numVuelo_fecha{fecha: fecha_hasta, numero_vuelo: tabla.clave_mayor}

	var claves []numVuelo_fecha
	tabla.dicc_fecha.IterarRango(&desde, &hasta, func(clave numVuelo_fecha, dato TDAVuelo.Vuelo) bool {
		claves = append(claves, clave)
		return true
	})
	for i := 0; i < len(claves); i++ {
		codigo := claves[i].numero_vuelo
		vuelo := tabla.base_datos.Obtener(codigo)
		fmt.Println(vuelo.ObtenerString())
		tabla.base_datos.Borrar(codigo)
		tabla.dicc_fecha.Borrar(claves[i])
		clave_conexion := conexion_fecha{fecha: vuelo.Fecha(), conexion: conexion{aeropuerto_origen: vuelo.AeropuertoOrigen(), aeropuerto_destino: vuelo.AeropuertoDestino()}}
		if tabla.dicc_conexiones.Pertenece(clave_conexion) {
			tabla.dicc_conexiones.Borrar(clave_conexion)
		}
	}

	fmt.Println("OK")
}

func (tabla *tabla) ImprimirError(comando string) {
	fmt.Fprintln(os.Stderr, "Error en comando", comando)
}

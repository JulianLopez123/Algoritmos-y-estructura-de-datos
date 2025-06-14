package tp2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tdas/diccionario"
)

func comparacion_fechas_ascendente(fecha1 string, fecha2 string) int {
	if fecha1 > fecha2 {
		return 1
	} else if fecha1 < fecha2 {
		return -1
	}
	return 0
}

func comparacion_fechas_descendente(fecha1, fecha2 string) int {
	return comparacion_fechas_ascendente(fecha2, fecha1)
}

func agregar_archivo(ruta string) {
	hash := diccionario.CrearHash[int, *Vuelo]()
	archivo, _ := os.Open(ruta)
	defer archivo.Close()
	lectura := bufio.NewScanner(archivo)
	for lectura.Scan() {
		linea := lectura.Text()
		linea_sep := strings.Split(linea, ",")
		vuelo := CrearVuelo(linea_sep)
		hash.Guardar(vuelo.Obtener_numero_vuelo(), &vuelo)
	}
	printVuelosEnArchivo(ruta, hash)
}

func main() {
	args := os.Args

	operacion := args[1]

	switch operacion {
	case "agregar_archivo":
		agregar_archivo(args[2])

	case "ver_tablero":
		cant_vuelos, _ := strconv.Atoi(args[2])
		ver_tablero(cant_vuelos, args[3], args[4], args[5])
	}
}

func ver_tablero(cant_vuelos int, modo, desde, hasta string) {
	var abb diccionario.DiccionarioOrdenado[string, Vuelo]
	if modo == "desc" {
		abb = diccionario.CrearABB[string, Vuelo](comparacion_fechas_descendente)
	} else if modo == "asc" {
		abb = diccionario.CrearABB[string, Vuelo](comparacion_fechas_ascendente)
	} else {
		panic("Error en el modo de ordenamiento")
	}

	datos_almacenados, _ := os.Open("datos_almacenados.csv")
	defer datos_almacenados.Close()
	lectura := bufio.NewScanner(datos_almacenados)
	for lectura.Scan() {
		linea := lectura.Text()
		linea_sep := strings.Split(linea, ",")
		vuelo := CrearVuelo(linea_sep)
		abb.Guardar(vuelo.Obtener_fecha(), vuelo)
	}
	var contador int

	abb.IterarRango(&desde, &hasta, func(clave string, dato Vuelo) bool {
		if contador == cant_vuelos {
			return false
		}
		fmt.Println()
		contador++
		return true
	})
}

// func recuperar_datos_almacenados(){
// 	datos_almacenados, _ := os.Open("datos_almacenados.csv")
// 	defer datos_almacenados.Close()
// 	lectura := bufio.NewScanner(datos_almacenados)
// 	for lectura.Scan() {
// 		vuelo := lectura.Text()
// 		vuelo_sep := strings.Split(vuelo,",")
// 		vuelos_juntos := strings.Join(vuelo_sep, "	")
// 		fmt.Println(vuelos_juntos)
// 	}
// }

// func printVuelos(ruta string) {
// 	archivo, _ := os.Open(ruta)
// 	defer archivo.Close()
// 	lectura := bufio.NewScanner(archivo)
// 	for lectura.Scan() {
// 		vuelo := lectura.Text()
// 		vuelo_sep := strings.Split(vuelo, ",")
// 		vuelos_juntos := strings.Join(vuelo_sep, "	")
// 		fmt.Println(vuelos_juntos)
// 	}
// }

// func print_vuelo(vuelo *vuelo) {
// 	vuelos_juntos := strings.Join(vuelo_sep, "	")
// }

func printVuelosEnArchivo(ruta string, hash *diccionario.Diccionario) {
	archivo, _ := os.Open(ruta)
	defer archivo.Close()
	write := bufio.NewWriter(archivo)
	for _, valor := range hash.tabla {
		vuelo := &valor
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

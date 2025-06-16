package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"tp2/TDATabla"
)

const (
	PARAMETROS_AGREGAR_ARCHIVO  = 2
	PARAMETROS_VER_TABLERO      = 5
	PARAMETROS_INFO_VUELO       = 2
	PARAMETROS_PRIORIDAD_VUELOS = 2
	PARAMETROS_SIGUIENTE_VUELO  = 4
	PARAMETROS_BORRAR           = 3
)

func main() {
	lectura := bufio.NewScanner(os.Stdin)
	tabla := TDATabla.CrearTabla()
	for {
		lectura.Scan()
		linea := lectura.Text()
		if linea == "" {
			break
		}
		ingreso := strings.Split(linea, " ")
		comando := ingreso[0]
		parametros := ingreso[1:]

		switch comando {
		case "agregar_archivo":
			if !cantParametrosCorrectos(ingreso, PARAMETROS_AGREGAR_ARCHIVO) {
				tabla.ImprimirError(comando)
				continue
			}
			tabla.AgregarArchivo(parametros[0])

		case "ver_tablero":
			if !cantParametrosCorrectos(ingreso, PARAMETROS_VER_TABLERO) {
				tabla.ImprimirError(comando)
				continue
			}
			cant_vuelos, err := strconv.Atoi(parametros[0])
			if err != nil {
				tabla.ImprimirError(comando)
				continue
			}
			tabla.VerTablero(cant_vuelos, parametros[1], parametros[2], parametros[3])

		case "info_vuelo":
			if !cantParametrosCorrectos(ingreso, PARAMETROS_INFO_VUELO) {
				tabla.ImprimirError(comando)
				continue
			}
			tabla.InfoVuelo(parametros[0])

		case "prioridad_vuelos":
			if !cantParametrosCorrectos(ingreso, PARAMETROS_PRIORIDAD_VUELOS) {
				tabla.ImprimirError(comando)
				continue
			}
			cant_vuelos, err := strconv.Atoi(parametros[0])
			if err != nil {
				tabla.ImprimirError(comando)
				continue
			}
			tabla.PrioridadVuelos(cant_vuelos)

		case "siguiente_vuelo":
			if !cantParametrosCorrectos(ingreso, PARAMETROS_SIGUIENTE_VUELO) {
				tabla.ImprimirError(comando)
				continue
			}
			tabla.SiguienteVuelo(parametros[0], parametros[1], parametros[2])

		case "borrar":
			if !cantParametrosCorrectos(ingreso, PARAMETROS_BORRAR) {
				tabla.ImprimirError(comando)
				continue
			}
			tabla.Borrar(parametros[0], parametros[1])

		default:
			tabla.ImprimirError(comando)
		}
	}
}

func cantParametrosCorrectos(parametros []string, cantidad int) bool {
	return len(parametros) == cantidad
}

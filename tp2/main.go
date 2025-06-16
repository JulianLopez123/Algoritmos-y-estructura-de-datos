package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tp2/TDATabla"
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
		parametros := strings.Split(linea, " ")
		operacion := parametros[0]
	
		switch operacion {
		case "agregar_archivo":
			tabla.Agregar_archivo(parametros)
		case "ver_tablero":
			tabla.Ver_tablero(parametros)
		case "info_vuelo":
			tabla.Info_vuelo(parametros)
		case "prioridad_vuelos":
			tabla.Prioridad_vuelos(parametros)
		case "siguiente_vuelo":
			tabla.Siguiente_vuelo(parametros)
		case "borrar":
			tabla.Borrar(parametros)
		default:
			fmt.Fprintln(os.Stderr, "Error en comando", operacion)
		}
	}

}
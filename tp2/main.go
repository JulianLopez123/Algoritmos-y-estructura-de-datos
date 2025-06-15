package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tp2/TDATabla"
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
			fmt.Println("Error")
		}
	}

}
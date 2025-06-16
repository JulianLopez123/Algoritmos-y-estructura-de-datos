package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tp2/TDATabla"
)

const(
	PARAMETROS_AGREGAR_ARCHIVO = 2
	PARAMETROS_VER_TABLERO = 5
	PARAMETROS_INFO_VUELO = 2
	PARAMETROS_PRIORIDAD_VUELOS = 2
	PARAMETROS_SIGUIENTE_VUELO = 4
	PARAMETROS_BORRAR = 3

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

		if !validarComando(comando,parametros,tabla){
			fmt.Fprintln(os.Stderr, "Error en comando", )
		}
	}
}		
	
func validarComando(comando string, parametros []string,tabla TDATabla.Tabla)bool{
	switch comando {
	case "agregar_archivo":
		if verificarCantParametros(parametros,PARAMETROS_AGREGAR_ARCHIVO){
			return tabla.AgregarArchivo(parametros[0])
		}
		return false

	case "ver_tablero":
		if verificarCantParametros(parametros,PARAMETROS_VER_TABLERO){
			cant_vuelos, err:= strconv.Atoi(parametros[1])
			if err != nil{
				return false
			}
			return tabla.VerTablero(cant_vuelos,parametros[1],parametros[2])
		}
		return false

	case "info_vuelo":
		if verificarCantParametros(parametros,PARAMETROS_INFO_VUELO){
			return tabla.InfoVuelo(parametros)	
		}
		return false

	case "prioridad_vuelos":
		if verificarCantParametros(parametros,PARAMETROS_PRIORIDAD_VUELOS){
			return tabla.PrioridadVuelos(parametros)
		}
		
	case "siguiente_vuelo":
		if verificarCantParametros(parametros,PARAMETROS_SIGUIENTE_VUELO){
			return tabla.SiguienteVuelo(parametros)
		}
		return false

	case "borrar":
		if verificarCantParametros(parametros,PARAMETROS_BORRAR){
			return tabla.Borrar(parametros)
		}
		return false

	default:
		return false
	}
}

func validarYEjecutarOperacion(primitiva func (TDATabla.Tabla))

func verificarCantParametros(parametros []string, cantidad int) bool {
	return len(parametros) == cantidad
}

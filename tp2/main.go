package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tp2/TDATabla"
)

const (
	PARAMETROS_AGREGAR_ARCHIVO  = 1
	PARAMETROS_VER_TABLERO      = 4
	PARAMETROS_INFO_VUELO       = 1
	PARAMETROS_PRIORIDAD_VUELOS = 1
	PARAMETROS_SIGUIENTE_VUELO  = 3
	PARAMETROS_BORRAR           = 2
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
		ejecutarparametros(ingreso,tabla)
	}
}

func ejecutarparametros(ingreso []string,tabla TDATabla.Tabla){
	comando := ingreso[0]
	parametros := ingreso[1:]

	switch comando {
	case "agregar_archivo":
		casoAgregarArchivo(tabla,parametros,comando)
	case "ver_tablero":
		casoVerTablero(tabla,parametros,comando)
	case "info_vuelo":
		casoInfoVuelo(tabla,parametros,comando)
	case "prioridad_vuelos":
		casoPrioridadVuelos(tabla,parametros,comando)
	case "siguiente_vuelo":
		casoSiguienteVuelo(tabla,parametros,comando)
	case "borrar":
		casoBorrar(tabla,parametros,comando)
	default:
		imprimirError(comando)
	}
}

func casoAgregarArchivo(tabla TDATabla.Tabla,parametros []string,comando string){
	if !cantParametrosCorrectos(parametros, PARAMETROS_AGREGAR_ARCHIVO) {
		imprimirError(comando)
		return
	}
	err := tabla.AgregarArchivo(parametros[0])
	if err{
		imprimirError(comando)
		return 
	}
	fmt.Println("OK")
	
}

func casoVerTablero(tabla TDATabla.Tabla,parametros []string,comando string){
	if !cantParametrosCorrectos(parametros, PARAMETROS_VER_TABLERO) {
		imprimirError(comando)
		return
	}
	cant_vuelos, err := strconv.Atoi(parametros[0])
	if err != nil {
		imprimirError(comando)
		return
	}

	vuelos,err2 := tabla.VerTablero(cant_vuelos, parametros[1], parametros[2], parametros[3])
	if err2{
		imprimirError(comando)
		return
	}
	for _,vuelo := range vuelos{
		fmt.Println(vuelo)
	}
	fmt.Println("OK")
}

func casoInfoVuelo(tabla TDATabla.Tabla,parametros []string,comando string){
	if !cantParametrosCorrectos(parametros, PARAMETROS_INFO_VUELO) {
		imprimirError(comando)
		return
	}
	vuelo,err:=tabla.InfoVuelo(parametros[0])
	if err{
		imprimirError(comando)
		return
	}
	fmt.Println(vuelo)
	fmt.Println("OK")
}

func casoPrioridadVuelos(tabla TDATabla.Tabla,parametros []string,comando string){
	if !cantParametrosCorrectos(parametros, PARAMETROS_PRIORIDAD_VUELOS) {
		imprimirError(comando)
		return
	}
	cant_vuelos, err := strconv.Atoi(parametros[0])
	if err != nil {
		imprimirError(comando)
		return
	}
	vuelos,err2 := tabla.PrioridadVuelos(cant_vuelos)
	if err2{
		imprimirError(comando)
		return
	}
	for _,vuelo := range vuelos{
		fmt.Println(vuelo)
	}
	fmt.Println("OK")
}

func casoSiguienteVuelo(tabla TDATabla.Tabla,parametros []string,comando string){
	if !cantParametrosCorrectos(parametros, PARAMETROS_SIGUIENTE_VUELO) {
		imprimirError(comando)
		return
	}
	origen := parametros[0]
	destino := parametros[1]
	fecha_desde := parametros[2]
	resultado ,err:= tabla.SiguienteVuelo(origen,destino,fecha_desde)
	if err{
		fmt.Println("No hay vuelo registrado desde", origen, "hacia", destino, "desde", fecha_desde)
	}else{
		fmt.Println(resultado)
	}
	fmt.Println("OK")
}

func casoBorrar(tabla TDATabla.Tabla,parametros []string,comando string){
	if !cantParametrosCorrectos(parametros, PARAMETROS_BORRAR) {
		imprimirError(comando)
		return
	}
	vuelos,err:=tabla.Borrar(parametros[0], parametros[1])
	if err{
		imprimirError(comando)
		return
	}
	for _,vuelo := range vuelos{
		fmt.Println(vuelo)
	}
	fmt.Println("OK")
}

func cantParametrosCorrectos(parametros []string, cantidad int) bool {
	return len(parametros) == cantidad
}

func imprimirError(comando string) {
	fmt.Fprintln(os.Stderr, "Error en comando", comando)
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tdas/cola_prioridad"
	"tdas/diccionario"
	"tp2/TDAVuelo"
)

type numVuelo_prioridad struct {
	numero_vuelo int
	prioridad int
}


func comparacion_fechas_ascendente(fecha1 string,fecha2 string) int{
	if fecha1 > fecha2{
		return 1
	}else if fecha1 < fecha2{
		return -1 
	}
	return 0
}

func comparar_numero_vuelo_prioridad(a, b numVuelo_prioridad) int { //heap maximos
	resultado :=  a.prioridad - b.prioridad 
	if resultado == 0{
		return a.numero_vuelo - b.numero_vuelo
	}
	return resultado
}


func main() {
	lectura := bufio.NewScanner(os.Stdin)
	hash := diccionario.CrearHash[int, TDAVuelo.Vuelo]()
	abb := diccionario.CrearABB[string,TDAVuelo.Vuelo](comparacion_fechas_ascendente)
	

	for {
		lectura.Scan()
		linea := lectura.Text()
		parametros := strings.Split(linea, " ")
		operacion := parametros[0]

		switch operacion {
		case "agregar_archivo":
			if !agregar_archivo(parametros[1],hash,abb){
				imprimirError(operacion)
			}
		case "ver_tablero":
			cant_vuelos, _ := strconv.Atoi(parametros[1])
			if !ver_tablero(cant_vuelos, parametros[2], parametros[3], parametros[4],abb){
				imprimirError(operacion)
			}
		case "info_vuelo":
			numero_vuelo,_:= strconv.Atoi(parametros[1])
			if !info_vuelo(numero_vuelo,hash){
				imprimirError(operacion)
			}
		case "prioridad_vuelos":
			cant_vuelos,_ := strconv.Atoi(parametros[1])
			if !prioridad_vuelos(cant_vuelos,hash){
				imprimirError(operacion)
			}
		case "siguiente_vuelo":

		}
	}
	
}

func ver_tablero(cant_vuelos int,modo,desde,hasta string,abb diccionario.DiccionarioOrdenado[string,TDAVuelo.Vuelo])bool{
	if modo != "desc" && modo != "asc" || cant_vuelos <= 0 { //|| comparacion_fechas_ascendente(desde,hasta) < 0
		return false //con errores 
	}
	
	// var abb diccionario.DiccionarioOrdenado[string, TDAVuelo.Vuelo]
	// if modo == "desc"{
	//
	//	}else if modo == "asc"{
	//		abb = diccionario.CrearABB[string,TDAVuelo.Vuelo](comparacion_fechas_ascendente)
	//	}
	datos_almacenados, _ := os.Open("datos_almacenados.csv")
	defer datos_almacenados.Close()
	lectura := bufio.NewScanner(datos_almacenados)
	for lectura.Scan() {
		linea := lectura.Text()
		linea_sep := strings.Split(linea,",")
		vuelo := TDAVuelo.CrearVuelo(linea_sep)
		clave := fmt.Sprintf("%s - %d",vuelo.Fecha(),vuelo.Numero_vuelo())
		abb.Guardar(clave,vuelo)
	}

	// clave_desde := fmt.Sprintf("%s - 00000000",desde)//clave minima
	// clave_hasta := fmt.Sprintf("%s - 99999999",hasta)
	var contador int
	if modo == "desc"{
		desde,hasta = hasta,desde
	}

	abb.IterarRango(&desde , &hasta, func(clave string, dato TDAVuelo.Vuelo)bool{
		if contador == cant_vuelos{
			return false
		}
		fmt.Println(clave)
		contador ++
		return true
	})
	fmt.Println("OK")
	return true //sin errores
}


func agregar_archivo(ruta string, hash diccionario.Diccionario[int, TDAVuelo.Vuelo],abb diccionario.DiccionarioOrdenado[string,TDAVuelo.Vuelo])bool {
	archivo, err := os.Open(ruta)
	if err != nil{
		return false
	}
	defer archivo.Close()
	lectura := bufio.NewScanner(archivo)
	for lectura.Scan() {
		linea := lectura.Text()
		linea_sep := strings.Split(linea, ",")
		vuelo := TDAVuelo.CrearVuelo(linea_sep)
		hash.Guardar(vuelo.Numero_vuelo(), vuelo)
		abb.Guardar(vuelo.Fecha(),vuelo)
	}
	fmt.Println("OK")
	return true
}

func info_vuelo(numero_vuelo int, hash diccionario.Diccionario[int,TDAVuelo.Vuelo])bool{
	if !hash.Pertenece(numero_vuelo){
		return false
	}
	vuelo:= hash.Obtener(numero_vuelo)
	fmt.Println(vuelo.Obtener_string()) 
	println("OK")
	return true
}

func prioridad_vuelos(cant_vuelos int,hash diccionario.Diccionario[int,TDAVuelo.Vuelo]) bool{
	heap := cola_prioridad.CrearHeap(comparar_numero_vuelo_prioridad) //heap maximos
	iterador := hash.Iterador()
	for iterador.HaySiguiente(){
		_,vuelo := iterador.VerActual()
		num_vuelo_prioridad := numVuelo_prioridad{numero_vuelo:vuelo.Numero_vuelo(),prioridad: vuelo.Prioridad()}
		heap.Encolar(num_vuelo_prioridad)
		iterador.Siguiente()
	}
	
	top := make([]numVuelo_prioridad,cant_vuelos)
	for i := 0; i < cant_vuelos;i++{
		top[i] = heap.Desencolar()
	}
	for i:= 0; i < cant_vuelos;i++{
		numero_vuelo,prioridad := top[i].numero_vuelo,top[i].prioridad
		fmt.Println(prioridad,"-",numero_vuelo)
	}
	fmt.Println("OK")
	return true
}


func imprimirError(comando string){
	fmt.Fprintln(os.Stderr, "Error en comando", comando)
}
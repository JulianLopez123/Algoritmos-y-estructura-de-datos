package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tdas/cola"
	"tdas/pila"
)

// Aclaro (por las dudas) que probe con un archivo main separado pero al entregarlo generaba errores
func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		linea := input.Text()
		fmt.Println(ConvertirANotacionPolacaInversa(linea))
	}
}

func ConvertirANotacionPolacaInversa(operacion string) string {
	var cadena_final []string
	operacion_lista := transformarOperacionALista(operacion)
	cola := cola.CrearColaEnlazada[string]()
	pila := pila.CrearPilaDinamica[string]()
	for _, valor := range operacion_lista {
		evaluarCasos(valor, cola, pila)
	}
	for !pila.EstaVacia() {
		cola.Encolar(pila.Desapilar())
	}
	for !cola.EstaVacia() {
		cadena_final = append(cadena_final, cola.Desencolar())
	}
	return strings.Join(cadena_final, " ")
}

func transformarOperacionALista(cadena string) []string {
	var cadena_normalizada []string
	cadena_caracteres := strings.Split(cadena, "")
	for _, caracter := range cadena_caracteres {
		longitud_cadena_normalizada := len(cadena_normalizada)
		if longitud_cadena_normalizada > 0 && isnumeric(cadena_normalizada[longitud_cadena_normalizada-1]) && isnumeric(caracter) {
			cadena_normalizada[longitud_cadena_normalizada-1] += caracter
		} else {
			cadena_normalizada = append(cadena_normalizada, caracter)
		}
	}
	return cadena_normalizada
}

func evaluarCasos(valor string, cola cola.Cola[string], pila pila.Pila[string]) {
	switch {
	case isnumeric(valor):
		cola.Encolar(valor)
	case isoperator(valor):
		for !pila.EstaVacia() && (calcularPrecedencia(pila.VerTope()) >= calcularPrecedencia(valor)) && valor != "^" {
			cola.Encolar(pila.Desapilar())
		}
		pila.Apilar(valor)
	case valor == "(":
		pila.Apilar(valor)
	case valor == ")":
		for pila.VerTope() != "(" {
			cola.Encolar(pila.Desapilar())
		}
		pila.Desapilar()
	}
}
func calcularPrecedencia(operador string) int {
	switch operador {
	case "+", "-":
		return 0
	case "*", "/":
		return 1
	case "^":
		return 2
	default:
		return -1
	}
}
func isnumeric(caracter string) bool {
	_, err := strconv.Atoi(caracter)
	return err == nil
}

func isoperator(caracter string) bool {
	switch caracter {
	case "+", "-", "/", "*", "^":
		return true
	}
	return false
}

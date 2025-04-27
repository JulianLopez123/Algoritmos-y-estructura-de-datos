package conversor

import (
	"strconv"
	"strings"
	"tdas/cola"
	"tdas/pila"
)

func ConvertirANotacionPolacaInversa(operacion string) string {
	var cadena_final []string
	operacion_lista := transformarOperacionALista(operacion)
	cola := cola.CrearColaEnlazada[string]()
	pila := pila.CrearPilaDinamica[string]()
	for _, valor := range operacion_lista {
		evaluarCaracter(valor, cola, pila)
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
	var normalizada []string
	caracteres := strings.Split(cadena, "")
	for _, caracter := range caracteres {
		longitud_normalizada := len(normalizada)
		if longitud_normalizada > 0 && esNumerico(normalizada[longitud_normalizada-1]) && esNumerico(caracter) {
			normalizada[longitud_normalizada-1] += caracter
		} else {
			normalizada = append(normalizada, caracter)
		}
	}
	return normalizada
}

func evaluarCaracter(valor string, cola cola.Cola[string], pila pila.Pila[string]) {
	switch {
	case esNumerico(valor):
		cola.Encolar(valor)
	case esOperador(valor):
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
func esNumerico(caracter string) bool {
	_, err := strconv.Atoi(caracter)
	return err == nil
}

func esOperador(caracter string) bool {
	switch caracter {
	case "+", "-", "/", "*", "^":
		return true
	}
	return false
}

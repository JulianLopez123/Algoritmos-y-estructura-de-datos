package ejercicios

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {
	*x, *y = *y, *x
}

// Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el el arreglo es de largo 0. Si el máximo
// elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
func Maximo(vector []int) int {
	if len(vector) == 0 {
		return -1
	}
	indice_mayor := 0
	for indice, valor := range vector {
		if valor > vector[indice_mayor] {
			indice_mayor = indice
		}
	}
	return indice_mayor
}

// Comparar compara dos arreglos de longitud especificada.
// Devuelve -1 si el primer arreglo es menor que el segundo; 0 si son iguales; o 1 si el primero es el mayor.
// Un arreglo es menor a otro cuando al compararlos elemento a elemento, el primer elemento en el que difieren
// no existe o es menor.
func Comparar(vector1 []int, vector2 []int) int {
	len_minimo := len(vector1)
	if len_minimo > len(vector2) {
		len_minimo = len(vector2)
	}
	for indice := 0; indice < len_minimo; indice++ {
		if vector1[indice] > vector2[indice] {
			return 1
		} else if vector1[indice] < vector2[indice] {
			return -1
		}
	}
	if len(vector1) > len_minimo {
		return 1
	} else if len(vector2) > len_minimo {
		return -1
	}
	return 0
}

// Seleccion ordena el arreglo recibido mediante el algoritmo de selección.
func Seleccion(vector []int) {
	if len(vector) == 0 {
		return
	}
	for indice := len(vector) - 1; indice > 0; indice-- {
		indice_maximo := Maximo(vector[:indice])
		if vector[indice] < vector[indice_maximo] {
			Swap(&vector[indice_maximo], &vector[indice])
		}
	}

}

// Suma devuelve la suma de los elementos de un arreglo. En caso de no tener elementos, debe devolver 0.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func Suma(vector []int) int {
	if len(vector) == 0 {
		return 0
	}
	return vector[0] + Suma(vector[1:])
}

// EsCadenaCapicua devuelve si la cadena es un palíndromo. Es decir, si se lee igual al derecho que al revés.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func EsCadenaCapicua(cadena string) bool {
	if len(cadena) <= 1 {
		return true
	} else if cadena[0] != cadena[len(cadena)-1] {
		return false
	}
	return EsCadenaCapicua(cadena[1 : len(cadena)-1])
}

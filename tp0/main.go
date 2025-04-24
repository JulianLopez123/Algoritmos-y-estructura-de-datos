package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tp0/ejercicios"
)

func armar_vectores(ruta string) []int {
	archivo, _ := os.Open(ruta)
	defer archivo.Close()
	lectura := bufio.NewScanner(archivo)
	vector := []int{}
	for lectura.Scan() {
		linea := lectura.Text()
		numero1, _ := strconv.Atoi(linea)
		vector = append(vector, numero1)
	}
	return vector
}

func mostrar_vector_mayor(vector []int) {
	ejercicios.Seleccion(vector)
	for _, valor := range vector {
		fmt.Println(valor)
	}
}
func main() {
	vector1, vector2 := armar_vectores("archivo1.in"), armar_vectores("archivo2.in")
	vector_mayor := ejercicios.Comparar(vector1, vector2)
	if vector_mayor == 1 {
		mostrar_vector_mayor(vector1)
	} else {
		mostrar_vector_mayor(vector2)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	args := os.Args
	switch args[1]{
	case "ver_tablero":
		printVuelos(args[2])

	}
	
	
}

func printVuelos(ruta string){
	archivo, _ := os.Open(ruta)
	defer archivo.Close()
	lectura := bufio.NewScanner(archivo)
	for lectura.Scan() {
		vuelo := lectura.Text()
		vuelo_sep := strings.Split(vuelo,",")
		vuelos_juntos := strings.Join(vuelo_sep, "	")
		fmt.Println(vuelos_juntos)
	}
}
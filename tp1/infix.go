package main

import (
	"bufio"
	"fmt"
	"os"
	"tp1/conversor"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		linea := input.Text()
		fmt.Println(conversor.ConvertirANotacionPolacaInversa(linea))
	}
}

package tp2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(ruta string){
	archivo, _ := os.Open(ruta)
	defer archivo.Close()
	lectura := bufio.NewScanner(archivo)
	for lectura.Scan() {
		vuelo := lectura.Text()
		vuelo_sep := strings.Split(vuelo,",")
		
		fmt.Println(vuelo_sep)
	}
	
}
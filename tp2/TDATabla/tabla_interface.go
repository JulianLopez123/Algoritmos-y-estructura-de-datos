package TDATabla

type Tabla interface {

	Agregar_archivo([]string)

	Ver_tablero([]string)

	Info_vuelo([]string)

	Siguiente_vuelo([]string)

	Prioridad_vuelos([]string)

	Borrar([]string)

}

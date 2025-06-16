package TDATabla

type Tabla interface {

	//Si  la cantidad de elementos 
	AgregarArchivo([]string)

	VerTablero([]string)

	InfoVuelo([]string)

	SiguienteVuelo([]string)

	PrioridadVuelos([]string)

	Borrar([]string)

	ImprimirError(string)
}

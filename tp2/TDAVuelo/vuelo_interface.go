package TDAVuelo

type Vuelo interface {
	Fecha() string

	Numero_vuelo() string

	Prioridad() int

	Obtener_string() string
}

package TDAVuelo

type Vuelo interface {
	Fecha() string

	Numero_vuelo() string

	Prioridad() int

	Aeropuerto_origen() string

	Aeropuerto_destino() string

	Obtener_string() string
}

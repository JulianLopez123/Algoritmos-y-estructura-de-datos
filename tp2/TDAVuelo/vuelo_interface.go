package TDAVuelo

type Vuelo interface {

	//Devuelve la fecha del vuelo
	Fecha() string

	//Devuelve el numero de vuelo
	NumeroVuelo() string

	//Devuelve la prioridad del vuelo
	Prioridad() int

	//Devuelve el aeropuerto de origen del vuelo
	AeropuertoOrigen() string

	//Devuelve el aeropuerto de destino del vuelo
	AeropuertoDestino() string

	//Devuelve toda la informacion del vuelo en un formato printeable
	ObtenerString() string
}

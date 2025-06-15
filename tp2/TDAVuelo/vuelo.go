package TDAVuelo

import (
	"fmt"
	"strconv"
)

type vuelo struct {
	numero_vuelo       string
	aerolinea          string
	aeropuerto_origen  string
	aeropuerto_destino string
	numero_cola        string
	prioridad          int
	fecha              string
	retraso_salida     int
	tiempo_vuelo       int
	cancelado          int
}

func CrearVuelo(entrada []string) Vuelo {
	//numero_vuelo, _ := strconv.Atoi(entrada[0])
	prioridad, _ := strconv.Atoi(entrada[5])
	retraso_salida, _ := strconv.Atoi(entrada[7])
	tiempo_vuelo, _ := strconv.Atoi(entrada[8])
	cancelado, _ := strconv.Atoi(entrada[9])

	return &vuelo{numero_vuelo: entrada[0], aerolinea: entrada[1], aeropuerto_origen: entrada[2],
		aeropuerto_destino: entrada[3], numero_cola: entrada[4], prioridad: prioridad, fecha: entrada[6],
		retraso_salida: retraso_salida, tiempo_vuelo: tiempo_vuelo, cancelado: cancelado}
}

func (vuelo *vuelo) Fecha() string {
	return vuelo.fecha
}

func (vuelo *vuelo) Numero_vuelo() string {
	return vuelo.numero_vuelo
}

func (vuelo *vuelo) Prioridad() int {
	return vuelo.prioridad
}

func (vuelo *vuelo) Aeropuerto_origen() string {
	return vuelo.aeropuerto_origen
}

func (vuelo *vuelo) Aeropuerto_destino() string {
	return vuelo.aeropuerto_destino
}

func (vuelo *vuelo) Obtener_string() string {
	return fmt.Sprintf("%s %s %s %s %s %d %s %d %d %d",
		vuelo.numero_vuelo,
		vuelo.aerolinea,
		vuelo.aeropuerto_origen,
		vuelo.aeropuerto_destino,
		vuelo.numero_cola,
		vuelo.prioridad,
		vuelo.fecha,
		vuelo.retraso_salida,
		vuelo.tiempo_vuelo,
		vuelo.cancelado,
	)
}

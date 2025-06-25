package TDATabla

type Tabla interface {

	//Agrega o actualiza los vuelos(dependiendo de si son nuevos) desde el archivo en ruta;
	//si el archivo no existe en la ruta recibida, devuelve true. Si no hay error, devuelve false al finalizar.
	AgregarArchivo(ruta string) bool

	//Devuelve cantidad_vuelos vuelos que se encuentren en el rango de fechas recibidas(desde, hasta), ordenados
	//según el modo especificado "asc" o "desc"; si el modo es invalido,o la cantidad_vuelos menor o igual a cero, o desde mayor que hasta,
	//devuelve true. Si no hay error, devuelve false al finalizar.
	VerTablero(cantidad_vuelos int, modo, desde, hasta string) ([]string, bool)

	//Devuelve la informacion del vuelo relacionado al numero_vuelo recibido por parametro,
	//si no se encuentra,devuelve true. Si no hay error, devuelve false al finalizar.
	InfoVuelo(numero_vuelo string) (string, bool)

	//Devuelve el primer vuelo encontrado que parte desde origen hacia destino con fecha igual o posterior a fecha_desde;
	//si no encuentra vuelos, devuelve true. Si no hay error, devuelve false al finalizar.
	SiguienteVuelo(origen, destino, fecha_desde string) (string, bool)

	//Devuelve las prioridades y los numeros de vuelo de la cantidad de vuelos recibida con mayor prioridad; si cantidad_vuelos es menor o igual a cero,
	//devuelve true. Si no hay error, devuelve false al finalizar.
	PrioridadVuelos(cantidad_vuelos int) ([]string, bool)

	//Elimina y devuelve la información de los vuelos en el rango de fechas recibidas;
	//si la fecha inicial es mayor a la final, devuelve true. Si no hay error, devuelve false al finalizar.
	Borrar(fecha_desde, fecha_hasta string) ([]string, bool)
}

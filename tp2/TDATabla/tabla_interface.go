package TDATabla

type Tabla interface {

	//Agrega o actualiza los vuelos(dependiendo de si son nuevos) desde el archivo en ruta;
	//si el archivo no existe en la ruta recibida, imprime error y retorna. Si no hay error, imprime "OK" al finalizar.
	AgregarArchivo(ruta string)

	//Imprime cantidad_vuelos vuelos que se encuentren en el rango de fechas recibidas(desde, hasta), ordenados 
	//según el modo especificado "asc" o "desc"; si el modo es invalido,o la cantidad_vuelos menor o igual a cero, o desde mayor que hasta,
	//imprime error y retorna. Si no hay error, imprime "OK" al finalizar.
	VerTablero(cantidad_vuelos int,modo,desde,hasta string)

	//Imprime la informacion del vuelo relacionado al numero_vuelo recibido por parametro, 
	//si no se encuentra,imprime error y retorna. Si no hay error, imprime "OK" al finalizar.
	InfoVuelo(numero_vuelo string)

	//Imprime el primer vuelo encontrado que parte desde origen hacia destino con fecha igual o posterior a fecha_desde;
	//si no encuentra vuelos, imprime un mensaje indicándolo. Si no hay error, imprime "OK" al finalizar.
	SiguienteVuelo(origen,destino,fecha_desde string)

	//Imprime las prioridades y los numeros de vuelo de la cantidad de vuelos recibida con mayor prioridad; si cantidad_vuelos es menor o igual a cero,
	//imprime error y retorna. Si no hay error, imprime "OK" al finalizar.
	PrioridadVuelos(cantidad_vuelos int)

	//Elimina e imprime la información de los vuelos en el rango de fechas recibidas;
	//si la fecha inicial es mayor a la final, imprime error y retorna. Si no hay error, imprime "OK" al finalizar.
	Borrar(fecha_desde,fecha_hasta string)

	//Imprime por stderr la linea "Error en comando", con el comando enviado por parametro.
	//Si no hay error, imprime "OK" al finalizar.
	ImprimirError(comando string)
}

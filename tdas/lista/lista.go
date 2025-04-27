package lista

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la lista no tiene elementos, false en caso contrario.
	EstaVacia() bool

	// InsetarPrimero agrega un elemento en la primera posicion de la lista.
	InsertarPrimero(T)

	// InsertarUltimo agrega un elemento en la ultima posicion de la lista.
	InsertarUltimo(T)

	// BorrarPrimero saca el primer elemento de la lista. Si la lista tiene elementos,se quita el primero de la misma,
	// y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero devuelve el primer elemento de la lista. Si la lista esta vacia,
	// entra en pánico con un mensaje "La lista esta vacía".
	VerPrimero() T

	// VerUltimo devuelve el ultimo elemento de la lista. Si la lista esta vacia,
	// entra en pánico con un mensaje "La lista esta vacía".
	VerUltimo() T

	// Largo devuelve el largo de la lista, o lo que es equivalente, su cantidad de elementos.
	Largo() int

	// Iterar recorre toda la lista aplicando la funcion visitar elemento a elemento
	// hasta que finaliza la misma o hasta que visitar devuelva false.
	Iterar(visitar func(T) bool)

	// Iterador se encarga de inicializar el iterador externo en la primera posicion de la lista
	// (el usuario es el encargado de realizar la iteracion a conveniencia).
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// VerActual devuelve el dato de la posicion actual en la que se encuentra el iterador; si no
	// hay elemento en la posicion actual entra en panico con un mensaje "El iterador termino de iterar"
	VerActual() T

	// HaySiguiente devuelve true si hay un elemento en la posicion actual del iterador;de no ser asi devuelve false.
	HaySiguiente() bool

	// Siguiente mueve la posicion actual del iterador a la del siguiente elemento.
	// Si no hay siguiente elemento entra en panico con un mensaje "El iterador termino de iterar".
	Siguiente()

	// Insertar agrega un nuevo elemento en la posicion actual del iterador.
	Insertar(T)

	// Borrar elimina el elemento de la posicion actual del iterador. Si la lista no tiene
	// elementos entra en panico con un mensaje "La lista esta vacia""
	Borrar() T
}

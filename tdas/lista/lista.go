package lista

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la lista no tiene elementos, false en caso contrario.
	EstaVacia() bool

	// InsetarPrimero agrega un elemento en la primera posicion de la lista.
	InsertarPrimero(T)

	// InsertarUltimo agrega un elemento en la ultima posicion de la lista.
	InsertarUltimo(T)

	// BorrarPrimero saca el primer elemento de la lista. Si la lista tiene elementos, se quita el primero de la misma,
	// y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	BorrarPrimero() T
	VerPrimero() T
	VerUltimo() T
	Largo() int
	Iterar(visitar func(T) bool)
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {
	VerActual() T

	// HaySiguiente devuelve true si hay un siguiente elemento a leer en la posicion actual del iterador;de no ser asi devuelve false.
	HaySiguiente() bool

	// Siguiente mueve la posicion actual del iterador a la siguiente. Si no hay siguiente elemento a leer entra en panico
	// con un mensaje "El iterador termino de iterar"
	Siguiente()
	Insertar(T)
	Borrar() T
}


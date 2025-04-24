package cola

type nodoCola[T any] struct {
	dato    T
	proximo *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{primero: nil, ultimo: nil}
}

func (cola *colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil
}

func (cola *colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.primero.dato
}

func (cola *colaEnlazada[T]) Desencolar() T {
	elemento_a_desencolar := cola.VerPrimero()
	cola.primero = cola.primero.proximo
	return elemento_a_desencolar
}

func (cola *colaEnlazada[T]) Encolar(elemento T) {
	nodo_nuevo := &nodoCola[T]{dato: elemento}
	if cola.EstaVacia() {
		cola.primero = nodo_nuevo
		cola.ultimo = nodo_nuevo
	} else {
		cola.ultimo.proximo = nodo_nuevo
		cola.ultimo = nodo_nuevo
	}
}

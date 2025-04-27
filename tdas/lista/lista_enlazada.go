package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iterListaEnlazada[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *listaEnlazada[T]
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{primero: nil, ultimo: nil, largo: 0}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(elemento T) {
	nodo_nuevo := &nodoLista[T]{dato: elemento}
	if lista.primero == nil {
		lista.primero = nodo_nuevo
		lista.ultimo = nodo_nuevo
	} else {
		nodo_nuevo.siguiente = lista.primero
		lista.primero = nodo_nuevo
	}
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(elemento T) {
	nodo_nuevo := &nodoLista[T]{dato: elemento}
	if lista.ultimo == nil {
		lista.ultimo = nodo_nuevo
		lista.primero = nodo_nuevo
	} else {
		lista.ultimo.siguiente = nodo_nuevo
		lista.ultimo = nodo_nuevo
	}
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	nodo_eliminado := lista.datoNodo(lista.primero)
	if lista.primero.siguiente != nil {
		lista.primero = lista.primero.siguiente
	} else {
		lista.primero = nil
		lista.ultimo = nil
	}
	lista.largo--
	return nodo_eliminado
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	return lista.datoNodo(lista.primero)
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	return lista.datoNodo(lista.ultimo)
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := lista.primero
	for actual != nil {
		if !visitar(actual.dato) {
			return
		}
		actual = actual.siguiente
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{actual: lista.primero, anterior: nil, lista: lista}
}

func (iter *iterListaEnlazada[T]) VerActual() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.actual.dato
}

func (iter *iterListaEnlazada[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iterListaEnlazada[T]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente
}

func (iter *iterListaEnlazada[T]) Insertar(elemento T) {
	nodo_nuevo := &nodoLista[T]{dato: elemento, siguiente: iter.actual}
	if iter.anterior == nil {
		iter.lista.primero = nodo_nuevo
		if iter.lista.EstaVacia() {
			iter.lista.ultimo = nodo_nuevo
		}
	} else {
		iter.anterior.siguiente = nodo_nuevo
		if iter.actual == nil {
			iter.lista.ultimo = nodo_nuevo
		}
	}
	iter.actual = nodo_nuevo
	iter.lista.largo++
}

func (iter *iterListaEnlazada[T]) Borrar() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo_eliminado := iter.actual
	if iter.anterior == nil {
		iter.lista.primero = iter.actual.siguiente
	} else {
		iter.anterior.siguiente = iter.actual.siguiente
	}
	if nodo_eliminado == iter.lista.ultimo {
		iter.lista.ultimo = iter.anterior
	}
	iter.actual = iter.actual.siguiente
	iter.lista.largo--
	return nodo_eliminado.dato
}

func (lista *listaEnlazada[T]) datoNodo(nodo *nodoLista[T]) T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return nodo.dato
}

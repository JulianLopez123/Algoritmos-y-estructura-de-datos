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
		lista.primero.siguiente = lista.primero
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
	if lista.largo == 0 {
		panic("La lista esta vacia")
	}
	nodo_eliminado := lista.primero.dato
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
	if lista.largo == 0 {
		panic("La lista esta vacia")
	}
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.largo == 0 {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{actual: lista.primero}
}

func (iter *iterListaEnlazada[T]) VerActual() T {
	return iter.actual.dato
}

func (iter *iterListaEnlazada[T]) HaySiguiente() bool {
	return iter.actual.siguiente != nil
}

func (iter *iterListaEnlazada[T]) Siguiente() {
	if iter.actual == nil {
		panic("El iterador termino de iterar")
	}
	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente
}

func (iter *iterListaEnlazada[T]) Insertar(elemento T) {
	nodo_nuevo := &nodoLista[T]{dato: elemento}
	if iter.anterior == nil {
		iter.lista.InsertarPrimero(elemento)
	} else if iter.actual == nil {
		iter.lista.InsertarUltimo(elemento)
	} else {
		nuevo_siguiente := iter.actual
		iter.actual = nodo_nuevo
		iter.actual.siguiente = nuevo_siguiente //si pusiese directamente iter.actual.sig = iter.actual ese iter.actual apuntaria a si mismo¿
		iter.anterior.siguiente = iter.actual
	}
	iter.lista.largo++
}

func (iter *iterListaEnlazada[T]) Borrar() T {
	if iter.actual == nil {
		panic("La lista esta vacia")
	}
	nodo_eliminado := iter.actual
	if iter.anterior == nil {
		iter.lista.primero = iter.actual.siguiente
	} else {
		iter.actual = iter.actual.siguiente
		iter.anterior.siguiente = iter.actual
	}

	if iter.actual.siguiente == nil {
		iter.lista.ultimo = iter.anterior
	} else {
		iter.actual = iter.actual.siguiente
		iter.anterior.siguiente = iter.actual
	}
	return nodo_eliminado.dato
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := lista.primero
	for actual != nil {
		if !visitar(actual.dato) {
			break
		}
		actual = lista.primero.siguiente
	}
}

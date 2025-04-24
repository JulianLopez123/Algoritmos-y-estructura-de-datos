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
    //??? externo
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{primero: nil, ultimo: nil, largo:0}
}

func (lista *listaEnlazada[T]) EstaVacia() bool{
	if lista.largo == 0{
		return true
	}
	return false
}

func (lista *listaEnlazada[T]) InsertarPrimero(elemento T){
	nodo_nuevo := &nodoLista[T]{dato: elemento}
	if lista.primero == nil{
		lista.primero = nodo_nuevo
		lista.ultimo = nodo_nuevo
	}else{
		lista.primero.siguiente = lista.primero
		lista.primero = nodo_nuevo
	}
	lista.largo ++
}

func (lista *listaEnlazada[T]) InsertarUltimo(elemento T){
	nodo_nuevo := &nodoLista[T]{dato: elemento}
	if lista.ultimo == nil{
		lista.ultimo = nodo_nuevo
		lista.primero = nodo_nuevo	
	}else{
		lista.ultimo.siguiente = nodo_nuevo
		lista.ultimo = nodo_nuevo
	}	
	lista.largo ++
}
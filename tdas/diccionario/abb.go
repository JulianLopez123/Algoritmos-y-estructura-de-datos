package diccionario

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	comparar func(K, K) int
}

type nodoAbb[K comparable, V any] struct {
	izq   *nodoAbb[K, V]
	der   *nodoAbb[K, V]
	clave K
	dato  V
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{raiz: nil, cantidad: 0, comparar: funcion_cmp}
}

func (abb *abb[K, V]) Cantidad() int{
	return abb.cantidad
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	abb.raiz = abb.hallarPosicionDeNodo(clave,dato,abb.raiz)
}


func (abb *abb[K, V]) hallarPosicionDeNodo(clave K,dato V, nodo *nodoAbb[K,V]) *nodoAbb[K,V]{
	if nodo == nil{
		abb.cantidad++
		return &nodoAbb[K, V]{clave:clave,dato:dato}
	}
	condicion := abb.comparar(clave,nodo.clave)
	switch {
	case condicion == 0:
	nodo.dato = dato
	case condicion > 0:
	nodo.der = abb.hallarPosicionDeNodo(clave,dato,nodo.der)
	case condicion < 0:
	nodo.izq = abb.hallarPosicionDeNodo(clave,dato,nodo.izq) 
	}
	return nodo
}
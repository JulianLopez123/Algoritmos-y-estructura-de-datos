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

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	abb.raiz = abb.guardar(clave,dato,abb.raiz)
}

func (abb *abb[K, V]) Borrar(clave K) V{
	var dato V;
	abb.raiz ,dato = abb.borrar(clave,abb.raiz)
	abb.cantidad--
	return dato
}

func (abb *abb[K, V]) borrar(clave K, nodo *nodoAbb[K,V]) (*nodoAbb[K,V], V){ 
	if nodo == nil{
		panic("La clave no pertenece al diccionario")
	}
	var dato V;
	condicion := abb.comparar(clave,nodo.clave)
	switch {
	case condicion == 0:
		dato := nodo.dato
		if nodo.izq == nil && nodo.der == nil{
			return nil,dato
		}else if nodo.izq == nil{
			return nodo.der,dato
		}else if nodo.der == nil{
			return nodo.izq,dato
		}else{
			nodo_maximo := buscarMaximo(nodo.izq)
			nodo.clave, nodo.dato = nodo_maximo.clave, nodo_maximo.dato
			nodo.izq, _ = abb.borrar(nodo_maximo.clave,nodo.izq)
			return nodo, dato
		}
	case condicion > 0:
	nodo.der,dato = abb.borrar(clave,nodo.der)
	return nodo, dato
	case condicion < 0:
	nodo.izq,dato = abb.borrar(clave,nodo.izq)
	return nodo, dato
	}
	return nil, nodo.dato
}

func (abb *abb[K, V]) guardar(clave K,dato V, nodo *nodoAbb[K,V]) *nodoAbb[K,V]{
	if nodo == nil{
	abb.raiz = abb.hallarPosicionDeNodo(clave, dato, abb.raiz)
}

func (abb *abb[K, V]) hallarPosicionDeNodo(clave K, dato V, nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	if nodo == nil {
		abb.cantidad++
		return &nodoAbb[K, V]{clave: clave, dato: dato}
	}
	condicion := abb.comparar(clave, nodo.clave)
	switch {
	case condicion == 0:
		nodo.dato = dato
	case condicion > 0:
	nodo.der = abb.guardar(clave,dato,nodo.der)
	case condicion < 0:
	nodo.izq = abb.guardar(clave,dato,nodo.izq) 
	}
	return nodo
}
func buscarMaximo[K comparable,V any](nodo *nodoAbb[K,V]) *nodoAbb[K,V]{
	if nodo.der == nil{
		return nodo
	}
	return buscarMaximo(nodo.der) 
}

func (abb *abb[K, V]) buscarClaveEnArbol(clave K, nodo *nodoAbb[K, V]) bool {
	if nodo == nil {
		return false
	}
	comparacion := abb.comparar(clave, nodo.clave)
	if comparacion == 0 {
		return true
	} else if comparacion < 0 {
		return abb.buscarClaveEnArbol(clave, nodo.izq)
	} else {
		return abb.buscarClaveEnArbol(clave, nodo.der)
	}
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	return abb.buscarClaveEnArbol(clave, abb.raiz)
}

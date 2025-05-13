package diccionario

import (
	TDAPila "tdas/pila"
)

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

type iterAbb[K comparable, V any] struct {
	pila TDAPila.Pila[*nodoAbb[K, V]]
}

type iterRangoAbb[K comparable, V any] struct {
	pila TDAPila.Pila[*nodoAbb[K, V]]
	abb *abb
	desde *K
	hasta *K
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{raiz: nil, cantidad: 0, comparar: funcion_cmp}
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	abb.raiz = abb.hallarPosicionDeNodo(clave, dato, abb.raiz)
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	return abb.buscarClaveEnArbol(clave, abb.raiz)
}

func (abb *abb[K, V]) Obtener(clave K) V {
	if !abb.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
	return abb.buscarDatoEnArbol(clave, abb.raiz)
}

func (abb *abb[K, V]) Borrar(clave K) V {
	var dato V
	abb.raiz, dato = abb.borrar(clave, abb.raiz)
	abb.cantidad--
	return dato
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	pila := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iter := &iterAbb[K, V]{pila: pila}

	for abb.raiz != nil {
		pila.Apilar(abb.raiz)
		abb.raiz = abb.raiz.izq
	}

	return iter
}

func (iter *iterAbb[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia()
}

func (iter *iterAbb[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.pila.VerTope().clave, iter.pila.VerTope().dato
}

func (iter *iterAbb[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo_eliminado := iter.pila.Desapilar()
	nodo_actual := nodo_eliminado.der
	for nodo_actual != nil {
		iter.pila.Apilar(nodo_actual)
		nodo_actual = nodo_actual.izq
	}
}

func  (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V]{
	pila := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	pila,_,_,_ = abb.inicializoPila(pila, abb.raiz, desde,hasta)

	return &iterRangoAbb[K, V]{abb:abb,pila: pila,desde:desde,hasta:hasta}
}

func (abb *abb[K, V])inicializoPila(pila TDAPila.Pila[*nodoAbb[K, V]], nodo *nodoAbb[K,V] , desde *K, hasta *K) (TDAPila.Pila[*nodoAbb[K, V]], nodoAbb[K,V], *K,*K){
	if nodo == nil{
		return pila,*nodo,nil,nil
	}else if abb.comparar(nodo.clave, *desde) >= 0 && abb.comparar(nodo.clave, *hasta) <= 0{
		pila.Apilar(nodo)
	}
	return abb.inicializoPila(pila, nodo.izq , desde, hasta)
}

func (iterRango *iterRangoAbb[K, V]) HaySiguiente() bool{
	return !iterRango.pila.EstaVacia()
}

func (iterRango *iterRangoAbb[K, V]) VerActual() (K, V){
	if !iterRango.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iterRango.pila.VerTope().clave, iterRango.pila.VerTope().dato
}

func (iterRango *iterRangoAbb[K, V]) Siguiente(){
	if !iterRango.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo_eliminado := iterRango.pila.Desapilar()
	nodo_actual := nodo_eliminado.der
	for nodo_actual != nil {
		clave := nodo_actual.clave
		if iterRango.condicion(clave){
			iterRango.pila.Apilar(nodo_actual)
		}
		nodo_actual = nodo_actual.izq
	}	
}

func (iterRango *iterRangoAbb[K, V]) condicion(clave K) bool{
	if iterRango.abb.comparar(clave, iterRango.desde) >= 0 && iterRango.abb.comparar(clave, iterRango.hasta) <= 0{
		return true
	}
	return false
}

func (abb abb[K, V]) Iterar(visitar func(K, V) bool){
	abb.raiz.iterar(visitar)
}

func (nodo *nodoAbb[K, V]) iterar(visitar func(K, V) bool) {
	if nodo == nil {
		return
	}
	nodo.izq.iterar(visitar)
	if !visitar(nodo.clave, nodo.dato) {
		return
	}
	nodo.der.iterar(visitar)
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
		nodo.der = abb.hallarPosicionDeNodo(clave, dato, nodo.der)
	case condicion < 0:
		nodo.izq = abb.hallarPosicionDeNodo(clave, dato, nodo.izq)
	}
	return nodo
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

func (abb *abb[K, V]) buscarDatoEnArbol(clave K, nodo *nodoAbb[K, V]) V {
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	comparacion := abb.comparar(clave, nodo.clave)
	if comparacion == 0 {
		return nodo.dato
	} else if comparacion < 0 {
		return abb.buscarDatoEnArbol(clave, nodo.izq)
	} else {
		return abb.buscarDatoEnArbol(clave, nodo.der)
	}
}

func (abb *abb[K, V]) borrar(clave K, nodo *nodoAbb[K, V]) (*nodoAbb[K, V], V) {
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	var dato V
	condicion := abb.comparar(clave, nodo.clave)
	switch {
	case condicion == 0:
		dato := nodo.dato
		if nodo.izq == nil && nodo.der == nil {
			return nil, dato
		} else if nodo.izq == nil {
			return nodo.der, dato
		} else if nodo.der == nil {
			return nodo.izq, dato
		} else {
			nodo_maximo := buscarMaximo(nodo.izq)
			nodo.clave, nodo.dato = nodo_maximo.clave, nodo_maximo.dato
			nodo.izq, _ = abb.borrar(nodo_maximo.clave, nodo.izq)
			return nodo, dato
		}
	case condicion > 0:
		nodo.der, dato = abb.borrar(clave, nodo.der)
		return nodo, dato
	case condicion < 0:
		nodo.izq, dato = abb.borrar(clave, nodo.izq)
		return nodo, dato
	}
	return nil, nodo.dato
}

func buscarMaximo[K comparable, V any](nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	if nodo.der == nil {
		return nodo
	}
	return buscarMaximo(nodo.der)
}

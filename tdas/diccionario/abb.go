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
	pila  TDAPila.Pila[*nodoAbb[K, V]]
	abb   *abb[K, V]
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
	return abb.buscarNodoEnArbol(clave, abb.raiz) != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	nodo := abb.buscarNodoEnArbol(clave, abb.raiz)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return nodo.dato
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
	nodo := abb.raiz
	for nodo != nil {
		iter.pila.Apilar(nodo)
		nodo = nodo.izq
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

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	pila := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iterador := &iterRangoAbb[K, V]{abb: abb, pila: pila, desde: desde, hasta: hasta}
	iterador.apilarElementosEnRango(abb.raiz)
	return iterador
}

func (iterRango *iterRangoAbb[K, V]) HaySiguiente() bool {
	return !iterRango.pila.EstaVacia()
}

func (iterRango *iterRangoAbb[K, V]) VerActual() (K, V) {
	if !iterRango.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iterRango.pila.VerTope().clave, iterRango.pila.VerTope().dato
}

func (iterRango *iterRangoAbb[K, V]) Siguiente() {
	if !iterRango.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo_eliminado := iterRango.pila.Desapilar()
	nodo_actual := nodo_eliminado.der
	iterRango.apilarElementosEnRango(nodo_actual)
}

func (iterRango *iterRangoAbb[K, V]) apilarElementosEnRango(nodo *nodoAbb[K, V]) {
	for nodo != nil {
		if iterRango.abb.comparar(nodo.clave, *iterRango.desde) < 0 {
			nodo = nodo.der
		} else if iterRango.abb.comparar(nodo.clave, *iterRango.hasta) > 0 {
			nodo = nodo.izq
		} else {
			iterRango.pila.Apilar(nodo)
			nodo = nodo.izq
		}
	}
}

func (abb abb[K, V]) Iterar(visitar func(K, V) bool) {
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

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	abb.iterarRango(abb.raiz, visitar, desde, hasta)
}

func (abb *abb[K, V]) iterarRango(n *nodoAbb[K, V], visitar func(K, V) bool, desde *K, hasta *K) {
	if n == nil {
		return
	}
	if desde == nil || abb.comparar(*desde, n.clave) < 0 {
		abb.iterarRango(n.izq, visitar, desde, hasta)
	}
	if (desde == nil || abb.comparar(*desde, n.clave) <= 0) && (hasta == nil || abb.comparar(*hasta, n.clave) >= 0) {
		if !visitar(n.clave, n.dato) {
			return
		}
	}
	if hasta == nil || abb.comparar(*hasta, n.clave) > 0 {
		abb.iterarRango(n.der, visitar, desde, hasta)
	}
}

func (abb *abb[K, V]) hallarPosicionDeNodo(clave K, dato V, nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	if nodo == nil {
		abb.cantidad++
		return &nodoAbb[K, V]{clave: clave, dato: dato}
	}
	rangoValido := abb.comparar(clave, nodo.clave)
	switch {
	case rangoValido == 0:
		nodo.dato = dato
	case rangoValido > 0:
		nodo.der = abb.hallarPosicionDeNodo(clave, dato, nodo.der)
	case rangoValido < 0:
		nodo.izq = abb.hallarPosicionDeNodo(clave, dato, nodo.izq)
	}
	return nodo
}

func (abb *abb[K, V]) buscarNodoEnArbol(clave K, nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	if nodo == nil {
		return nil
	}
	comparacion := abb.comparar(clave, nodo.clave)
	if comparacion == 0 {
		return nodo
	} else if comparacion < 0 {
		return abb.buscarNodoEnArbol(clave, nodo.izq)
	} else {
		return abb.buscarNodoEnArbol(clave, nodo.der)
	}
}

func (abb *abb[K, V]) borrar(clave K, nodo *nodoAbb[K, V]) (*nodoAbb[K, V], V) {
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	var dato V
	rangoValido := abb.comparar(clave, nodo.clave)
	switch {
	case rangoValido == 0:
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
	case rangoValido > 0:
		nodo.der, dato = abb.borrar(clave, nodo.der)
		return nodo, dato
	case rangoValido < 0:
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

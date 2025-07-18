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

type iterRangoAbb[K comparable, V any] struct {
	pila        TDAPila.Pila[*nodoAbb[K, V]]
	comparacion func(K, K) int
	desde       *K
	hasta       *K
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{raiz: nil, cantidad: 0, comparar: funcion_cmp}
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	if abb.raiz == nil {
		abb.raiz = crearNodo(clave, dato)
		abb.cantidad++
		return
	}

	nodo, padre := abb.buscarNodoYPadreEnArbol(clave, abb.raiz, nil)
	if nodo == nil {
		nodo = crearNodo(clave, dato)
		condicion := abb.comparar(nodo.clave, padre.clave)
		if condicion < 0 {
			padre.izq = nodo
		} else {
			padre.der = nodo
		}
		abb.cantidad++
	}
	nodo.dato = dato
}
func (abb *abb[K, V]) Pertenece(clave K) bool {
	nodo, _ := abb.buscarNodoYPadreEnArbol(clave, abb.raiz, nil)
	return nodo != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	nodo, _ := abb.buscarNodoYPadreEnArbol(clave, abb.raiz, nil)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return nodo.dato
}

func (abb *abb[K, V]) Borrar(clave K) V {
	nodo, padre := abb.buscarNodoYPadreEnArbol(clave, abb.raiz, nil)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	dato := nodo.dato
	nodo_reemplazante := abb.borrar(nodo)

	if padre == nil {
		abb.raiz = nodo_reemplazante
	} else {
		condicion := abb.comparar(nodo.clave, padre.clave)
		if condicion < 0 {
			padre.izq = nodo_reemplazante
		} else {
			padre.der = nodo_reemplazante
		}
	}
	abb.cantidad--
	return dato
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return abb.IteradorRango(nil, nil)
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	pila := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iterador := &iterRangoAbb[K, V]{comparacion: abb.comparar, pila: pila, desde: desde, hasta: hasta}
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

func (abb *abb[K, V]) Iterar(visitar func(K, V) bool) {
	abb.IterarRango(nil, nil, visitar)
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	abb.iterarRango(abb.raiz, visitar, desde, hasta)
}

func (abb *abb[K, V]) iterarRango(n *nodoAbb[K, V], visitar func(K, V) bool, desde *K, hasta *K) bool {
	if n == nil {
		return true
	}
	if desde == nil || abb.comparar(n.clave, *desde) > 0 {
		if !abb.iterarRango(n.izq, visitar, desde, hasta) {
			return false
		}
	}
	if (desde == nil || abb.comparar(n.clave, *desde) >= 0) && (hasta == nil || abb.comparar(n.clave, *hasta) <= 0) {
		if !visitar(n.clave, n.dato) {
			return false
		}
	}
	if hasta == nil || abb.comparar(n.clave, *hasta) < 0 {
		if !abb.iterarRango(n.der, visitar, desde, hasta) {
			return false
		}
	}
	return true
}

func (iterRango *iterRangoAbb[K, V]) apilarElementosEnRango(nodo *nodoAbb[K, V]) {
	if nodo == nil {
		return
	}
	if iterRango.desde != nil && iterRango.comparacion(nodo.clave, *iterRango.desde) < 0 {
		iterRango.apilarElementosEnRango(nodo.der)
	}
	if (iterRango.desde == nil || iterRango.comparacion(nodo.clave, *iterRango.desde) >= 0) && (iterRango.hasta == nil || iterRango.comparacion(nodo.clave, *iterRango.hasta) <= 0) {
		iterRango.pila.Apilar(nodo)
		iterRango.apilarElementosEnRango(nodo.izq)
	}
	if iterRango.hasta != nil && iterRango.comparacion(nodo.clave, *iterRango.hasta) > 0 {
		iterRango.apilarElementosEnRango(nodo.izq)
	}
}

func (abb *abb[K, V]) borrar(nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	if nodo.izq == nil {
		return nodo.der
	} else if nodo.der == nil {
		return nodo.izq
	}
	nodo_maximo, padre := buscarMaximoYPadre(nodo.izq, nodo)
	nodo.clave, nodo.dato = nodo_maximo.clave, nodo_maximo.dato
	if padre == nodo {
		nodo.izq = nodo_maximo.izq
	} else {
		padre.der = nodo_maximo.izq
	}
	return nodo

}

func (abb *abb[K, V]) buscarNodoYPadreEnArbol(clave K, nodo *nodoAbb[K, V], padre *nodoAbb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if nodo == nil {
		return nil, padre
	}
	comparacion := abb.comparar(clave, nodo.clave)
	if comparacion == 0 {
		return nodo, padre
	} else if comparacion > 0 {
		return abb.buscarNodoYPadreEnArbol(clave, nodo.der, nodo)
	} else {
		return abb.buscarNodoYPadreEnArbol(clave, nodo.izq, nodo)
	}
}

func buscarMaximoYPadre[K comparable, V any](nodo, padre *nodoAbb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if nodo.der == nil {
		return nodo, padre
	}
	return buscarMaximoYPadre(nodo.der, nodo)
}

func crearNodo[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{clave: clave, dato: dato}
}

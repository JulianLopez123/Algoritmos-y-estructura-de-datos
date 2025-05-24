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
	pila  TDAPila.Pila[*nodoAbb[K, V]]
	abb   *abb[K, V]
	desde *K
	hasta *K
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{raiz: nil, cantidad: 0, comparar: funcion_cmp}
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	encontrado := abb.Pertenece(clave)
	nodo := abb.buscarNodoEnArbol(clave, &dato, abb.raiz)
	if !encontrado {
		abb.cantidad++
	}
	abb.raiz = nodo
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	return abb.buscarNodoEnArbol(clave, nil, abb.raiz) != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	nodo := abb.buscarNodoEnArbol(clave, nil, abb.raiz)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return nodo.dato
}

// var dato V
// abb.raiz, dato = abb.borrar(clave, abb.raiz)
// abb.cantidad--
// return dato
func (abb *abb[K, V]) Borrar(clave K) V {
	nodo := abb.buscar(clave, abb.raiz)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	nodo_eliminado := nodo.dato
	abb.cantidad--
	// if nodo.izq == nil && nodo.der == nil {
	// 	nodo = nil
	// 	return nodo_eliminado
	// } else if nodo.izq == nil {
	// 	nodo.clave, nodo.dato = nodo.der.clave, nodo.der.dato
	// 	return nodo_eliminado
	// } else if nodo.der == nil {
	// 	nodo.clave, nodo.dato = nodo.izq.clave, nodo.izq.dato
	// 	return nodo_eliminado
	// } else {
	// 	nodo_maximo := buscarMaximo(nodo.izq)
	// 	nodo.clave, nodo.dato = nodo_maximo.clave, nodo_maximo.dato
	// 	abb.Borrar(nodo_maximo.clave)
	// 	return nodo_eliminado
	// }
	if nodo.izq != nil && nodo.der != nil {
		nodo_maximo := buscarMaximo(nodo.izq)
		nodo.clave, nodo.dato = nodo_maximo.clave, nodo_maximo.dato
		abb.Borrar(nodo_maximo.clave)
		return nodo_eliminado
	}

}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return abb.IteradorRango(nil, nil)
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
	if iterRango.desde != nil && iterRango.abb.comparar(nodo.clave, *iterRango.desde) < 0 {
		iterRango.apilarElementosEnRango(nodo.der)
	}
	if (iterRango.desde == nil || iterRango.abb.comparar(nodo.clave, *iterRango.desde) >= 0) && (iterRango.hasta == nil || iterRango.abb.comparar(nodo.clave, *iterRango.hasta) <= 0) {
		iterRango.pila.Apilar(nodo)
		iterRango.apilarElementosEnRango(nodo.izq)
	}
	if iterRango.hasta != nil && iterRango.abb.comparar(nodo.clave, *iterRango.hasta) > 0 {
		iterRango.apilarElementosEnRango(nodo.izq)
	}
}

func (abb *abb[K, V]) buscarNodoEnArbol(clave K, dato *V, nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	if nodo == nil {
		if dato != nil {
			return &nodoAbb[K, V]{clave: clave, dato: *dato}
		}
		return nil
	}
	comparacion := abb.comparar(clave, nodo.clave)
	switch {
	case comparacion == 0 && dato != nil:
		nodo.dato = *dato
	case comparacion > 0:
		nodo_der := abb.buscarNodoEnArbol(clave, dato, nodo.der)
		if dato == nil {
			return nodo_der
		}
		nodo.der = nodo_der
	case comparacion < 0:
		nodo_izq := abb.buscarNodoEnArbol(clave, dato, nodo.izq)
		if dato == nil {
			return nodo_izq
		}
		nodo.izq = nodo_izq
	}
	return nodo
}

// func (abb *abb[K, V]) borrar(clave K, nodo *nodoAbb[K, V]) (*nodoAbb[K, V], V) {
// 	if nodo == nil {
// 		panic("La clave no pertenece al diccionario")
// 	}
// 	var dato V
// 	rangoValido := abb.comparar(clave, nodo.clave)
// 	switch {
// 	case rangoValido == 0:
// 		dato := nodo.dato
// 		if nodo.izq == nil && nodo.der == nil {
// 			return nil, dato
// 		} else if nodo.izq == nil {
// 			return nodo.der, dato
// 		} else if nodo.der == nil {
// 			return nodo.izq, dato
// 		} else {
// 			nodo_maximo := buscarMaximo(nodo.izq)
// 			nodo.clave, nodo.dato = nodo_maximo.clave, nodo_maximo.dato
// 			nodo.izq, _ = abb.borrar(nodo_maximo.clave, nodo.izq)
// 			return nodo, dato
// 		}
// 	case rangoValido > 0:
// 		nodo.der, dato = abb.borrar(clave, nodo.der)
// 		return nodo, dato
// 	case rangoValido < 0:
// 		nodo.izq, dato = abb.borrar(clave, nodo.izq)
// 		return nodo, dato
// 	}
// 	return nil, nodo.dato
// }

func (abb *abb[K, V]) buscar(clave K, nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	if nodo == nil {
		return nil
	}
	comparacion := abb.comparar(clave, nodo.clave)
	switch {
	case comparacion == 0:
		return nodo
	case comparacion > 0:
		return abb.buscar(clave, nodo.der)
	default:
		return abb.buscar(clave, nodo.izq)
	}
}

func buscarMaximo[K comparable, V any](nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	if nodo.der == nil {
		return nodo
	}
	return buscarMaximo(nodo.der)
}

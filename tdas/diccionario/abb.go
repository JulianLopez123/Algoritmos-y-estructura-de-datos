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
type operacion int

const (
	GUARDAR operacion = iota
	BORRAR
	PERTENECE_U_OBTENER
)

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{raiz: nil, cantidad: 0, comparar: funcion_cmp}
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	nodo, hallado := abb.buscarNodoEnArbol(clave, &dato, abb.raiz, GUARDAR)
	if hallado == nil {
		abb.cantidad++
	}
	abb.raiz = nodo
}
func (abb *abb[K, V]) Pertenece(clave K) bool {
	nodo, _ := abb.buscarNodoEnArbol(clave, nil, abb.raiz, PERTENECE_U_OBTENER)
	return nodo != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	nodo, _ := abb.buscarNodoEnArbol(clave, nil, abb.raiz, PERTENECE_U_OBTENER)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return nodo.dato
}

func (abb *abb[K, V]) Borrar(clave K) V {
	nodo, dato := abb.buscarNodoEnArbol(clave, nil, abb.raiz, BORRAR)
	if dato == nil {
		panic("La clave no pertenece al diccionario")
	}
	abb.raiz = nodo
	abb.cantidad--
	return *dato
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

func (abb *abb[K, V]) buscarNodoEnArbol(clave K, dato *V, nodo *nodoAbb[K, V], operacion operacion) (*nodoAbb[K, V], *V) {
	if nodo == nil {
		if operacion == GUARDAR {
			return crearNodo(clave, *dato), nil
		}
		return nil, nil
	}
	comparacion := abb.comparar(clave, nodo.clave)
	switch {
	case comparacion == 0:
		switch operacion {
		case GUARDAR:
			nodo.dato = *dato
			return nodo, &nodo.dato
		case BORRAR:
			nodo_eliminado, dato := abb.borrar(nodo)
			return nodo_eliminado, &dato
		default:
			return nodo, &nodo.dato
		}
	case comparacion > 0:
		nodo_der, resultado := abb.buscarNodoEnArbol(clave, dato, nodo.der, operacion)
		if operacion == PERTENECE_U_OBTENER {
			return nodo_der, nil
		}
		nodo.der = nodo_der
		return nodo, resultado

	case comparacion < 0:
		nodo_izq, resultado := abb.buscarNodoEnArbol(clave, dato, nodo.izq, operacion)
		if operacion == PERTENECE_U_OBTENER {
			return nodo_izq, nil
		}
		nodo.izq = nodo_izq
		return nodo, resultado
	}
	return nodo, nil
}

func (abb *abb[K, V]) borrar(nodo *nodoAbb[K, V]) (*nodoAbb[K, V], V) {
	var nul V
	if nodo == nil {
		return nil, nul
	}
	dato := nodo.dato
	if nodo.izq == nil {
		return nodo.der, dato
	} else if nodo.der == nil {
		return nodo.izq, dato
	} else {
		predecesor := buscarMaximo(nodo.izq)
		nodo.clave, nodo.dato = predecesor.clave, predecesor.dato
		nodo.izq, _ = abb.buscarNodoEnArbol(nodo.clave, nil, nodo.izq, BORRAR)
		return nodo, dato
	}
}

func buscarMaximo[K comparable, V any](nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	if nodo.der == nil {
		return nodo
	}
	return buscarMaximo(nodo.der)
}

func crearNodo[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{clave: clave, dato: dato}
}

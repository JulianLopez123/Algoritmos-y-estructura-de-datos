package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

const TAMAÑO_HASH = 20

type parClaveValor[K comparable, V any] struct {
	clave K
	dato  V
}

type hashAbierto[K comparable, V any] struct {
	tabla    []TDALista.Lista[parClaveValor[K, V]]
	tam      int
	cantidad int
}

type iterHash[K comparable, V any] struct {
	actual   *parClaveValor[K, V]
	anterior *parClaveValor[K, V]
	lista    *hashAbierto[K, V]
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	tabla := make([]TDALista.Lista[parClaveValor[K, V]], TAMAÑO_HASH)
	for i := range tabla {
		tabla[i] = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
	}
	return &hashAbierto[K, V]{tabla: tabla, tam: TAMAÑO_HASH, cantidad: 0}
}

func (hash *hashAbierto[K, V]) Guardar(clave K, dato V) {
	hashIndex := hashFunc(clave)
	index := hashIndex % hash.tam

	hash.tabla[index].InsertarUltimo(parClaveValor[K, V]{clave: clave, dato: dato})
	hash.cantidad++
}

func (hash *hashAbierto[K, V]) Cantidad() int {
	return hash.cantidad
}

func(hash *hashAbierto[K, V]) Obtener(clave K) V {
	iterador_lista := hash.hallarPosicionParClaveValor(clave)
	return iterador_lista.VerActual().dato
}

func(hash *hashAbierto[K, V]) Borrar(clave K) V {
	iterador_lista := hash.hallarPosicionParClaveValor(clave)
	hash.cantidad --
	return iterador_lista.Borrar().dato
}

func (hash *hashAbierto[K, V])hallarPosicionParClaveValor(clave K) TDALista.IteradorLista[parClaveValor[K, V]]{
	if !clave.Pertenece(){
		panic("La clave no pertenece al diccionario")
	}
	indice := hashFunc(clave) % hash.tam
	lista := hash.tabla[indice]
	iterador_lista := lista.Iterador()
	for iterador_lista.HaySiguiente() {
		if iterador_lista.VerActual().clave == clave{
			return iterador_lista
		}
		iterador_lista.Siguiente()
	}
	panic("error de ejecucion")
}

func hashFunc[K comparable](clave K) int {
	bytes := convertirABytes(clave)
	h := uint64(14695981039346656037)
	for _, b := range bytes {
		h *= 1099511628211
		h ^= uint64(b)
	}
	return int(h)
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func (iter *iterHash[K, V]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iterHash[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar.")
	}
	return iter.actual.clave, iter.actual.dato
}

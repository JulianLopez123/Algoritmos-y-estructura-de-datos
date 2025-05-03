package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

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
	// actual   *parClaveValor[K, V]
	// anterior *parClaveValor[K, V]
	// lista    *hashAbierto[K, V]
	index   int
	hashMap *hashAbierto[K, V]
	lista   TDALista.IteradorLista[parClaveValor[K, V]]
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	tabla := make([]TDALista.Lista[parClaveValor[K, V]], 5)
	for i := range tabla {
		tabla[i] = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
	}
	return &hashAbierto[K, V]{tabla: tabla, tam: 5, cantidad: 0}
}

func (hash *hashAbierto[K, V]) Guardar(clave K, dato V) {
	hashIndex := hashFunc(clave)
	index := hashIndex % hash.tam
	lista := hash.tabla[index]

	iter := lista.Iterador()
	for iter.HaySiguiente() {
		if iter.VerActual().clave == clave {
			iter.Borrar()
			iter.Insertar(parClaveValor[K, V]{clave: clave, dato: dato})
			return
		}
		iter.Siguiente()
	}

	hash.tabla[index].InsertarUltimo(parClaveValor[K, V]{clave: clave, dato: dato})
	hash.cantidad++
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

func (hash *hashAbierto[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashAbierto[K, V]) Pertenece(clave K) bool {
	hashIndex := hashFunc(clave)
	index := hashIndex % hash.tam
	lista := hash.tabla[index]
	iter := lista.Iterador()

	for iter.HaySiguiente() {
		if iter.VerActual().clave == clave {
			return true
		}
		iter.Siguiente()
	}

	return false
}

func (iter *iterHash[K, V]) HaySiguiente() bool {
	return iter.index == iter.hashMap.cantidad
}

func (iter *iterHash[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar.")
	}
	return iter.lista.VerActual().clave, iter.lista.VerActual().dato
}

func (iter *iterHash[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador terminó de iterar.")
	}

}

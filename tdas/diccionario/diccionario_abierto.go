package diccionario

import (
	"fmt"
	"math"
	TDALista "tdas/lista"
)

const TAMAÑO_INICIAL_HASH = 17
const FACTOR_CARGA = 3
const FACTOR_REDIMENSION = 2

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
	index   int
	hashMap *hashAbierto[K, V]
	lista   TDALista.IteradorLista[parClaveValor[K, V]]
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	tabla := inicializarTabla[K, V](TAMAÑO_INICIAL_HASH)
	return &hashAbierto[K, V]{tabla: tabla, tam: TAMAÑO_INICIAL_HASH, cantidad: 0}
}

func (hash *hashAbierto[K, V]) Guardar(clave K, dato V) {
	if hash.cantidad/hash.tam > FACTOR_CARGA {
		hash.redimensionar(hash.tam * FACTOR_REDIMENSION)
	}
	iterador_lista := hash.hallarPosicionParClaveValor(clave)
	if !iterador_lista.HaySiguiente() {
		hash.cantidad++
	} else {
		iterador_lista.Borrar()
	}
	iterador_lista.Insertar(parClaveValor[K, V]{clave: clave, dato: dato})
}

func (hash *hashAbierto[K, V]) Pertenece(clave K) bool {
	iterador_lista := hash.hallarPosicionParClaveValor(clave)
	return iterador_lista.HaySiguiente()

}

func (hash *hashAbierto[K, V]) Obtener(clave K) V {
	iterador_lista := hash.hallarPosicionParClaveValor(clave)
	if !iterador_lista.HaySiguiente() {
		panic("La clave no pertenece al diccionario")
	}
	return iterador_lista.VerActual().dato
}

func (hash *hashAbierto[K, V]) Borrar(clave K) V {
	if hash.cantidad/hash.tam < FACTOR_CARGA {
		hash.redimensionar(hash.tam / FACTOR_REDIMENSION)
	}
	iterador_lista := hash.hallarPosicionParClaveValor(clave)
	if !iterador_lista.HaySiguiente() {
		panic("La clave no pertenece al diccionario")
	}
	hash.cantidad--
	return iterador_lista.Borrar().dato
}

func (hash *hashAbierto[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashAbierto[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for i := 0; i < hash.tam; i++ {
		if hash.tabla[i].EstaVacia() {
			continue
		} else {
			iter := hash.tabla[i].Iterador()
			for iter.HaySiguiente() {
				if !visitar(iter.VerActual().clave, iter.VerActual().dato) {
					return
				}
				iter.Siguiente()
			}
		}
	}
}

func (hash *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	iter := &iterHash[K, V]{hashMap: hash}

	for i := 0; i < len(hash.tabla); i++ {
		if !hash.tabla[i].EstaVacia() {
			iter.index = i
			iter.lista = hash.tabla[i].Iterador()
			break
		}
	}

	return iter
}

func (iter *iterHash[K, V]) HaySiguiente() bool {
	if iter.index == iter.hashMap.tam {
		return false
	}
	return iter.lista != nil && iter.lista.HaySiguiente()
}

func (iter *iterHash[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.lista.VerActual().clave, iter.lista.VerActual().dato
}

func (iter *iterHash[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	iter.lista.Siguiente()
	if iter.lista.HaySiguiente() {
		return
	}

	iter.index++
	for iter.index < iter.hashMap.tam {
		if !iter.hashMap.tabla[iter.index].EstaVacia() {
			iter.lista = iter.hashMap.tabla[iter.index].Iterador()
			return
		}
		iter.index++
	}

	iter.lista = nil
}

func (hash *hashAbierto[K, V]) hallarPosicionParClaveValor(clave K) TDALista.IteradorLista[parClaveValor[K, V]] {
	indice := int(math.Abs(float64(hashFunc(clave) % hash.tam)))
	lista := hash.tabla[indice]
	iterador_lista := lista.Iterador()
	for iterador_lista.HaySiguiente() {
		if iterador_lista.VerActual().clave == clave {
			return iterador_lista
		}
		iterador_lista.Siguiente()
	}
	return iterador_lista
}

func (hash *hashAbierto[K, V]) redimensionar(nuevoTam int) {
	nuevoTam = tamañoPrimo(nuevoTam)
	nuevoHash := inicializarTabla[K, V](nuevoTam)

	for _, lista := range hash.tabla {
		iter := lista.Iterador()
		for iter.HaySiguiente() {
			nuevoIndex := int(math.Abs(float64(hashFunc(iter.VerActual().clave) % nuevoTam)))
			nuevoHash[nuevoIndex].InsertarUltimo(iter.VerActual())
			iter.Siguiente()
		}
	}

	hash.tabla = nuevoHash
	hash.tam = nuevoTam
}

func inicializarTabla[K comparable, V any](tamaño int) []TDALista.Lista[parClaveValor[K, V]] {
	tabla := make([]TDALista.Lista[parClaveValor[K, V]], tamaño)
	for i := range tabla {
		tabla[i] = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
	}
	return tabla
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

func esPrimo(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if (n % i) == 0 {
			return false
		}
	}
	return true
}

func tamañoPrimo(n int) int {
	for !esPrimo(n) {
		n++
	}
	return n
}

package diccionario

import (
	"fmt"
	"math"
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
	index   int
	hashMap *hashAbierto[K, V]
	lista   TDALista.IteradorLista[parClaveValor[K, V]]
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
	index :=  int(math.Abs(float64(hashIndex % hash.tam)))
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

func (hash *hashAbierto[K, V]) Pertenece(clave K) bool {
	hashIndex := hashFunc(clave)
	index := int(math.Abs(float64(hashIndex % hash.tam)))
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

func (hash *hashAbierto[K, V]) Obtener(clave K) V {
	iterador_lista := hash.hallarPosicionParClaveValor(clave)
	return iterador_lista.VerActual().dato
}

func (hash *hashAbierto[K, V]) Borrar(clave K) V {
	iterador_lista := hash.hallarPosicionParClaveValor(clave)
	hash.cantidad--
	return iterador_lista.Borrar().dato
}

func (hash *hashAbierto[K, V]) Cantidad() int {
	return hash.cantidad
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
	panic("La clave no pertenece al diccionario")
	
}

func(hash *hashAbierto[K, V])Iterar(visitar func(clave K,dato V) bool){ 
	for i:= 0; i < hash.tam; i++{
		if hash.tabla[i].EstaVacia() {
			continue
		}else{
			iter := hash.tabla[i].Iterador()
			for iter.HaySiguiente(){
				if !visitar(iter.VerActual().clave,iter.VerActual().dato){
					break
				}
				iter.Siguiente()
			}
		}
	}
}
func (hash *hashAbierto[K, V])Iterador() IterDiccionario[K, V]{
	return &iterHash[K,V ]{index:0 , hashMap:hash}
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

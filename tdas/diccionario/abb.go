package diccionario

type abb[K comparable, V any] struct {
	raiz     *nodoAb[K, V]
	cantidad int
	comparar func(K, K) int
	// destruir dato¿?
}

type nodoAb[K comparable, V any] struct {
	izq   *nodoAb[K, V]
	der   *nodoAb[K, V]
	clave K
	dato  V
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{raiz: nil, cantidad: 0, comparar: funcion_cmp}
}



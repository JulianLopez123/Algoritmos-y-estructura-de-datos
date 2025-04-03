package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func (pila *pilaDinamica[T]) Apilar(elemento T) {
	if len(pila.datos) == pila.cantidad {
		redimensionar("*", pila)
	}
	pila.datos[pila.cantidad] = elemento
	pila.cantidad += 1
}

func (pila *pilaDinamica[T]) Desapilar() T {
	elemento_a_desapilar := pila.VerTope()
	pila.cantidad -= 1
	if len(pila.datos) >= 4*pila.cantidad {
		redimensionar("/", pila)
	}
	return elemento_a_desapilar
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	return pila.datos[pila.cantidad-1]
}

func redimensionar[T any](caso string, pila *pilaDinamica[T]) {
	var nuevo_slice []T
	if caso == "*" {
		nuevo_slice = make([]T, cap(pila.datos)*2)
	} else if caso == "/" {
		if cap(pila.datos)/2 < 1 {
			nuevo_slice = make([]T, 1)
		} else {
			nuevo_slice = make([]T, cap(pila.datos)/2)
		}
	}
	copy(nuevo_slice, pila.datos)
	pila.datos = nuevo_slice

}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{datos: make([]T, 1), cantidad: 0}
}

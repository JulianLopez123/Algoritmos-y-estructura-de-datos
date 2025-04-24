package pila

/* Definición del struct pila proporcionado por la cátedra. */
const PROPORCION_DE_OCUPACION_MINIMA = 4
const CONSTANTE_DE_REDIMENSION = 2
const CAPACIDAD_MINIMA = 5

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{datos: make([]T, CAPACIDAD_MINIMA), cantidad: 0}
}

func (pila *pilaDinamica[T]) Apilar(elemento T) {
	if len(pila.datos) == pila.cantidad {
		nuevo_tamaño := cap(pila.datos) * CONSTANTE_DE_REDIMENSION
		redimensionar(nuevo_tamaño, pila)
	}
	pila.datos[pila.cantidad] = elemento
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	elemento_a_desapilar := pila.VerTope()
	pila.cantidad--
	if len(pila.datos) >= PROPORCION_DE_OCUPACION_MINIMA*pila.cantidad {
		var nuevo_tamaño int
		capacidad_pila_datos := cap(pila.datos)
		if capacidad_pila_datos/CONSTANTE_DE_REDIMENSION < CAPACIDAD_MINIMA {
			nuevo_tamaño = CAPACIDAD_MINIMA
		} else {
			nuevo_tamaño = capacidad_pila_datos / CONSTANTE_DE_REDIMENSION
		}
		redimensionar(nuevo_tamaño, pila)
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

func redimensionar[T any](nuevo_tamaño int, pila *pilaDinamica[T]) {
	nuevo_slice := make([]T, nuevo_tamaño)
	copy(nuevo_slice, pila.datos)
	pila.datos = nuevo_slice
}

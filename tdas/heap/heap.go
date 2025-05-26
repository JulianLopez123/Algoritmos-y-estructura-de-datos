package cola_prioridad

type colaConPrioridad[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &colaConPrioridad[T]{cant: 0, cmp: funcion_cmp}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &colaConPrioridad[T]{datos: arreglo, cant: len(arreglo), cmp: funcion_cmp}
}

func (heap *colaConPrioridad[T]) EstaVacia() bool {
	return heap.cant == 0
}

func (heap *colaConPrioridad[T]) Encolar(elemento T) {
	heap.cant++
	heap.datos[heap.cant-1] = elemento
	heap.heapify(heap.datos, heap.cant-1)
}

func (heap *colaConPrioridad[T]) VerMax() T {
	if heap.EstaVacia() {
		panic("El heap esta vacio")
	}
	return heap.datos[0]
}

func (heap *colaConPrioridad[T]) Desencolar() T {

}

func (heap *colaConPrioridad[T]) Cantidad() int {
	return heap.cant
}

func (heap *colaConPrioridad[T]) heapify(arr []T, i int) {
	if i == 0 {
		return
	}
	nodo_padre := (i - 1) / 2
	comparacion := heap.cmp(arr[i], arr[nodo_padre])
	if comparacion > 0 {
		arr[i], arr[nodo_padre] = arr[nodo_padre], arr[i]
		heap.heapify(arr, nodo_padre)
	}
}

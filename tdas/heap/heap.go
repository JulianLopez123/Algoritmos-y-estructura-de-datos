package cola_prioridad

type colaConPrioridad[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

const PROPORCION_DE_OCUPACION_MINIMA = 4
const CONSTANTE_DE_REDIMENSION = 2
const CAPACIDAD_MINIMA = 5

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &colaConPrioridad[T]{datos: make([]T, CAPACIDAD_MINIMA), cant: 0, cmp: funcion_cmp}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heapify(arreglo, len(arreglo), funcion_cmp)
	cola := &colaConPrioridad[T]{datos: arreglo, cant: len(arreglo), cmp: funcion_cmp}
	return cola
}

func (heap *colaConPrioridad[T]) EstaVacia() bool {
	return heap.cant == 0
}

func (heap *colaConPrioridad[T]) Encolar(elemento T) {
	if len(heap.datos) == heap.cant {
		nuevo_tamaño := cap(heap.datos) * CONSTANTE_DE_REDIMENSION
		heap.redimensionar(nuevo_tamaño)
	}
	heap.cant++
	heap.datos[heap.cant-1] = elemento
	upHeap(heap.datos, heap.cant, heap.cant-1, heap.cmp)
}

func (heap *colaConPrioridad[T]) VerMax() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	return heap.datos[0]
}

func (heap *colaConPrioridad[T]) Desencolar() T {
	dato_eliminado := heap.VerMax()

	if len(heap.datos) >= PROPORCION_DE_OCUPACION_MINIMA*heap.cant {
		capacidad_heap_datos := cap(heap.datos)
		nuevo_tamaño := max(capacidad_heap_datos/CONSTANTE_DE_REDIMENSION, CAPACIDAD_MINIMA)
		heap.redimensionar(nuevo_tamaño)
	}

	heap.datos[0], heap.datos[heap.cant-1] = heap.datos[heap.cant-1], heap.datos[0]
	heap.cant--
	downHeap(heap.datos, heap.cant, 0, heap.cmp)
	return dato_eliminado
}

func (heap *colaConPrioridad[T]) Cantidad() int {
	return heap.cant
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	largo_relativo := len(elementos) - 1
	heapify(elementos, len(elementos), funcion_cmp)
	for largo_relativo != 0 {
		elementos[0], elementos[largo_relativo] = elementos[largo_relativo], elementos[0]
		largo_relativo--
		downHeap(elementos, largo_relativo, 0, funcion_cmp)

	}

}

func heapify[T any](arr []T, largo int, cmp func(T, T) int) {
	if largo == 0 {
		return
	}
	largo_relativo := largo / 2
	ultimo_nodo_con_hijos := largo_relativo - 1
	for i := ultimo_nodo_con_hijos; i > -1; i-- {
		downHeap(arr, largo_relativo, i, cmp)
	}
}

func upHeap[T any](arr []T, largo int, posicion_hijo int, func_cmp func(T, T) int) {
	posicion_padre := (posicion_hijo - 1) / 2
	if posicion_hijo >= largo || posicion_padre < 0 {
		return
	}
	padre := arr[posicion_padre]
	hijo := arr[posicion_hijo]

	if func_cmp(hijo, padre) > 0 {
		arr[posicion_padre], arr[posicion_hijo] = arr[posicion_hijo], arr[posicion_padre]
		upHeap(arr, largo, posicion_padre, func_cmp)
	}
}

func downHeap[T any](arr []T, largo int, posicion_padre int, func_cmp func(T, T) int) {
	pos_hijo_izq := 2*posicion_padre + 1
	pos_hijo_der := 2*posicion_padre + 2
	if pos_hijo_izq >= largo {
		return
	}
	padre := arr[posicion_padre]
	hijo_izq := arr[pos_hijo_izq]

	if pos_hijo_der >= largo {
		if func_cmp(hijo_izq, padre) > 0 {
			arr[posicion_padre], arr[pos_hijo_izq] = arr[pos_hijo_izq], arr[posicion_padre]
		}
		return
	}

	hijo_der := arr[pos_hijo_der]

	if func_cmp(hijo_der, padre) > 0 || func_cmp(hijo_izq, padre) > 0 {
		if func_cmp(hijo_der, hijo_izq) > 0 {
			arr[posicion_padre], arr[pos_hijo_der] = arr[pos_hijo_der], arr[posicion_padre]
			downHeap(arr, largo, pos_hijo_der, func_cmp)
		} else {
			arr[posicion_padre], arr[pos_hijo_izq] = arr[pos_hijo_izq], arr[posicion_padre]
			downHeap(arr, largo, pos_hijo_der, func_cmp)
		}
	}
}

func (heap *colaConPrioridad[T]) redimensionar(nuevo_tamaño int) {
	nuevo_slice := make([]T, nuevo_tamaño)
	copy(nuevo_slice, heap.datos)
	heap.datos = nuevo_slice
}

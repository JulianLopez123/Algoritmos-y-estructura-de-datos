package cola_prioridad

type colaConPrioridad[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

const PROPORCION_DE_OCUPACION_MINIMA = 4
const CONSTANTE_DE_REDIMENSION = 2
const CAPACIDAD_MINIMA = 5
const CORRIMIENTO_HIJO_IZQUIERDO = 1
const CORRIMIENTO_HIJO_DERECHO = 2
const CORRIMIENTO_PADRE = 1
const FACTOR_DE_BUSQUEDA = 2

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &colaConPrioridad[T]{datos: make([]T, CAPACIDAD_MINIMA), cant: 0, cmp: funcion_cmp}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	nuevo_arr := make([]T, max(CAPACIDAD_MINIMA, len(arreglo))*CONSTANTE_DE_REDIMENSION)
	copy(nuevo_arr, arreglo)
	heapify(nuevo_arr, len(arreglo), funcion_cmp)
	cola := &colaConPrioridad[T]{datos: nuevo_arr, cant: len(arreglo), cmp: funcion_cmp}

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
	upHeap(heap.datos, heap.cant-1, heap.cmp)
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
	heapify(elementos, len(elementos), funcion_cmp)
	for i := len(elementos) - 1; i > 0; i-- {
		elementos[0], elementos[i] = elementos[i], elementos[0]
		downHeap(elementos, i, 0, funcion_cmp)
	}
}

func heapify[T any](arr []T, largo int, cmp func(T, T) int) {
	if largo == 0 {
		return
	}
	ultimo_nodo_con_hijos := (largo / FACTOR_DE_BUSQUEDA) - CORRIMIENTO_PADRE
	for i := ultimo_nodo_con_hijos; i > -1; i-- {
		downHeap(arr, largo, i, cmp)
	}
}

func upHeap[T any](arr []T, posicion_hijo int, func_cmp func(T, T) int) {
	if posicion_hijo == 0 {
		return
	}
	posicion_padre := (posicion_hijo - CORRIMIENTO_PADRE) / FACTOR_DE_BUSQUEDA
	if func_cmp(arr[posicion_hijo], arr[posicion_padre]) > 0 {
		arr[posicion_padre], arr[posicion_hijo] = arr[posicion_hijo], arr[posicion_padre]
		upHeap(arr, posicion_padre, func_cmp)
	}
}

func downHeap[T any](arr []T, largo int, posicion_padre int, func_cmp func(T, T) int) {
	pos_hijo_izq := FACTOR_DE_BUSQUEDA*posicion_padre + CORRIMIENTO_HIJO_IZQUIERDO
	pos_hijo_der := FACTOR_DE_BUSQUEDA*posicion_padre + CORRIMIENTO_HIJO_DERECHO
	pos_elemento_mayor := posicion_padre

	if pos_hijo_izq < largo && func_cmp(arr[pos_hijo_izq], arr[pos_elemento_mayor]) > 0 {
		pos_elemento_mayor = pos_hijo_izq
	}
	if pos_hijo_der < largo && func_cmp(arr[pos_hijo_der], arr[pos_elemento_mayor]) > 0 {
		pos_elemento_mayor = pos_hijo_der
	}
	if pos_elemento_mayor != posicion_padre {
		arr[posicion_padre], arr[pos_elemento_mayor] = arr[pos_elemento_mayor], arr[posicion_padre]
		downHeap(arr, largo, pos_elemento_mayor, func_cmp)
	}
}

func (heap *colaConPrioridad[T]) redimensionar(nuevo_tamaño int) {
	nuevo_slice := make([]T, nuevo_tamaño)
	copy(nuevo_slice, heap.datos)
	heap.datos = nuevo_slice
}

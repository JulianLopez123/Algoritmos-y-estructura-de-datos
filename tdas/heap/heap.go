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
	cola := &colaConPrioridad[T]{datos: arreglo, cant: len(arreglo), cmp: funcion_cmp}
	cola.heapify(cola.datos,cola.cant,cola.cmp)
	return cola
}

func (heap *colaConPrioridad[T]) EstaVacia() bool {
	return heap.cant == 0
}

func (heap *colaConPrioridad[T]) Encolar(elemento T) {
	heap.cant++
	heap.datos[heap.cant-1] = elemento
	heap.heapify(heap.datos, heap.cant-1,heap.cmp)
}

func (heap *colaConPrioridad[T]) VerMax() T {
	return heap.datos[0]
}

func (heap *colaConPrioridad[T]) Desencolar() T {
	if heap.EstaVacia(){
		panic("La cola esta vacia")
	}
	dato_eliminado:= heap.datos[0]
	heap.datos[0], heap.datos[heap.cant-1] = heap.datos[heap.cant-1], heap.datos[0]
	heap.cant--
	heap.downHeap(0)
	return dato_eliminado
}

func (heap *colaConPrioridad[T]) Cantidad() int {
	return heap.cant
}

func (heap *colaConPrioridad[T]) heapify(arr []T, largo int, cmp func(T, T) int) {
	if largo == 0{
		return 
	}
	heap.datos = arr
	heap.cant = largo
	heap.cmp = cmp
	
	ultimo_nodo_con_hijos := (largo /2) -1
	for i := ultimo_nodo_con_hijos; i > -1;i--{
		heap.downHeap(i)
	}
}

func (heap *colaConPrioridad[T]) downHeap(posicion_padre int){
	pos_hijo_izq := 2 * posicion_padre + 1
	pos_hijo_der := 2 * posicion_padre + 2
	if pos_hijo_izq >= heap.cant{
		return
	}
	padre := heap.datos[posicion_padre]
	hijo_izq := heap.datos[pos_hijo_izq]
	
	if pos_hijo_der >= heap.cant{
		if heap.cmp(hijo_izq,padre) > 0{
			heap.datos[posicion_padre], heap.datos[pos_hijo_izq] = heap.datos[pos_hijo_izq], heap.datos[posicion_padre]
		}
		return
	}

	hijo_der := heap.datos[pos_hijo_der]

	if heap.cmp(hijo_der,padre) > 0 ||  heap.cmp(hijo_izq,padre) > 0{
		if heap.cmp(hijo_der,hijo_izq) > 0{
			heap.datos[posicion_padre], heap.datos[pos_hijo_der] = heap.datos[pos_hijo_der], heap.datos[posicion_padre]
			heap.downHeap(pos_hijo_der)
		}else{
			heap.datos[posicion_padre], heap.datos[pos_hijo_izq] = heap.datos[pos_hijo_izq], heap.datos[posicion_padre]
			heap.downHeap(pos_hijo_izq)
		}
	}
}
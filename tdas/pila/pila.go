package pila

type Pila[T any] interface {

	// EstaVacia devuelve verdadero si la pila no tiene elementos apilados, false en caso contrario.
	EstaVacia() bool

	// VerTope obtiene el valor del tope de la pila. Si la pila tiene elementos se devuelve el valor del tope.
	// Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
	VerTope() T

	// Apilar agrega un nuevo elemento a la pila.
	Apilar(T)

	// Desapilar saca el elemento tope de la pila. Si la pila tiene elementos, se quita el tope de la pila, y
	// se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
	Desapilar() T
}

func esta_ord(arr []int) bool {
	return esta_ord_rec(arr, 0, len(arr)-1)
}

func esta_ord_rec(arr []int, inicio, fin int) bool {
	if inicio >= fin {
		return true
	}
	mitad := (inicio + fin) / 2

	if arr[inicio] > arr[mitad] || arr[fin] < arr[mitad] {
		return false
	}
	return esta_ord_rec(arr, mitad+1, fin) && esta_ord_rec(arr, inicio, mitad-1)
}

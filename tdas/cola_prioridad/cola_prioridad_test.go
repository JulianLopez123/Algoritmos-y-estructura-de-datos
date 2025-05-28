package cola_prioridad_test

import (
	"fmt"
	TDAHeap "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

var TAMS_VOLUMEN = []int{12500, 25000, 50000, 100000, 200000, 400000}

func TestHeapVacio(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](func(a, b int) int { return a - b })

	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "El heap esta vacio", func() { heap.VerMax() })
	require.PanicsWithValue(t, "El heap esta vacio", func() { heap.Desencolar() })
}

func TestUnElemento(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](func(a, b int) int { return a - b })

	heap.Encolar(1)
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, 1, heap.VerMax())
	require.Equal(t, 1, heap.Desencolar())
	require.PanicsWithValue(t, "El heap esta vacio", func() { heap.VerMax() })
}

func TestHeapGuardar(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](func(a, b int) int { return a - b })

	require.PanicsWithValue(t, "El heap esta vacio", func() { heap.VerMax() })
	heap.Encolar(1)
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, 1, heap.VerMax())

	heap.Encolar(2)
	require.Equal(t, 2, heap.Cantidad())
	require.Equal(t, 2, heap.VerMax())

	heap.Encolar(4)
	require.Equal(t, 3, heap.Cantidad())
	require.Equal(t, 4, heap.VerMax())

	heap.Encolar(3)
	require.Equal(t, 4, heap.Cantidad())
	require.Equal(t, 4, heap.VerMax())
}

func TestEncoloMuchoElementos(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](func(a, b int) int { return a - b })

	for i := 0; i < 200; i++ {
		heap.Encolar(i)
		require.Equal(t, i, heap.VerMax())
	}

	for i := 0; i < 200; i++ {
		heap.Desencolar()
	}

	require.True(t, heap.EstaVacia())
}

func TestDiferenteTipoDeDato(t *testing.T) {
	heap := TDAHeap.CrearHeap[string](func(a, b string) int {
		if a == b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	})

	heap.Encolar("hola")
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, "hola", heap.VerMax())
	heap.Desencolar()
	require.PanicsWithValue(t, "El heap esta vacio", func() { heap.VerMax() })
}

func TestGuardarYBorrarRepetidasVeces(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](func(a, b int) int { return a - b })

	for i := 0; i < 1000; i++ {
		heap.Encolar(i)
		require.Equal(t, i, heap.VerMax())
		heap.Desencolar()
		require.Equal(t, 0, heap.Cantidad())
	}
}

func TestDesencolarEnOrden(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](func(a, b int) int { return a - b })
	elementos := []int{5, 8, 9, 6, 3, 2}

	for _, element := range elementos {
		heap.Encolar(element)
	}

	esperado := []int{9, 8, 6, 5, 3, 2}

	for i, _ := range elementos {
		require.Equal(t, esperado[i], heap.Desencolar())
	}
}

func TestEncoloYDesencoloElementosIguales(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](func(a, b int) int { return a - b })

	for i := 0; i < 10; i++ {
		heap.Encolar(2)
	}

	for i := 0; i < 10; i++ {
		require.Equal(t, 2, heap.Desencolar())
	}

	require.True(t, heap.EstaVacia())
}

func TestHeapifyArregloDesordenado(t *testing.T) {
	arr := []int{5, 8, 9, 6, 3, 2}
	heap := TDAHeap.CrearHeapArr[int](arr, func(a, b int) int { return a - b })

	require.Equal(t, len(arr), heap.Cantidad())

	arr_ordenado := []int{9, 8, 6, 5, 3, 2}

	for i := 0; i < len(arr); i++ {
		require.Equal(t, arr_ordenado[i], heap.Desencolar())
	}
}

func TestHeapifyArregloInvertido(t *testing.T) {
	arr := []int{2, 3, 5, 6, 8, 9}
	heap := TDAHeap.CrearHeapArr[int](arr, func(a, b int) int { return a - b })

	require.Equal(t, len(arr), heap.Cantidad())

	arr_ordenado := []int{9, 8, 6, 5, 3, 2}

	for i := 0; i < len(arr); i++ {
		require.Equal(t, arr_ordenado[i], heap.Desencolar())
	}
}

func TestDosHeapsMismoArr(t *testing.T) {
	arr := []int{4, 2, 4, 6, 1, 2, 8}

	heap_arr := TDAHeap.CrearHeapArr[int](arr, func(a, b int) int { return a - b })
	heap := TDAHeap.CrearHeap(func(a, b int) int { return a - b })

	for i := 0; i < len(arr); i++ {
		heap.Encolar(arr[i])
	}

	for i := 0; i < len(arr); i++ {
		require.Equal(t, heap.Desencolar(), heap_arr.Desencolar())
	}

}

func TestHeapifyArregloOrdenado(t *testing.T) {
	arr := []int{9, 8, 6, 5, 3, 2}
	heap := TDAHeap.CrearHeapArr[int](arr, func(a, b int) int { return a - b })

	require.Equal(t, len(arr), heap.Cantidad())

	arr_ordenado := []int{9, 8, 6, 5, 3, 2}

	for i := 0; i < len(arr); i++ {
		require.Equal(t, arr_ordenado[i], heap.Desencolar())
	}
}

func ejecutarPruebaVolumen(b *testing.B, n int) {
	heap := TDAHeap.CrearHeap[int](func(a, b int) int { return a - b })

	for i := 0; i < n; i++ {
		heap.Encolar(i)
	}

	require.Equal(b, n, heap.Cantidad())
	require.Equal(b, n-1, heap.VerMax())

	for i := 0; i < n; i++ {
		require.Equal(b, (n-1)-i, heap.Desencolar())
	}

	require.Equal(b, 0, heap.Cantidad())
}

func BenchmarkHeap(b *testing.B) {
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumen(b, n)
			}
		})
	}
}

func TestRedimension(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](func(a, b int) int { return a - b })

	for i := 0; i < 100; i++ {
		heap.Encolar(i)
	}

	require.Equal(t, 100, heap.Cantidad())
	require.Equal(t, 99, heap.VerMax())

	for i := 99; i >= 0; i-- {
		require.Equal(t, i, heap.Desencolar())
	}

	for i := 0; i < 200; i++ {
		heap.Encolar(i)
	}

	require.Equal(t, 200, heap.Cantidad())
	require.Equal(t, 199, heap.VerMax())

	for i := 199; i >= 0; i-- {
		require.Equal(t, i, heap.Desencolar())
	}

	require.True(t, heap.EstaVacia())
}

func TestHeapSort(t *testing.T) {
	arr := []int{3, 1, 4, 1, 5, 9, 2}
	TDAHeap.HeapSort(arr, func(a, b int) int { return a - b })

	esperado := []int{1, 1, 2, 3, 4, 5, 9}

	require.Equal(t, esperado, arr)
}

func TestHeapSortStrings(t *testing.T) {
	arr := []string{"hola", "si", "como estas?"}
	TDAHeap.HeapSort(arr, func(a, b string) int { return len(a) - len(b) })

	esperado := []string{"si", "hola", "como estas?"}

	require.Equal(t, esperado, arr)
}

func TestHeapSortArrVacio(t *testing.T) {
	arr := []int{}
	TDAHeap.HeapSort(arr, func(a, b int) int { return a - b })

	esperado := []int{}

	require.Equal(t, arr, esperado)
}

func TestDesencolarYHeapsortDevuelvenArrConOrdenOpuesto(t *testing.T) {
	arr := []int{3, 1, 4, 1, 5, 9, 2}
	heap := TDAHeap.CrearHeapArr(arr, func(a, b int) int { return a - b })

	var arr_esperado []int
	for i := 0; i < len(arr); i++ {
		arr_esperado = append(arr_esperado, heap.Desencolar())
	}
	TDAHeap.HeapSort(arr, func(a, b int) int { return b - a })
	require.Equal(t, arr_esperado, arr)

}

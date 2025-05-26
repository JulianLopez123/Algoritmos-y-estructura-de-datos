package cola_prioridad_test

import (
	"fmt"
	TDAHeap "tdas/heap"
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

func TestHeapify(t *testing.T) {
	arr := []int{5, 8, 9, 6, 3, 2}
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

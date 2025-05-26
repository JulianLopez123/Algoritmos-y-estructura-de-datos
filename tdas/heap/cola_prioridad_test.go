package cola_prioridad_test

import (
	TDAHeap "tdas/heap"
	"testing"

	"github.com/stretchr/testify/require"
)

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

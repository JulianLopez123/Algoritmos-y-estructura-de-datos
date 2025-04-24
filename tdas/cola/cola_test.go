package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.True(t, cola.EstaVacia())
	cola.Encolar(2)
	require.False(t, cola.EstaVacia())
	cola.Desencolar()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
}

func TestFIFO(t *testing.T) {
	cola_string := TDACola.CrearColaEnlazada[string]()
	cola_string.Encolar("primero")
	cola_string.Encolar("segundo")
	require.Equal(t, "primero", cola_string.Desencolar())
}

func TestVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for k := 0; k < 10000; k++ {
		cola.Encolar(k)
	}
	for i := 0; i < 10000; i++ {
		require.Equal(t, i, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
}

func TestVerPrimeroDespuesDeDesencolar(t *testing.T) {
	cola_float := TDACola.CrearColaEnlazada[float64]()
	cola_float.Encolar(7.102122)
	cola_float.Encolar(3.142332)
	cola_float.Encolar(1.233212)
	require.Equal(t, 7.102122, cola_float.VerPrimero())
	require.Equal(t, 7.102122, cola_float.Desencolar())
	require.Equal(t, 3.142332, cola_float.VerPrimero())
}

package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}

func TestLIFO(t *testing.T) {
	pila_string := TDAPila.CrearPilaDinamica[string]()
	pila_string.Apilar("hola")
	pila_string.Apilar("como")
	pila_string.Apilar("va")
	require.Equal(t, "va", pila_string.Desapilar())
	require.Equal(t, "como", pila_string.Desapilar())
	require.Equal(t, "hola", pila_string.Desapilar())

	pila_float64 := TDAPila.CrearPilaDinamica[float64]()
	pila_float64.Apilar(7.102122)
	pila_float64.Apilar(3.142332)
	require.Equal(t, 3.142332, pila_float64.Desapilar())
	require.Equal(t, 7.102122, pila_float64.Desapilar())

	pila_int := TDAPila.CrearPilaDinamica[int]()
	for l := 0; l < 5; l++ {
		pila_int.Apilar(l)
	}
	for l := 4; l > -1; l-- {
		require.Equal(t, l, pila_int.Desapilar())
	}
}

func TestVolumen(t *testing.T) {
	//valido los topes anterior y posterior a cada apilar y desapilar.
	pila := TDAPila.CrearPilaDinamica[int]()

	for k := 0; k < 10000; k++ {
		if k != 0 {
			require.Equal(t, k-1, pila.VerTope())
		}
		pila.Apilar(k)
		require.Equal(t, k, pila.VerTope())
	}
	for i := 9999; i > -1; i-- {
		require.Equal(t, i, pila.VerTope())
		require.Equal(t, i, pila.Desapilar())
		if i != 0 {
			require.Equal(t, i-1, pila.VerTope())
		}
	}
	require.True(t, pila.EstaVacia())
}

func TestVaciarPila(t *testing.T) {
	//A mi parecer este test es necesario debido que al achicar TestPilaVacia,
	//debo probar si una pila está vacía luego de apilar y de desapilar
	//en un test aparte,y verificar que esta se comporte como tal
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(2)
	require.False(t, pila.EstaVacia())
	pila.Desapilar()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}

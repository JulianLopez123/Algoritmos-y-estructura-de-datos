package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
}

func TestVerPrimeroListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "Se espera un panic cuando se intenta ver el primero de una lista vacia")
}

func TestVerUltimoListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() }, "Se espera un panic cuando se intenta ver el ultimo de una lista vacia")
}

func Test(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)

	primero := lista.BorrarPrimero()
	require.Equal(t, 3, primero)

	lista.InsertarPrimero(4)

	primero = lista.BorrarPrimero()
	require.Equal(t, 4, primero)
	primero = lista.BorrarPrimero()
	require.Equal(t, 2, primero)
	primero = lista.BorrarPrimero()
	require.Equal(t, 1, primero)

	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() }, "Se espera un panic cuando se intenta borrar un elemento de una lista vacia")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "Se espera un panic cuando se intenta ver el primero de una lista vacia")
}

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

func TestInsertarPrimeroYBorrarPrimero(t *testing.T) {
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

func TestVerificarInsercionAlCrearIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(1)

	primero := iter.VerActual()
	require.Equal(t, 1, primero)
	require.Equal(t, 1, lista.VerPrimero())
}

func TestBorrarElementoIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()

	iter.Insertar(1)
	iter.Insertar(2)

	require.Equal(t, 2, lista.VerPrimero())

	iter.Borrar()

	require.Equal(t, 1, lista.VerPrimero())
}

func TestBorrarUltimoElementoIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	require.Equal(t, 3, lista.VerUltimo())

	iter := lista.Iterador()
	for iter.HaySiguiente() {
		if iter.VerActual() == 3 {
			iter.Borrar()
			break
		}
		iter.Siguiente()
	}

	require.Equal(t, 2, lista.VerUltimo())
	require.Equal(t, 2, lista.Largo())
}

func TestBorrarElementoDelMedioIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iter := lista.Iterador()
	for iter.HaySiguiente() {
		if iter.VerActual() == 2 {
			iter.Borrar()
			break
		}
		iter.Siguiente()
	}

	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())
}

func TestIteradorExterno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	iterador := lista.Iterador()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
	iterador.Insertar("todo bien")
	iterador.Insertar("hola")
	require.Equal(t, "hola", iterador.VerActual())
	require.Equal(t, "hola", lista.VerPrimero())
	iterador.Siguiente()
	require.Equal(t, "todo bien", iterador.VerActual())
}

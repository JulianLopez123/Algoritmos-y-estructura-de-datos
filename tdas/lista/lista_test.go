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

func TestVerPrimeroYVerUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarUltimo(1)
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())

	lista.InsertarUltimo(2)
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())

	lista.InsertarPrimero(3)
	require.Equal(t, 3, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())
}

func TestInsertarUltimoYBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	require.Equal(t, 1, lista.BorrarPrimero())
	require.Equal(t, 2, lista.BorrarPrimero())
	require.Equal(t, 3, lista.BorrarPrimero())

	require.True(t, lista.EstaVacia())
}

func TestIterar(t *testing.T){
	lista := TDALista.CrearListaEnlazada[float64]()
	lista.InsertarUltimo(7.102122)
	lista.InsertarUltimo(3.142332)
	lista.InsertarUltimo(1.233212)
	contador := 0
	lista.Iterar(func (elemento float64) bool {if elemento * 2 > 0{contador ++;return true}else{return false}})
	require.Equal(t,3,contador)
	contador = 0
	lista.Iterar(func (elemento float64) bool {if elemento > 5{contador ++;return true}else{return false}})
	require.Equal(t,1,contador)
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

func TestInsertarIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	iterador := lista.Iterador()
	iterador.Insertar("todo bien")
	require.Equal(t, "todo bien", iterador.VerActual())
	require.Equal(t, "todo bien", lista.VerPrimero())
	iterador.Insertar("hola")
	require.Equal(t, "hola", iterador.VerActual())
	require.Equal(t, "todo bien", lista.VerUltimo())
	require.Equal(t, "hola", lista.VerPrimero())
	iterador.Siguiente()
	require.True(t, iterador.HaySiguiente())
	iterador.Siguiente()
	iterador.Insertar("que tal")
	require.Equal(t, "que tal", lista.VerUltimo())
	iterador.Siguiente()
	require.False(t, iterador.HaySiguiente())
	require.Equal(t, 3, lista.Largo())
}

func TestListaVaciaIterador(t *testing.T){
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Borrar() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Siguiente() })
	require.False(t, iterador.HaySiguiente())
}


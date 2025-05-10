package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
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

	lista.InsertarPrimero(5)
	lista.InsertarPrimero(6)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 6, lista.VerPrimero())
	require.Equal(t, 5, lista.VerUltimo())
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

func TestInsertarTiposDiferentes(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()

	lista.InsertarUltimo("hola")
	lista.InsertarUltimo("todo bien")

	require.False(t, lista.EstaVacia())

	valor_borrado := lista.BorrarPrimero()
	require.Equal(t, "hola", valor_borrado)
	require.Equal(t, 1, lista.Largo())
	valor_borrado = lista.BorrarPrimero()
	require.Equal(t, "todo bien", valor_borrado)

	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() }, "Se espera un panic cuando se intenta borrar un elemento de una lista vacia")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "Se espera un panic cuando se intenta ver el primero de una lista vacia")
}

func TestEstadoConSlices(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[[]string]()
	lista.InsertarPrimero([]string{"a", "b"})
	lista.InsertarUltimo([]string{"c", "d"})
	lista.InsertarPrimero([]string{"e", "f"})
	require.Equal(t, []string{"e", "f"}, lista.BorrarPrimero())
	lista.InsertarUltimo([]string{"g", "h"})
	require.Equal(t, []string{"a", "b"}, lista.BorrarPrimero())

	require.Equal(t, 2, lista.Largo())
	require.Equal(t, []string{"c", "d"}, lista.VerPrimero())
	require.Equal(t, []string{"g", "h"}, lista.VerUltimo())
}

func TestElementoNil(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[*int]()
	var n *int = nil

	lista.InsertarPrimero(n)
	require.False(t, lista.EstaVacia())
	require.Equal(t, n, lista.VerPrimero())
}

func TestVolumenInsertarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	cant := 10000
	for i := 0; i < cant; i++ {
		lista.InsertarUltimo(i)
		ultimo := lista.VerUltimo()
		require.Equal(t, i, ultimo)
	}

	for i := 0; i < cant; i++ {
		primero := lista.VerPrimero()
		require.Equal(t, i, primero)
		valor_borrado := lista.BorrarPrimero()
		require.Equal(t, i, valor_borrado)
	}

	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() }, "Se espera un panic cuando se intenta borrar un elemento de una lista vacia")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "Se espera un panic cuando se intenta ver el primero de una lista vacia")
}

func TestVolumenInsertarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	cant := 10000
	for i := 0; i < cant; i++ {
		lista.InsertarPrimero(i)
		primero := lista.VerPrimero()
		require.Equal(t, i, primero)
	}

	for i := 0; i < cant; i++ {
		ultimo := lista.VerUltimo()
		require.Equal(t, 0, ultimo)
		valor_borrado := lista.BorrarPrimero()
		require.Equal(t, (cant-1)-i, valor_borrado)
	}

	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() }, "Se espera un panic cuando se intenta borrar un elemento de una lista vacia")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "Se espera un panic cuando se intenta ver el primero de una lista vacia")
}

func TestIterarConCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[float64]()
	lista.InsertarUltimo(7.102122)
	lista.InsertarUltimo(3.142332)
	lista.InsertarUltimo(1.233212)
	contador := 0
	lista.Iterar(func(elemento float64) bool {
		if elemento > 5 {
			contador++
			return true
		} else {
			return false
		}
	})
	require.Equal(t, 1, contador)

	contador = 0
	lista.Iterar(func(elemento float64) bool {
		return false
	})
	require.Equal(t, 0, contador)
}
func TestIterarSinCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(7)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(20)
	contador := 0
	lista.Iterar(func(elemento int) bool {
		if elemento*2 > 0 {
			contador++
			return true
		} else {
			return false //no se ejecuta nunca en este caso
		}
	})
	require.Equal(t, 3, contador)

	contador = 0
	lista.Iterar(func(elemento int) bool {
		contador++
		return true
	})
	require.Equal(t, 3, contador)
}

func TestIterarListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	contador := 0
	lista.Iterar(func(elemento int) bool {
		contador++
		return true
	})
	require.Equal(t, 0, contador)
}

func TestIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 50; i++ {
		lista.InsertarUltimo(i)
	}
	iterador := lista.Iterador()
	numero := 0
	for iterador.HaySiguiente() {
		require.Equal(t, numero, iterador.VerActual())
		numero++
		iterador.Siguiente()
	}

	require.False(t, iterador.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Borrar() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Siguiente() })
}

func TestIteradorListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	require.False(t, iterador.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Borrar() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Siguiente() })
}

func TestIteradorVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 100000; i++ {
		lista.InsertarUltimo(i)
	}
	iter := lista.Iterador()
	contador := 0
	for iter.HaySiguiente() {
		require.Equal(t, contador, iter.VerActual())
		iter.Siguiente()
		contador++
	}
	require.Equal(t, 100000, contador)
}

func TestVerificarInsercionAlCrearIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(1)
	iter.Insertar(2)
	require.Equal(t, 2, lista.VerPrimero())
	iter.Borrar()
	primero := iter.VerActual()
	require.Equal(t, 1, primero)
	require.Equal(t, 1, lista.VerPrimero())
}

func TestBorrarElementoIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(11)
	lista.InsertarUltimo(32)
	iter := lista.Iterador()
	iter.Siguiente()
	require.Equal(t, 32, iter.VerActual())
	iter.Insertar(54)
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 54, iter.VerActual())
	require.Equal(t, 32, lista.VerUltimo())

	iter.Borrar()
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 32, lista.VerUltimo())
	require.Equal(t, 32, iter.VerActual())

	iter.Siguiente()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })

	iter = lista.Iterador()
	require.Equal(t, 11, iter.VerActual())

	iter.Borrar()
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 32, lista.VerPrimero())
	require.Equal(t, 32, lista.VerUltimo())
	require.Equal(t, 32, iter.VerActual())

	iter.Borrar()
	require.Equal(t, 0, lista.Largo())
	require.True(t, lista.EstaVacia())
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

func TestDosIteradores(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(5)
	lista.InsertarUltimo(23)
	lista.InsertarUltimo(32)
	iterador_1 := lista.Iterador()
	lista.InsertarPrimero(100)
	iterador_2 := lista.Iterador()

	require.Equal(t, 5, iterador_1.VerActual())
	require.Equal(t, 100, iterador_2.VerActual())
	require.Equal(t, 100, lista.VerPrimero())

	iterador_2.Siguiente()
	iterador_1.Siguiente()
	require.Equal(t, 23, iterador_1.VerActual())
	require.Equal(t, 5, iterador_2.VerActual())

	iterador_1.Insertar(15)
	require.Equal(t, 15, iterador_1.VerActual())

	iterador_1.Siguiente()
	iterador_2.Siguiente()
	require.Equal(t, 15, iterador_2.VerActual()) //la insercion en el iterador_1 se vio reflejada en el iterador_2

	iterador_1.Borrar()
	iterador_2.Siguiente()
	require.Equal(t, 32, iterador_1.VerActual()) //el iterador_1 al borrar apunto al siguiente elemento al borrado(ahora actual)
	require.Equal(t, 32, iterador_2.VerActual()) //al borrarse su proximo elemento, al moverse lo hizo hacia su nuevo proximo(el que le seguia al eliminado)

	iterador_1.Siguiente()
	iterador_1.Insertar(999)
	require.Equal(t, 999, lista.VerUltimo()) //se actualizan los punteros de la lista de primero y ultimo con las primitivas del iterador

	lista.InsertarUltimo(100) //aunque no es parte del contrato, la lista si se ve modificada con las primitivas de la lista, y esos cambios reflejados en los iters
	iterador_1.Siguiente()
	require.Equal(t, 100, iterador_1.VerActual())
	require.Equal(t, 100, lista.VerUltimo())
}

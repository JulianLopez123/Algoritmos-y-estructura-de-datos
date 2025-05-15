package diccionario_test

import (
	"fmt"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

var TAMS_VOLUMEN_ABB = []int{1500, 2000, 5000, 10000, 20000} //menores que en hash(al ser fors incrementales el abb se vuelve lineal en tiempo al recorrerlo)

func TestABBDiccionarioVacio(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, string](func(a string, b string) int {
		if a == b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	})
	require.Equal(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}

func TestABBDiccionarioClaveDefault(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, string](func(a string, b string) int {
		if a == b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	})
	require.False(t, dic.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("") })

	dicNum := TDADiccionario.CrearABB[int, string](func(a int, b int) int { return a - b })
	require.False(t, dicNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Borrar(0) })
}

func TestABBUnElement(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](func(a string, b string) int {
		if a == b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	})
	dic.Guardar("A", 10)
	require.Equal(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.Equal(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestABBDiccionarioGuardar(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, string](func(a string, b string) int {
		if a == b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	})
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestABBReemplazoDato(t *testing.T) {
	clave := "Gato"
	clave2 := "Perro"
	dic := TDADiccionario.CrearABB[string, string](func(a string, b string) int {
		if a == b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	})
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.Equal(t, "miau", dic.Obtener(clave))
	require.Equal(t, "guau", dic.Obtener(clave2))
	require.Equal(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.Equal(t, 2, dic.Cantidad())
	require.Equal(t, "miu", dic.Obtener(clave))
	require.Equal(t, "baubau", dic.Obtener(clave2))
}

func TestABBReemplazoDatoHopscotch(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a int, b int) int { return a - b })
	for i := 0; i < 500; i++ {
		dic.Guardar(i, i)
	}
	for i := 0; i < 500; i++ {
		dic.Guardar(i, 2*i)
	}
	ok := true
	for i := 0; i < 500 && ok; i++ {
		ok = dic.Obtener(i) == 2*i
	}
	require.True(t, ok, "Los elementos no fueron actualizados correctamente")
}

func TestABBDiccionarioBorrar(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, string](func(a string, b string) int {
		if a == b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	})
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[0]) })

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[1]) })
}

func TestABBReutlizacionDeBorrados(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, string](func(a string, b string) int {
		if a == b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	})
	clave := "hola"
	dic.Guardar(clave, "mundo!")
	dic.Borrar(clave)
	require.Equal(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(clave))
	dic.Guardar(clave, "mundooo!")
	require.True(t, dic.Pertenece(clave))
	require.Equal(t, 1, dic.Cantidad())
	require.Equal(t, "mundooo!", dic.Obtener(clave))
}

func TestABBConClavesNumericas(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](func(a int, b int) int { return a - b })
	clave := 10
	valor := "Gatito"

	dic.Guardar(clave, valor)
	require.Equal(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave))
	require.Equal(t, valor, dic.Obtener(clave))
	require.Equal(t, valor, dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestABBClaveVacia(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, string](func(a string, b string) int {
		if a == b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	})
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.Equal(t, 1, dic.Cantidad())
	require.Equal(t, clave, dic.Obtener(clave))
}

func TestABBValorNulo(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, *int](func(a string, b string) int {
		if a == b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	})
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.Equal(t, 1, dic.Cantidad())
	require.Equal(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestABBCadenaLargaParticular(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, string](func(a string, b string) int {
		if a == b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	})
	claves := make([]string, 10)
	cadena := "%d~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~" +
		"~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"
	valores := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	for i := 0; i < 10; i++ {
		claves[i] = fmt.Sprintf(cadena, i)
		dic.Guardar(claves[i], valores[i])
	}
	require.EqualValues(t, 10, dic.Cantidad())

	ok := true
	for i := 0; i < 10 && ok; i++ {
		ok = dic.Obtener(claves[i]) == valores[i]
	}

	require.True(t, ok, "Obtener clave larga funciona")
}

func TestABBGuardarYBorrarRepetidasVeces(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a int, b int) int { return a - b })
	for i := 0; i < 1000; i++ {
		dic.Guardar(i, i)
		require.True(t, dic.Pertenece(i))
		dic.Borrar(i)
		require.False(t, dic.Pertenece(i))
	}
}

func TestABBGuardarYBorrarRepetidasVecesComparacionDistinta(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a int, b int) int { return 40*b - 40*a })
	for i := 0; i < 1000; i++ {
		dic.Guardar(i, i)
		require.True(t, dic.Pertenece(i))
		dic.Borrar(i)
		require.False(t, dic.Pertenece(i))
	}
}

func buscarABB(clave string, claves []string) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}

func TestABBIteradorInternoClaves(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, *int](func(a string, b string) int {
		if a == b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	})
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	dic.Guardar(clave1, nil)
	dic.Guardar(clave2, nil)
	dic.Guardar(clave3, nil)

	cs := []string{"", "", ""}
	cantidad := 0

	dic.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		cantidad++
		return true
	})

	require.EqualValues(t, 3, cantidad)
	require.Equal(t, cs[0], clave1)
	require.Equal(t, cs[1], clave2)
	require.Equal(t, cs[2], clave3)
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

func TestABBIteradorInternoValores(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	dic := TDADiccionario.CrearABB[string, int](func(a string, b string) int {
		if a == b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	})
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	multi := 1
	dic.Iterar(func(_ string, dato int) bool {
		multi *= dato
		return true
	})

	require.EqualValues(t, 720, multi)
}

func TestABBIteradorInternoValoresConBorrados(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno, sin recorrer datos borrados")
	dic := TDADiccionario.CrearABB[string, int](func(a string, b string) int {
		if a == b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	})
	clave0 := "Elefante"
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic.Guardar(clave0, 7)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	dic.Borrar(clave0)

	multi := 1
	dic.Iterar(func(_ string, dato int) bool {
		multi *= dato
		return true
	})

	require.EqualValues(t, 720, multi)
}

func ejecutarPruebaVolumenABB(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[string, int](func(a string, b string) int {
		if a == b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	})

	claves := make([]string, n)
	valores := make([]int, n)

	for i := 0; i < n; i++ {
		valores[i] = i
		claves[i] = fmt.Sprintf("%08d", i)
		dic.Guardar(claves[i], valores[i])
	}

	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	ok := true
	for i := 0; i < n; i++ {
		ok = dic.Pertenece(claves[i])
		if !ok {
			break
		}
		ok = dic.Obtener(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	for i := 0; i < n; i++ {
		ok = dic.Borrar(claves[i]) == valores[i]
		if !ok {
			break
		}
		ok = !dic.Pertenece(claves[i])
		if !ok {
			break
		}
	}

	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, dic.Cantidad())
}

func BenchmarkDiccionarioABB(b *testing.B) {
	b.Log("Prueba de stress del Diccionario. Prueba guardando distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que podemos obtener y ver si pertenece cada una de las claves geeneradas, " +
		"y que luego podemos borrar sin problemas")
	for _, n := range TAMS_VOLUMEN_ABB {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenABB(b, n)
			}
		})
	}
}

func TestABBIterarDiccionarioVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccionario.CrearABB[string, int](func(a string, b string) int {
		if a == b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	})
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestABBDiccionarioIterar(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](func(a string, b string) int {
		if a == b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	})
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	iter := dic.Iterador()

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscarABB(primero, claves))

	iter.Siguiente()
	segundo, segundo_valor := iter.VerActual()
	require.NotEqualValues(t, -1, buscarABB(segundo, claves))
	require.EqualValues(t, valores[buscarABB(segundo, claves)], segundo_valor)
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscarABB(tercero, claves))
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestABBIteradorNoLlegaAlFinal(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccionario.CrearABB[string, string](func(a string, b string) int {
		if a == b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	})
	claves := []string{"A", "B", "C"}
	dic.Guardar(claves[0], "")
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[2], "")

	dic.Iterador()
	iter2 := dic.Iterador()
	iter2.Siguiente()
	iter3 := dic.Iterador()
	primero, _ := iter3.VerActual()
	iter3.Siguiente()
	segundo, _ := iter3.VerActual()
	iter3.Siguiente()
	tercero, _ := iter3.VerActual()
	iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, -1, buscarABB(primero, claves))
	require.NotEqualValues(t, -1, buscarABB(segundo, claves))
	require.NotEqualValues(t, -1, buscarABB(tercero, claves))
}
func TestABBIteradorDiccionarioBalanceado(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccionario.CrearABB[string, string](func(a string, b string) int {
		if a == b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	})
	claves := []string{"B", "A", "D", "C", "Z", "F"}
	for _, i := range claves {
		dic.Guardar(i, "")
	}

	iter := dic.Iterador()
	primero, _ := iter.VerActual()
	iter.Siguiente()
	segundo, _ := iter.VerActual()
	iter.Siguiente()
	tercero, _ := iter.VerActual()
	iter.Siguiente()
	cuarto, _ := iter.VerActual()
	iter.Siguiente()
	quinto, _ := iter.VerActual()
	iter.Siguiente()
	sexto, _ := iter.VerActual()
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
	require.EqualValues(t, primero, claves[1])
	require.EqualValues(t, segundo, claves[0])
	require.EqualValues(t, tercero, claves[3])
	require.EqualValues(t, cuarto, claves[2])
	require.EqualValues(t, quinto, claves[5])
	require.EqualValues(t, sexto, claves[4])

}

func TestABBPruebaIterarTrasBorrados(t *testing.T) {
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"

	dic := TDADiccionario.CrearABB[string, string](func(a string, b string) int {
		if a == b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	})
	dic.Guardar(clave1, "")
	dic.Guardar(clave2, "")
	dic.Guardar(clave3, "")
	dic.Borrar(clave1)
	dic.Borrar(clave2)
	dic.Borrar(clave3)
	iter := dic.Iterador()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	dic.Guardar(clave1, "A")
	iter = dic.Iterador()

	require.True(t, iter.HaySiguiente())
	c1, v1 := iter.VerActual()
	require.EqualValues(t, clave1, c1)
	require.EqualValues(t, "A", v1)
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}

func TestIterarRangoABBVacio(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a int, b int) int { return a - b })
	desde := 0
	hasta := 99
	iter := dic.IteradorRango(&desde, &hasta)
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIterarRangoDesdeMayorQueHasta(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a int, b int) int { return a - b })
	desde := 99
	hasta := 0
	iter := dic.IteradorRango(&desde, &hasta)
	for i := 0; i < 20; i++ {
		dic.Guardar(i, i)
	}
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestABBDiccionarioIterarRango(t *testing.T) {
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](func(a string, b string) int {
		if a == b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	})
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	iter := dic.IteradorRango(&claves[1], &claves[2])

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscarABB(primero, claves))

	iter.Siguiente()
	segundo, segundo_valor := iter.VerActual()
	require.NotEqualValues(t, -1, buscarABB(segundo, claves))
	require.EqualValues(t, valores[buscarABB(segundo, claves)], segundo_valor)
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIterarFueraDeRango(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	for i := 0; i < 20; i++ {
		dic.Guardar(i, i)
	}

	desde := 32
	hasta := 100
	iter := dic.IteradorRango(&desde, &hasta)
	require.False(t, iter.HaySiguiente())

}

func TestABBDosIteradoresRango(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccionario.CrearABB[string, string](func(a string, b string) int {
		if a == b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	})
	claves := []string{"B", "A", "C"}
	dic.Guardar(claves[0], "")
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[2], "")

	iter := dic.IteradorRango(&claves[1], &claves[0])
	iter2 := dic.IteradorRango(&claves[0], &claves[2])
	iter_primero, _ := iter.VerActual()
	iter.Siguiente()
	iter_segundo, _ := iter.VerActual()
	iter2_primero, _ := iter2.VerActual()
	iter2.Siguiente()
	iter2_segundo, _ := iter2.VerActual()
	iter.Siguiente()
	iter2.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.False(t, iter2.HaySiguiente())

	require.NotEqualValues(t, iter_primero, iter2_primero)
	require.EqualValues(t, iter_segundo, iter2_primero)
	require.NotEqualValues(t, iter_segundo, iter2_segundo)
	require.EqualValues(t, iter_primero, claves[1])
	require.EqualValues(t, iter_segundo, claves[0])
	require.EqualValues(t, iter2_primero, claves[0])
	require.EqualValues(t, iter2_segundo, claves[2])
}

func TestIterarRangoDesdeYHastaIguales(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, string](func(a string, b string) int {
		if a == b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	})

	dic.Guardar("C", "")
	dic.Guardar("A", "")
	dic.Guardar("B", "")
	dic.Guardar("D", "")
	desde := "B"

	iter := dic.IteradorRango(&desde, &desde)

	require.True(t, iter.HaySiguiente())
	clave, _ := iter.VerActual()
	require.Equal(t, clave, "B")
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIterarRangoTodoElABB(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a int, b int) int { return a - b })
	for i := 0; i < 50; i++ {
		dic.Guardar(i, i)
	}
	desde := 0
	hasta := 100
	iterRango := dic.IteradorRango(&desde, &hasta)
	iter := dic.Iterador()
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		clave_rango, _ := iterRango.VerActual()
		require.Equal(t, clave_rango, clave)
		iter.Siguiente()
		iterRango.Siguiente()
	}
	require.False(t, iter.HaySiguiente())
	require.False(t, iterRango.HaySiguiente())
}

func TestIterarRangoABBCiertosElementosEnOrdenDescendente(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return b - a })
	dic.Guardar(32, 2)
	dic.Guardar(22, 2)
	dic.Guardar(1, 2)
	dic.Guardar(13, 2)
	dic.Guardar(17, 2)
	dic.Guardar(5, 2)

	desde := 23
	hasta := 2
	iter := dic.IteradorRango(&desde, &hasta)
	clave, _ := iter.VerActual()
	require.Equal(t, 22, clave)
	iter.Siguiente()
	clave, _ = iter.VerActual()
	require.Equal(t, 17, clave)
	iter.Siguiente()
	clave, _ = iter.VerActual()
	require.Equal(t, 13, clave)
	iter.Siguiente()
	clave, _ = iter.VerActual()
	require.Equal(t, 5, clave)
	iter.Siguiente()

}
func TestIterarRangoABBOrdenDescendente(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return b - a })
	for i := 0; i < 50; i++ {
		dic.Guardar(i, i)
	}
	desde := 23
	hasta := 2
	iter := dic.IteradorRango(&desde, &hasta)

	var claves_esperadas []int
	for i := 23; i > 1; i-- {
		claves_esperadas = append(claves_esperadas, i)
	}

	for _, clave_esperada := range claves_esperadas {
		clave, _ := iter.VerActual()
		require.Equal(t, clave_esperada, clave)
		iter.Siguiente()
	}
}

func TestIterarRangoVolumen(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a int, b int) int { return a - b })
	for i := 0; i < 7000; i++ {
		dic.Guardar(i, i)
	}
	desde := 1900
	hasta := 6700
	iter := dic.IteradorRango(&desde, &hasta)
	contador := 1900
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		require.True(t, clave >= desde)
		require.True(t, clave <= hasta)
		require.Equal(t, contador, clave)
		contador++
		iter.Siguiente()
	}
	require.Equal(t, 6701, contador)
}

func ejecutarPruebasVolumenIteradorABB(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[string, *int](func(a string, b string) int {
		if a == b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	})

	claves := make([]string, n)
	valores := make([]int, n)

	for i := 0; i < n; i++ {
		claves[i] = fmt.Sprintf("%08d", i)
		valores[i] = i
		dic.Guardar(claves[i], &valores[i])
	}

	iter := dic.Iterador()
	require.True(b, iter.HaySiguiente())

	ok := true
	var i int
	var clave string
	var valor *int

	for i = 0; i < n; i++ {
		if !iter.HaySiguiente() {
			ok = false
			break
		}
		c1, v1 := iter.VerActual()
		clave = c1
		if clave == "" {
			ok = false
			break
		}
		valor = v1
		if valor == nil {
			ok = false
			break
		}
		*valor = n
		iter.Siguiente()
	}
	require.True(b, ok, "Iteracion en volumen no funciona correctamente")
	require.EqualValues(b, n, i, "No se recorrió todo el largo")
	require.False(b, iter.HaySiguiente(), "El iterador debe estar al final luego de recorrer")

	ok = true
	for i = 0; i < n; i++ {
		if valores[i] != n {
			ok = false
			break
		}
	}
	require.True(b, ok, "No se cambiaron todos los elementos")
}

func BenchmarkIteradorABB(b *testing.B) {
	b.Log("Prueba de stress del Iterador del Diccionario. Prueba guardando distinta cantidad de elementos " +
		"(muy grandes) b.N elementos, iterarlos todos sin problemas. Se ejecuta cada prueba b.N veces para generar " +
		"un benchmark")
	for _, n := range TAMS_VOLUMEN_ABB {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebasVolumenIteradorABB(b, n)
			}
		})
	}
}

func TestABBVolumenIteradorCorte(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	dic := TDADiccionario.CrearABB[int, int](func(a int, b int) int { return a - b })

	for i := 0; i < 5000; i++ {
		dic.Guardar(i, i)
	}

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	dic.Iterar(func(c int, v int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c%100 == 0 {
			seguirEjecutando = false
			return false
		}
		return true
	})

	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia, "No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}

func TestIterarRango(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })

	for i := 0; i < 20; i++ {
		dic.Guardar(i, i)
	}

	desde := 8
	hasta := 15
	claves := []int{}
	esperado := []int{8, 9, 10, 11, 12, 13, 14, 15}
	dic.IterarRango(&desde, &hasta, func(clave, dato int) bool {
		claves = append(claves, clave)
		return true
	})

	require.Equal(t, len(claves), len(esperado))
	require.Equal(t, esperado, claves)
}

func TestIterarRangoSumaElementos(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })

	for i := 0; i < 20; i++ {
		dic.Guardar(i, i)
	}

	desde := 8
	hasta := 15
	sum := 0
	dic.IterarRango(&desde, &hasta, func(clave, dato int) bool {
		sum += dato
		return true
	})

	require.Equal(t, 92, sum)
}

func TestIterarRangoConCorte(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })

	for i := 0; i < 20; i++ {
		dic.Guardar(i, i)
	}

	desde := 8
	hasta := 15
	claves_visitadas := []int{}

	dic.IterarRango(&desde, &hasta, func(clave, dato int) bool {
		if clave != 10 {
			claves_visitadas = append(claves_visitadas, clave)
			return true
		}
		return false
	})

	require.Equal(t, 2, len(claves_visitadas))
	require.Equal(t, []int{8, 9}, claves_visitadas)
}

func TestIterarRangoSinDesde(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })

	for i := 0; i < 20; i++ {
		dic.Guardar(i, i)
	}

	desde := (*int)(nil)
	hasta := 15
	claves := []int{}
	esperado := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	dic.IterarRango(desde, &hasta, func(clave, dato int) bool {
		claves = append(claves, clave)
		return true
	})

	require.Equal(t, len(claves), len(esperado))
	require.Equal(t, esperado, claves)
}

func TestIterarRangoSinHasta(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })

	for i := 0; i < 20; i++ {
		dic.Guardar(i, i)
	}

	desde := 8
	hasta := (*int)(nil)
	claves := []int{}
	esperado := []int{8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	dic.IterarRango(&desde, hasta, func(clave, dato int) bool {
		claves = append(claves, clave)
		return true
	})

	require.Equal(t, len(claves), len(esperado))
	require.Equal(t, esperado, claves)
}

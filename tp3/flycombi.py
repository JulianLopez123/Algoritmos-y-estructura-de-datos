import sys
from TDAGrafo.grafo import Grafo
import queue
from collections import deque


def cargar_grafo(archivo_aeropuertos, archivo_vuelos):
    grafo = Grafo(es_dirigido=True)

    with open(archivo_aeropuertos, "r") as aeropuertos:
        for aeropuerto in aeropuertos:
            aeropuerto = aeropuerto.rstrip("\n").split(",")
            ciudad, codigo_aeropuerto, latitud, longitud = (
                aeropuerto[0],
                aeropuerto[1],
                aeropuerto[2],
                aeropuerto[3],
            )
            grafo.agregar_vertice(
                {
                    codigo_aeropuerto: {
                        "ciudad": ciudad,
                        "latitud": latitud,
                        "longitud": longitud,
                    }
                },
            )
    with open(archivo_vuelos, "r") as vuelos:
        for vuelo in vuelos:
            vuelo = vuelo.rstrip("\n").split(",")
            (
                aeropuerto_i,
                aeropuerto_j,
                tiempo_promedio,
                precio,
                cant_vuelos_entre_aeropuertos,
            ) = (vuelo[0], vuelo[1], vuelo[2], vuelo[3], vuelo[4])
            grafo.agregar_arista(
                aeropuerto_i,
                aeropuerto_j,
                {
                    "tiempo_promedio": tiempo_promedio,
                    "precio": precio,
                    "cant_vuelos_entre_aeropuertos": cant_vuelos_entre_aeropuertos,
                },
            )
    return grafo


def camino_mas(grafo, parametros):
    if len(parametros) > 3 or len(parametros) == 0:
        print("error")
        return 2
    forma = parametros[0]
    if forma != "barato" and forma != "rapido":
        print("error")
        return 3

    origen = parametros[1]
    destino = parametros[2]
    camino = camino_minimo_barato_o_rapido(grafo, forma, origen, destino)
    if camino == None:
        print("error")
        return 4

    mostrar_camino(camino)


def mostrar_camino(camino):
    for i in range(len(camino)):
        print(camino[i])
        if i != len(camino):
            print(" -> ")


def reconstruir_camino(padre, origen, destino):
    resultado = []
    while padre[destino] != None:
        resultado.append(padre[destino])
        destino = padre[destino]
    return resultado


def camino_minimo_barato_o_rapido(grafo, forma, origen, destino):
    dist = {}
    padre = {}
    for v in grafo:
        dist[v] = float("inf")
    dist[origen] = 0
    padre[origen] = None
    heap = queue.PriorityQueue()
    heap.put((0, origen))
    while not heap.empty():
        _, v = heap.get()
        if v == destino:
            return reconstruir_camino(padre, destino)
        for w in grafo.adyacentes(v):
            distancia_por_aca = dist[v] + grafo.peso(v, w)[forma]
            if distancia_por_aca < dist[w]:
                dist[w] = distancia_por_aca
                padre[w] = v
                heap.encolar((dist[w], w))
                # o: heap.actualizar(w, dist[w])
    return None


def camino_minimo_escalas(grafo, origen, destino):
    visitados = set()
    padres = {}
    orden = {}
    padres[origen] = None
    orden[origen] = 0
    visitados.add(origen)
    cola = deque()
    cola.append(origen)
    while cola:
        v = cola.popleft()
        if v == destino:
            return reconstruir_camino(padres, destino)
        for w in grafo.adyacentes(v):
            if w not in visitados:
                padres[w] = v
                orden[w] = orden[v] + 1
                visitados.add(w)
                cola.append(w)
    return orden


def camino_escalas(grafo, parametros):
    if len(parametros) != 2:
        print("error")
        return 5

    origen = parametros[0]
    destino = parametros[1]
    camino = camino_minimo_escalas(grafo, origen, destino)
    if camino == None:
        print("error")
        return 4

    mostrar_camino(camino)


# def centralidad(grafo,parametros):


def main():
    entrada = sys.argv
    if len(entrada) < 3:
        print("error")
        return 1

    archivo_aeropuertos = entrada[1]
    archivo_vuelos = entrada[2]

    grafo = cargar_grafo(archivo_aeropuertos, archivo_vuelos)

    while True:
        ingreso = sys.stdin.read()
        ingreso = ingreso.rstrip("\n").split(" ")
        comando = ingreso[0]
        parametros = ingreso[1].split(",")

        if comando == "camino_mas":
            camino_mas(grafo, parametros)
        elif comando == "camino_escalas":
            camino_escalas(grafo, parametros)
        elif comando == "centralidad":
            centralidad(grafo, parametros)
        elif comando == "nueva_aerolinea":
            nueva_aerolinea(grafo, parametros)
        elif comando == "itinerario":
            itinerario(grafo, parametros)
        elif comando == "exportar_kml":
            exportar_kml(grafo, parametros)
        else:
            print("error")
            return 6


main()

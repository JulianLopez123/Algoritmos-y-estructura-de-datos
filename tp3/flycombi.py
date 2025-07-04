import sys
from TDAGrafo.grafo import Grafo
import queue
from collections import deque


def cargar_grafo(archivo_aeropuertos, archivo_vuelos):
    aeropuertos_de_ciudades = {}
    info_aeropuertos = {}
    info_vuelos = {}

    grafo_precio = Grafo(es_dirigido=True)
    grafo_tiempo_promedio = Grafo(es_dirigido=True)
    grafo_cant_vuelos_entre_aeropuertos = Grafo(es_dirigido=True)

    with open(archivo_aeropuertos, "r") as aeropuertos:
        for aeropuerto in aeropuertos:
            aeropuerto = aeropuerto.rstrip("\n").split(",")
            ciudad, codigo_aeropuerto, latitud, longitud = (
                aeropuerto[0],
                aeropuerto[1],
                aeropuerto[2],
                aeropuerto[3],
            )
            if ciudad not in aeropuertos_de_ciudades:
                aeropuertos_de_ciudades[ciudad] = []
            aeropuertos_de_ciudades[ciudad].append(codigo_aeropuerto)
            info_aeropuertos[codigo_aeropuerto] = {
                "ciudad": ciudad,
                "latitud": latitud,
                "longitud": longitud,
            }
            grafo_precio.agregar_vertice(codigo_aeropuerto)
            grafo_tiempo_promedio.agregar_vertice(codigo_aeropuerto)
            grafo_cant_vuelos_entre_aeropuertos.agregar_vertice(codigo_aeropuerto)

    with open(archivo_vuelos, "r") as vuelos:
        for vuelo in vuelos:
            vuelo = vuelo.rstrip("\n").split(",")
            (
                aeropuerto_i,
                aeropuerto_j,
                tiempo_promedio,
                precio,
                cant_vuelos_entre_aeropuertos,
            ) = (vuelo[0], vuelo[1], int(vuelo[2]), int(vuelo[3]), int(vuelo[4]))

            info_vuelos[(aeropuerto_i, aeropuerto_j)] = {
                "tiempo_promedio": tiempo_promedio,
                "precio": precio,
                "cant_vuelos_entre_aeropuertos": cant_vuelos_entre_aeropuertos,
            }

            grafo_precio.agregar_arista(aeropuerto_i, aeropuerto_j, precio)
            grafo_tiempo_promedio.agregar_arista(
                aeropuerto_i, aeropuerto_j, tiempo_promedio
            )
            grafo_cant_vuelos_entre_aeropuertos.agregar_arista(
                aeropuerto_i, aeropuerto_j, cant_vuelos_entre_aeropuertos
            )
    return (
        grafo_precio,
        grafo_tiempo_promedio,
        grafo_cant_vuelos_entre_aeropuertos,
        info_aeropuertos,
        aeropuertos_de_ciudades,
        info_vuelos,
    )


def camino_mas(
    grafo_precio,
    grafo_tiempo_promedio,
    parametros,
    aeropuertos_de_ciudades,
):
    if len(parametros) != 3:
        print("error_cant param")
        return 2
    forma = parametros[0]
    origen = parametros[1]
    destino = parametros[2]
    if forma == "barato":
        grafo = grafo_precio
    elif forma == "rapido":
        grafo = grafo_tiempo_promedio
    else:
        return 3

    padres_minimos = {}
    aeropuerto_origen_minimo = ""
    aeropuerto_destino_minimo = ""
    camino = None
    distancia_minima = float("inf")
    for aeropuerto_origen in aeropuertos_de_ciudades[origen]:
        padres, distancia = camino_minimo_barato_o_rapido(grafo, aeropuerto_origen)
        for aeropuerto_destino in aeropuertos_de_ciudades[destino]:
            if aeropuerto_destino not in aeropuertos_de_ciudades[destino]:
                continue
            if distancia[aeropuerto_destino] < distancia_minima:
                distancia_minima = distancia[aeropuerto_destino]
                padres_minimos = padres
                aeropuerto_destino_minimo = aeropuerto_destino
                aeropuerto_origen_minimo = aeropuerto_origen

    camino = reconstruir_camino(
        padres_minimos, aeropuerto_origen_minimo, aeropuerto_destino_minimo
    )
    if camino == None:
        print("error no hay camino")
        return 4

    mostrar_camino(camino)


def mostrar_camino(camino):
    imprimible = ""
    for i in range(len(camino)):
        if i != len(camino) - 1:
            imprimible += camino[i] + " -> "
        else:
            imprimible += camino[i]
    print(imprimible)


def reconstruir_camino(padre, origen, destino):
    resultado = []
    aeropuerto = destino
    while aeropuerto:
        resultado.append(aeropuerto)
        aeropuerto = padre[aeropuerto]
    return resultado[::-1]


def camino_escalas(grafo, parametros, aeropuertos_de_ciudades):
    if len(parametros) != 2:
        print("error")
        return 5

    origen = parametros[0]
    destino = parametros[1]

    padres_minimos = {}
    aeropuerto_origen_minimo = ""
    aeropuerto_destino_minimo = ""
    camino = None
    distancia_minima = float("inf")
    for aeropuerto_origen in aeropuertos_de_ciudades[origen]:
        padres, distancia = camino_minimo_escalas(grafo, aeropuerto_origen)
        for aeropuerto_destino in aeropuertos_de_ciudades[destino]:
            if aeropuerto_destino not in aeropuertos_de_ciudades[destino]:
                continue
            if distancia[aeropuerto_destino] < distancia_minima:
                distancia_minima = distancia[aeropuerto_destino]
                padres_minimos = padres
                aeropuerto_destino_minimo = aeropuerto_destino
                aeropuerto_origen_minimo = aeropuerto_origen

    camino = reconstruir_camino(
        padres_minimos, aeropuerto_origen_minimo, aeropuerto_destino_minimo
    )
    if camino == None:
        print("error no hay camino")
        return 4

    mostrar_camino(camino)


# def centralidad(grafo,parametros):


def camino_minimo_barato_o_rapido(grafo, origen):
    dist = {}
    padre = {}
    for v in grafo.obtener_vertices():
        dist[v] = float("inf")
    dist[origen] = 0
    padre[origen] = None
    heap = queue.PriorityQueue()
    heap.put((0, origen))
    while not heap.empty():
        _, v = heap.get()
        # if v == destino:
        #     return reconstruir_camino(padre, destino)
        for w in grafo.adyacentes(v):
            distancia_por_aca = dist[v] + grafo.peso_arista(v, w)
            if distancia_por_aca < dist[w]:
                dist[w] = distancia_por_aca
                padre[w] = v
                heap.put((dist[w], w))
    return padre, dist


def camino_minimo_escalas(grafo, origen):
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
        # if v == destino:
        #     return reconstruir_camino(padres, destino)
        for w in grafo.adyacentes(v):
            if w not in visitados:
                padres[w] = v
                orden[w] = orden[v] + 1
                visitados.add(w)
                cola.append(w)
    return padres, orden


def mst_prim(grafo):
    v = grafo.vertice_aleatorio()
    visitados = set()
    visitados.add(v)
    q = queue.PriorityQueue()
    for w in grafo.adyacentes(v):
        q.put(((v, w), grafo.peso_arista(v, w)))
    arbol = Grafo(es_dirigido=False, lista_vertices=grafo.obtener_vertices())
    while not q.empty():
        (v, w), peso = q.get()
        if w in visitados:
            continue
        arbol.agregar_arista(v, w, peso)
        visitados.add(w)
        for x in grafo.adyacentes(w):
            if x not in visitados:
                q.put(((w, x), grafo.peso_arista(w, x)))
    return arbol


def obtener_vuelos(grafo):
    aristas = []
    visitados = set()
    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            if w not in visitados:
                aristas.append((v, w, grafo.peso_arista(v, w)))
        visitados.add(v)
    return aristas


def escribir_archivo_vuelos(vuelos, ruta, informacion_vuelos):
    with open(ruta, "w") as archivo:
        for vuelo in vuelos:
            aeropuerto_origen = vuelo[0]
            aeropuerto_destino = vuelo[1]
            aeropuertos_conectados = (aeropuerto_origen, aeropuerto_destino)
            if aeropuertos_conectados not in informacion_vuelos:
                aeropuertos_conectados = (aeropuerto_destino, aeropuerto_origen)
            info_vuelo = informacion_vuelos[aeropuertos_conectados]
            linea = f"{aeropuertos_conectados[0]},{aeropuertos_conectados[1]},{info_vuelo["tiempo_promedio"]},{info_vuelo["precio"]},{info_vuelo["cant_vuelos_entre_aeropuertos"]}\n"
            archivo.write(linea)
    print("OK")


def nueva_aerolinea(grafo, parametros, informacion_vuelos):
    if len(parametros) != 1:
        print("error cant param nueva aerolinea")
        return 6
    ruta = parametros[0]
    arbol = mst_prim(grafo)
    vuelos = obtener_vuelos(arbol)
    escribir_archivo_vuelos(vuelos, ruta, informacion_vuelos)


def main():
    entrada = sys.argv
    if len(entrada) < 3:
        print("error")
        return 1

    archivo_aeropuertos = entrada[1]
    archivo_vuelos = entrada[2]

    (
        grafo_precio,
        grafo_tiempo_promedio,
        grafo_cant_vuelos_entre_aeropuertos,
        info_aeropuertos,
        aeropuertos_de_ciudades,
        info_vuelos,
    ) = cargar_grafo(archivo_aeropuertos, archivo_vuelos)

    while True:
        ingreso = input()
        ingreso = ingreso.rstrip().split(" ", maxsplit=1)
        comando = ingreso[0]
        parametros = ingreso[1].split(",")

        if comando == "camino_mas":
            camino_mas(
                grafo_precio,
                grafo_tiempo_promedio,
                parametros,
                aeropuertos_de_ciudades,
            )
        elif comando == "camino_escalas":
            camino_escalas(
                grafo_cant_vuelos_entre_aeropuertos,
                parametros,
                aeropuertos_de_ciudades,
            )
        # elif comando == "centralidad":
        #     centralidad(grafo, parametros)
        elif comando == "nueva_aerolinea":
            nueva_aerolinea(grafo_precio, parametros, info_vuelos)
        # elif comando == "itinerario":
        #     itinerario(grafo, parametros)
        # elif comando == "exportar_kml":
        #     exportar_kml(grafo, parametros)
        else:
            print("error")
            return 6


main()

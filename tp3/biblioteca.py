from queue import PriorityQueue, LifoQueue
from collections import deque
from TDAGrafo.grafo import Grafo


def dijkstra(grafo, origen, destino=None):
    dist = {}
    padre = {}
    for v in grafo.obtener_vertices():
        dist[v] = float("inf")
    dist[origen] = 0
    padre[origen] = None
    heap = PriorityQueue()
    heap.put((0, origen))
    while not heap.empty():
        _, v = heap.get()
        if v == destino:
            break
        for w in grafo.adyacentes(v):
            distancia_por_aca = dist[v] + grafo.peso_arista(v, w)
            if distancia_por_aca < dist[w]:
                dist[w] = distancia_por_aca
                padre[w] = v
                heap.put((dist[w], w))
    return padre, dist


def bfs(grafo, origen, destino=None):
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
    q = PriorityQueue()
    for w in grafo.adyacentes(v):
        q.put((grafo.peso_arista(v, w), v, w))
    arbol = Grafo(es_dirigido=False, lista_vertices=grafo.obtener_vertices())
    while not q.empty():
        peso, v, w = q.get()
        if w in visitados:
            continue
        arbol.agregar_arista(v, w, peso)
        visitados.add(w)
        for x in grafo.adyacentes(w):
            if x not in visitados:
                q.put((grafo.peso_arista(w, x), w, x))
    return arbol


def betweeness_centrality(grafo):
    cent = {}
    for v in grafo.obtener_vertices():
        cent[v] = 0
    for v in grafo.obtener_vertices():
        padre, distancias = dijkstra(grafo, v)
        cent_aux = {}
        for w in grafo.obtener_vertices():
            cent_aux[w] = 0
        vertices_ordenados = sorted(
            distancias, key=lambda i: distancias[i], reverse=True
        )
        for w in vertices_ordenados:
            if padre[w] == None:
                continue
            cent_aux[padre[w]] += 1 + cent_aux[w]
        for w in grafo.obtener_vertices():
            if w == v:
                continue
            cent[w] += cent_aux[w]
    return cent


def _dfs(grafo, v, visitados, pila):
    for w in grafo.adyacentes(v):
        if w not in visitados:
            visitados.add(w)
            _dfs(grafo, w, visitados, pila)
    pila.put(v)


def topologico_dfs(grafo):
    visitados = set()
    pila = LifoQueue()
    for v in grafo.obtener_vertices():
        if v not in visitados:
            visitados.add(v)
            _dfs(grafo, v, visitados, pila)
    return pila_a_lista(pila)


def pila_a_lista(pila):
    lista = []
    while not pila.empty():
        lista.append(pila.get())
    return lista


def cargar_grafo(archivo_aeropuertos, archivo_vuelos):
    aeropuertos_de_ciudades = {}
    info_aeropuertos = {}
    info_vuelos = {}

    grafo_precio = Grafo(es_dirigido=False)
    grafo_tiempo_promedio = Grafo(es_dirigido=False)
    grafo_escalas = Grafo(es_dirigido=False)
    grafo_escalas_inverso = Grafo(es_dirigido=False)

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
            grafo_escalas.agregar_vertice(codigo_aeropuerto)
            grafo_escalas_inverso.agregar_vertice(codigo_aeropuerto)

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
            grafo_escalas.agregar_arista(
                aeropuerto_i, aeropuerto_j, cant_vuelos_entre_aeropuertos
            )
            grafo_escalas_inverso.agregar_arista(
                aeropuerto_i, aeropuerto_j, 1 / cant_vuelos_entre_aeropuertos
            )
    return (
        grafo_precio,
        grafo_tiempo_promedio,
        grafo_escalas,
        grafo_escalas_inverso,
        info_aeropuertos,
        aeropuertos_de_ciudades,
        info_vuelos,
    )


def camino_minimo(
    grafo,
    origen,
    destino,
    aeropuertos_de_ciudades,
    funcion,
):
    padres_minimos = {}
    aeropuerto_origen_minimo = ""
    aeropuerto_destino_minimo = ""
    distancia_minima = float("inf")
    for aeropuerto_origen in aeropuertos_de_ciudades[origen]:
        padres, distancias = funcion(grafo, aeropuerto_origen)
        for aeropuerto_destino in aeropuertos_de_ciudades[destino]:
            if aeropuerto_destino not in distancias:
                continue
            if distancias[aeropuerto_destino] < distancia_minima:
                distancia_minima = distancias[aeropuerto_destino]
                padres_minimos = padres
                aeropuerto_destino_minimo = aeropuerto_destino
                aeropuerto_origen_minimo = aeropuerto_origen

    if distancia_minima == float("inf"):
        return None, None, None
    return padres_minimos, aeropuerto_origen_minimo, aeropuerto_destino_minimo


def camino_mas(
    grafo_precio,
    grafo_tiempo_promedio,
    parametros,
    aeropuertos_de_ciudades,
):
    if len(parametros) != 3:
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

    padres_minimos, aeropuerto_origen_minimo, aeropuerto_destino_minimo = camino_minimo(
        grafo, origen, destino, aeropuertos_de_ciudades, dijkstra
    )
    if padres_minimos == None:
        return 4

    camino = reconstruir_camino(padres_minimos, aeropuerto_destino_minimo)

    linea = armar_camino_imprimible(camino)
    print(linea)
    return linea


def armar_camino_imprimible(camino):
    linea = ""
    for i in range(len(camino)):
        if i != len(camino) - 1:
            linea += camino[i] + " -> "
        else:
            linea += camino[i]
    return linea


def reconstruir_camino(padre, destino):
    if padre == None:
        return None
    resultado = []
    aeropuerto = destino
    while aeropuerto:
        resultado.append(aeropuerto)
        aeropuerto = padre[aeropuerto]
    return resultado[::-1]


def camino_escalas(grafo, parametros, aeropuertos_de_ciudades):
    if len(parametros) != 2:
        return 5

    origen = parametros[0]
    destino = parametros[1]

    padres_minimos, aeropuerto_origen_minimo, aeropuerto_destino_minimo = camino_minimo(
        grafo, origen, destino, aeropuertos_de_ciudades, bfs
    )
    camino = reconstruir_camino(padres_minimos, aeropuerto_destino_minimo)
    if camino == None:
        return 4

    linea = armar_camino_imprimible(camino)
    print(linea)
    return linea


def obtener_vuelos(grafo):
    vuelos = []
    visitados = set()
    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            if w not in visitados:
                vuelos.append((v, w, grafo.peso_arista(v, w)))
        visitados.add(v)
    return vuelos


def escribir_archivo_vuelos(vuelos, ruta, informacion_vuelos):
    with open(ruta, "w") as archivo:
        for vuelo in vuelos:
            aeropuerto_origen = vuelo[0]
            aeropuerto_destino = vuelo[1]
            aeropuertos_conectados = (aeropuerto_origen, aeropuerto_destino)
            if aeropuertos_conectados not in informacion_vuelos:
                aeropuertos_conectados = (aeropuerto_destino, aeropuerto_origen)
            info_vuelo = informacion_vuelos[aeropuertos_conectados]
            linea = f"{aeropuertos_conectados[0]},{aeropuertos_conectados[1]},{info_vuelo['tiempo_promedio']},{info_vuelo['precio']},{info_vuelo['cant_vuelos_entre_aeropuertos']}\n"
            archivo.write(linea)
    print("OK")


def nueva_aerolinea(grafo, parametros, informacion_vuelos):
    if len(parametros) != 1:
        return 7
    ruta = parametros[0]
    arbol = mst_prim(grafo)
    vuelos = obtener_vuelos(arbol)
    escribir_archivo_vuelos(vuelos, ruta, informacion_vuelos)


def centralidad(grafo, parametros):
    if len(parametros) != 1:
        return 9
    if not parametros[0].isnumeric():
        return 10
    cantidad_aeropuertos = int(parametros[0])
    centralidad = betweeness_centrality(grafo)
    aeropuertos_centrales = sorted(
        centralidad.items(), key=lambda aeropuerto: aeropuerto[1], reverse=True
    )
    linea = ""
    for i in range(cantidad_aeropuertos):
        linea += aeropuertos_centrales[i][0] + ", "
    print(linea)


def obtener_itinerario(ruta):
    itinerario = Grafo(es_dirigido=True)
    with open(ruta, "r") as archivo:
        ciudades_a_visitar = archivo.readline().strip().split(",")
        for ciudad in ciudades_a_visitar:
            itinerario.agregar_vertice(ciudad)
        for linea in archivo:
            if not linea:
                continue
            prioridades = linea.strip().split(",")
            itinerario.agregar_arista(prioridades[0], prioridades[1])
    return itinerario


def itinerario(grafo_escalas, parametros, aeropuertos_de_ciudades):
    if len(parametros) != 1:
        return 8
    ruta = parametros[0]
    itinerario = obtener_itinerario(ruta)
    orden = topologico_dfs(itinerario)
    ciudades_en_orden = ",".join(orden)
    print(ciudades_en_orden)
    ciudades_en_orden = ciudades_en_orden.split(",")

    for i in range(len(ciudades_en_orden) - 1):
        origen = ciudades_en_orden[i]
        destino = ciudades_en_orden[i + 1]
        padres_minimos, aeropuerto_origen_minimo, aeropuerto_destino_minimo = (
            camino_minimo(
                grafo_escalas,
                origen,
                destino,
                aeropuertos_de_ciudades,
                bfs,
            )
        )
        camino = reconstruir_camino(padres_minimos, aeropuerto_destino_minimo)
        if camino == None:
            return 4

        linea = armar_camino_imprimible(camino)
        print(linea)


def exportar_kml(parametros, ruta_kml, info_aeropuertos):
    if len(parametros) != 1 or ruta_kml == "":
        return 9

    archivo = parametros[0]
    escribir_kml(archivo, info_aeropuertos, ruta_kml)


def escribir_lineas_kml(archivo, aeropuertos, info_aeropuertos):
    for i in range(len(aeropuertos) - 1):
        latitud_origen = info_aeropuertos[aeropuertos[i]]["latitud"]
        longitud_origen = info_aeropuertos[aeropuertos[i]]["longitud"]
        latitud_destino = info_aeropuertos[aeropuertos[i + 1]]["latitud"]
        longitud_destino = info_aeropuertos[aeropuertos[i + 1]]["longitud"]
        archivo.write("		<Placemark>\n")
        archivo.write("        <LineString>\n")
        archivo.write(
            f"                <coordinates>{latitud_origen} {longitud_origen} {latitud_destino} {longitud_destino}</coordinates>\n"
        )
        archivo.write("        </LineString>\n")
        archivo.write("		</Placemark>\n")


def escribir_lugares_kml(archivo, aeropuertos, info_aeropuertos):
    for aeropuerto in aeropuertos:
        ciudad = info_aeropuertos[aeropuerto]["ciudad"]
        latitud = info_aeropuertos[aeropuerto]["latitud"]
        longitud = info_aeropuertos[aeropuerto]["longitud"]
        archivo.write("		<Placemark>\n")
        archivo.write(f"			<name>{ciudad}</name>\n")
        archivo.write("			<Point>\n")
        archivo.write(
            f"                <coordinates>{latitud}, {longitud}</coordinates>\n"
        )
        archivo.write("			</Point>\n")
        archivo.write("		</Placemark>\n\n")


def escribir_inicio_kml(archivo, ciudad_origen, ciudad_destino):
    archivo.write('<?xml version="1.0" encoding="UTF-8"?>\n')
    archivo.write('<kml xmlns="http://earth.google.com/kml/2.1">\n')
    archivo.write("	<Document>\n")
    archivo.write(
        f"        <name>Camino desde {ciudad_origen} hacia {ciudad_destino}</name>\n\n"
    )


def escribir_declaracion_kml(archivo):
    archivo.write(" </Document>\n")
    archivo.write("</kml>")


def escribir_kml(archivo, info_aeropuertos, ruta_kml):
    with open(archivo, "w") as kml:
        aeropuertos = ruta_kml.split(" -> ")
        aeropuerto_origen = aeropuertos[0]
        aeropuerto_destino = aeropuertos[-1]
        ciudad_origen = info_aeropuertos[aeropuerto_origen]["ciudad"]
        ciudad_destino = info_aeropuertos[aeropuerto_destino]["ciudad"]

        escribir_inicio_kml(kml, ciudad_origen, ciudad_destino)
        escribir_lugares_kml(kml, aeropuertos, info_aeropuertos)
        escribir_lineas_kml(kml, aeropuertos, info_aeropuertos)
        escribir_declaracion_kml(kml)
    print("OK")

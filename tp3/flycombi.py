#!/usr/bin/python3
import sys
from TDAGrafo.grafo import Grafo
from biblioteca import (
    cargar_grafo,
    camino_mas,
    camino_escalas,
    centralidad,
    nueva_aerolinea,
    itinerario,
    exportar_kml,
)


def main():
    entrada = sys.argv
    instrucciones = sys.stdin
    if len(entrada) != 3:
        return 1
    archivo_aeropuertos = entrada[1]
    archivo_vuelos = entrada[2]
    ejecucion(archivo_aeropuertos, archivo_vuelos, instrucciones)


def ejecucion(archivo_aeropuertos, archivo_vuelos, instrucciones):
    ruta_kml = ""
    grafo, info_aeropuertos, info_vuelos, aeropuertos_de_ciudades = cargar_grafo(
        archivo_aeropuertos, archivo_vuelos
    )
    for linea in instrucciones:
        ingreso = linea.rstrip("\n").split(" ", maxsplit=1)
        comando = ingreso[0]
        parametros = ingreso[1].split(",")
        if comando == "camino_mas":
            ruta_kml = camino_mas(
                grafo,
                parametros,
                aeropuertos_de_ciudades,
            )
        elif comando == "camino_escalas":
            ruta_kml = camino_escalas(
                grafo,
                parametros,
                aeropuertos_de_ciudades,
            )
        elif comando == "centralidad":
            centralidad(grafo, parametros)
        elif comando == "nueva_aerolinea":
            nueva_aerolinea(grafo, parametros, info_vuelos)
        elif comando == "itinerario":
            itinerario(
                grafo,
                parametros,
                aeropuertos_de_ciudades,
            )
        elif comando == "exportar_kml":
            exportar_kml(parametros, ruta_kml, info_aeropuertos)
        else:
            return 6


main()

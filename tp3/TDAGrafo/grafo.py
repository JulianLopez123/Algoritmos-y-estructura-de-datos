import random


class Grafo:
    def __init__(self, es_dirigido, vertices_init=None):
        self.es_dirigido = es_dirigido
        self.vertices = {}
        self.aristas = {}
        if vertices_init != None:
            for v in vertices_init:
                self.agregar_vertice(v)

    def agregar_vertice(self, v):
        if v not in self.vertices.keys():
            self.vertices.add(v)
            self.aristas[v] = {}

    def borrar_vertice(self, v):
        if v in self.vertices:
            self.vertices.remove(v)
            del self.aristas[v]
            for arista in self.aristas:
                if v in self.aristas[arista]:
                    del self.aristas[arista][v]

    def agregar_arista(self, v, w, peso=1):
        if v in self.vertices and w in self.vertices:
            self.aristas[v][w] = peso
            if not self.es_dirigido:
                self.aristas[w][v] = peso

    def borrar_arista(self, v, w):
        if v in self.aristas and w in self.aristas[v]:
            del self.aristas[v][w]
            if not self.es_dirigido:
                del self.aristas[w][v]

    def estan_unidos(self, v, w):
        if v in self.aristas and w in self.aristas[v]:
            return True
        return False

    def peso_arista(self, v, w):
        if self.estan_unidos(v, w):
            return self.aristas[v][w]
        return None

    def obtener_vertices(self):
        return list(self.vertices)

    def vertice_aleatorio(self):
        if len(self.vertices) > 0:
            return random.choice(list(self.vertices))
        return None

    def adyacentes(self, v):
        if v not in self.aristas:
            return []
        return list(self.aristas[v].keys())

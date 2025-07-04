import random


class Grafo:
    def __init__(self, es_dirigido, vertices_init=[]):
        self.es_dirigido = es_dirigido
        self.aristas = {}
        if vertices_init:
            for v in vertices_init:
                self.agregar_vertice(v)

    def agregar_vertice(self, v):
        if v not in self.aristas:
            self.aristas[v] = {}

    def borrar_vertice(self, v):
        if v in self.aristas:
            del self.aristas[v]
            for w in self.aristas:
                if v in self.aristas[w]:
                    del self.aristas[w][v]

    def agregar_arista(self, v, w, peso=1):
        if v in self.aristas and w in self.aristas:
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
        return list(self.aristas)

    def vertice_aleatorio(self):
        if len(self.aristas) > 0:
            return random.choice(list(self.aristas))
        return None

    def adyacentes(self, v):
        if v not in self.aristas:
            return []
        return list(self.aristas[v].keys())

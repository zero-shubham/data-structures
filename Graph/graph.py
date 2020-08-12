# graph with adjacency list representation
class Graph:
    def __init__(self):
        self.vertices = dict()

    def add_vertex(self, value: int):
        self.vertices[value] = list()
        return self.vertices[value]

    def add_edge(self, vertex1: int, vertex2: int):
        self.vertices[vertex1].append(vertex2)
        return self.vertices[vertex1]

    def __str__(self):
        for vertex, edges in self.vertices.items():
            for edge in edges:
                print(f"{vertex} -> {edge}")
            print("\n")
        return ""

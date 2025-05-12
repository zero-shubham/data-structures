from graph import Graph


def main():
    g = Graph()
    g.add_vertex(5)
    g.add_vertex(7)
    g.add_edge(5, 7)
    g.add_vertex(3)
    g.add_edge(3, 5)
    g.add_vertex(1)
    g.add_edge(1, 7)
    g.add_vertex(2)
    g.add_edge(2, 1)
    g.add_edge(2, 7)
    g.add_edge(7, 3)
    print(g)
    print(g.vertices)


if __name__ == "__main__":
    main()

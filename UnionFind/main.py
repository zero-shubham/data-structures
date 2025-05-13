class UnionFind:
    def __init__(self, n: int):
        self.par = dict()
        self.rank = dict()

        for i in range(1, n+1):
            self.par[i] = i
            self.rank[i] = 0

    def find(self, n):
        p = self.par[n]

        while p != self.par[p]:
            self.par[p] = self.par[self.par[p]]
            p = self.par[p]

        return p

    def union(self, n1, n2) -> bool:
        p1, p2 = self.find(n1), self.find(n2)

        if p1 == p2:
            return False

        if self.rank[p1] > self.rank[p2]:
            self.par[p2] = p1
        elif self.rank[p2] > self.rank[p1]:
            self.par[p1] = p2
        else:
            self.par[p1] = p2
            self.rank[p2] = +1

        return True


def main():
    edges = [[1, 2], [1, 3], [2, 3]]
    uf = UnionFind(len(edges))

    for edge in edges:
        print("added : ", uf.union(edge[0], edge[1]))

    print(uf.par, uf.rank)


if __name__ == "__main__":
    main()

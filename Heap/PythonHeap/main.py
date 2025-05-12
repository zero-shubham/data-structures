from heap import MinHeap


def main():
    h = MinHeap()
    h.push(3)
    print(h.arr)
    h.push(5)
    print(h.arr)
    h.push(2)
    print(h.arr)
    h.push(55)
    print(h.arr)
    h.push(6)
    print(h.arr)
    h.push(1)
    print(h.arr)

    pop = h.pop()
    print(h.arr, pop)

    pop = h.pop()
    print(h.arr, pop)

    pop = h.pop()
    print(h.arr, pop)

    pop = h.pop()
    print(h.arr, pop)

    pop = h.pop()
    print(h.arr, pop)

    pop = h.pop()
    print(h.arr, pop)

    pop = h.pop()
    print(h.arr, pop)


if __name__ == "__main__":
    main()

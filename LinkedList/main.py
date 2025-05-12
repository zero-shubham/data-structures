from linked_list import LinkedList


def main():
    ll = LinkedList()
    print(ll.empty())
    ll.push_back(1)
    ll.push_back(2)
    ll.push_back(3)
    print(len(ll))
    print(ll)
    print(ll.empty())
    print(ll.value_at(2))
    print("pop_back (1) = ", ll.pop_back())
    print("pop_back (2) = ", ll.pop_back())
    print("pop_back (3) = ", ll.pop_back())
    ll.push_back(1)
    ll.push_back(2)
    ll.push_back(3)
    ll.push_front(5)
    print(ll)
    print("pop_front (1) = ", ll.pop_front())
    print(ll)
    print("pop_back (1) = ", ll.pop_back())
    print("pop_back (2) = ", ll.pop_back())
    print("pop_back (3) = ", ll.pop_back())
    print(ll)


if __name__ == "__main__":
    main()

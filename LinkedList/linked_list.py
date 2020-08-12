
class Node:
    def __init__(self, data: int, next=None):
        self.data = data
        self.next = next


class LinkedList:

    def __init__(self):
        self.root = None

    def push_back(self, value: int) -> int:
        if not self.root:
            self.root = Node(value)
        else:
            last_node = self.root
            while last_node.next:
                last_node = last_node.next
            last_node.next = Node(value)
        return value

    def pop_back(self) -> int:
        if not self.root:
            raise IndexError("LinkedList is empty")
        elif not self.root.next:
            value = self.root.data
            self.root = None
            return value

        ptr = self.root
        prev_ptr = ptr
        while ptr and ptr.next:
            prev_ptr = ptr
            ptr = ptr.next

        prev_ptr.next = None
        return ptr.data

    def __len__(self) -> int:
        count = 0
        ptr = self.root
        while ptr is not None:
            if ptr.data:
                count += 1
            ptr = ptr.next
        return count

    def __str__(self):
        ptr = self.root
        while ptr and ptr.data:
            print(f"{ptr.data} ->", end="\t")
            ptr = ptr.next
        return "None"

    def empty(self) -> bool:
        return self.root is None

    def value_at(self, index: int) -> int:
        idx = 0
        ptr = self.root
        while idx != index:
            if not ptr:
                raise IndexError("Index out of bound.")
            ptr = ptr.next
            idx += 1

        return ptr.data

    def push_front(self, value) -> int:
        new_node = Node(value, self.root)
        self.root = new_node
        return self.root.data

    def pop_front(self) -> int:
        if not self.root:
            raise IndexError("LinkedList is empty")

        popped = self.root
        self.root = self.root.next
        return popped.data

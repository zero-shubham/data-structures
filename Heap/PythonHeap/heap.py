
class MinHeap():
    def __init__(self):
        self.arr = [None]

    def leftchild(self, idx: int) -> int:
        if len(self.arr) <= idx*2:
            return None
        return idx*2

    def rightchild(self, idx: int) -> int:
        if len(self.arr) <= (idx*2) + 1:
            return None
        return (idx*2) + 1

    def get_push_parent(self, idx) -> int:
        return idx//2

    def percolate_down(self, idx):
        if idx is None:
            return

        left_child = self.leftchild(idx)
        right_child = self.rightchild(idx)

        if left_child and left_child < len(self.arr):
            self.percolate_down(left_child)

        if right_child and right_child < len(self.arr):
            self.percolate_down(right_child)

        smallest = idx

        if left_child and left_child < len(self.arr) and self.arr[left_child] < self.arr[smallest]:
            smallest = left_child
        if right_child and right_child < len(self.arr) and self.arr[right_child] < self.arr[smallest]:
            smallest = right_child

        if smallest != idx:
            tmp = self.arr[smallest]
            self.arr[smallest] = self.arr[idx]
            self.arr[idx] = tmp

    def percolate_up(self, ele_idx):

        curr_push_parent = self.get_push_parent(ele_idx)
        if not self.arr[curr_push_parent]:
            return

        if self.arr[ele_idx] < self.arr[curr_push_parent]:
            tmp = self.arr[ele_idx]
            self.arr[ele_idx] = self.arr[curr_push_parent]
            self.arr[curr_push_parent] = tmp

        self.percolate_up(curr_push_parent)

    def push(self, item: int):
        if len(self.arr) == 1:
            self.arr.append(item)
            return

        self.arr.append(item)
        self.percolate_up(len(self.arr)-1)

    def pop(self) -> int:
        if len(self.arr) == 1:
            return
        if len(self.arr) == 2:
            return self.arr.pop()

        pop_item = self.arr[1]

        self.arr[1] = self.arr.pop()
        self.percolate_down(1)

        return pop_item

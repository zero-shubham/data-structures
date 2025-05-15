from typing import List


class SegmentTree:

    def __init__(self, left: int, right: int, total: int):
        self.L: SegmentTree = None
        self.R: SegmentTree = None
        self.range = (left, right)
        self.sum = total

    @staticmethod
    def build(nums: List[int], L: int, R: int):
        if L == R:
            return SegmentTree(L, R, nums[L])

        M = (L+R)//2
        root = SegmentTree(L, R, 0)
        root.L = SegmentTree.build(nums, L, M)
        root.R = SegmentTree.build(nums, M+1, R)
        root.sum = root.L.sum + root.R.sum

        return root

    def update(self, index: int, val: int) -> None:
        if self.range[0] == self.range[1]:
            self.sum = val
            return

        M = (self.range[0] + self.range[1])//2

        if index > M:
            self.R.update(index, val)
        else:
            self.L.update(index, val)

        self.sum = self.L.sum + self.R.sum

    def sumRange(self, left: int, right: int) -> int:
        if self.range[0] == left and self.range[1] == right:
            return self.sum

        M = (self.range[0] + self.range[1])//2

        if left > M:
            return self.R.sumRange(left, right)
        if right <= M:
            return self.L.sumRange(left, right)
        return (self.L.sumRange(left, M)+self.R.sumRange(M+1, right))


def main():
    l = [1, 3, 5]
    st = SegmentTree.build(l, 0, len(l)-1)
    print(st.sum)
    st.update(1, 2)
    print(st.sum)


if __name__ == "__main__":
    main()

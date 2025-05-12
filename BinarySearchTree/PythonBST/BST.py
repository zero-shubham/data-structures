class BST:
    def __init__(self, initVal=None):
        self.value = initVal
        if self.value:
            self.left = BST()
            self.right = BST()
        else:
            self.left = None
            self.right = None
        return

    def isEmpty(self) -> bool:
        return (self.value == None)

    def isLeaf(self):
        return (self.left.isEmpty() and self.right.isEmpty())

    def remove(self, val):
        if self.isEmpty():
            return
        # print("removing => ", val)
        if val < self.value:
            self.left = self.left.remove(val) or BST()
        elif val > self.value:
            self.right = self.right.remove(val) or BST()
        else:
            # print("removing => ", val)
            if self.left.isEmpty():
                return self.right
            elif self.right.isEmpty():
                return self.left
            else:
                minRight = self.minVal(self.right)
                self.value = minRight.value
                self.right = self.right.remove(minRight.value)

        return self

    def insert(self, val):
        if self.isEmpty():
            self.value = val
            self.left = BST()
            self.right = BST()

        if self.value == val:
            return

        if val < self.value:
            self.left.insert(val)
            return

        if val > self.value:
            self.right.insert(val)
            return

    def inorder(self):

        if self.isEmpty():
            return ('')
        else:
            return (self.left.inorder()+'--'+str(self.value)+'--'+self.right.inorder()+'\n')

    def __str__(self):
        return str(self.inorder())

    def foo(self):
        if self.isEmpty():
            return (0)
        elif self.isLeaf():
            return (self.value)
        else:
            return (self.value + max(self.left.foo(), self.right.foo()))

    def minVal(self, node):
        curr = node

        while curr and curr.left and not curr.left.isEmpty():
            curr = curr.left

        return curr

from BST import BST

root = BST(5)
root.insert(3)
root.insert(4)

root.insert(6)
root.insert(2)


# print(str(root))

# print("\n=>",root.foo())


root = root.remove(5)
print(str(root))

root = root.remove(6)
print(str(root))

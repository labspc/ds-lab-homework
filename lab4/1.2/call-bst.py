import ctypes
from ctypes import c_int, POINTER

# 加载共享库
bst = ctypes.CDLL('./libbst.so')

# 定义 initialize_tree 和 inorder_traversal_api 函数的参数和返回类型
bst.initialize_tree.restype = ctypes.POINTER(ctypes.c_void_p)
bst.initialize_tree.argtypes = [POINTER(c_int), c_int]

bst.inorder_traversal_api.argtypes = [ctypes.POINTER(ctypes.c_void_p)]

# 定义要插入的键
keys = [16, 5, 17, 29, 11, 3, 15, 20]
n = len(keys)

# 创建一个C类型的数组
keys_array = (c_int * n)(*keys)

# 调用 initialize_tree 函数
root = bst.initialize_tree(keys_array, n)

# 调用 inorder_traversal_api 函数
bst.inorder_traversal_api(root)

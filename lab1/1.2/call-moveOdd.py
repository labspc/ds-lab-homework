import ctypes

# 加载 C 库
lib = ctypes.CDLL("./libmoveOdd_cc.so")  

# 定义 C 函数的参数类型
lib.moveOdd_api.argtypes = [
    ctypes.POINTER(ctypes.c_int),  # A 数组指针
    ctypes.c_int                   # 数组大小
]

# 定义 C 函数的返回类型
lib.moveOdd_api.restype = None

def typeTrans(A):
    # 将 Python 列表转换为 C 数组
    A_arr = (ctypes.c_int * len(A))(*A)
    # 调用 C 函数
    lib.moveOdd_api(A_arr, len(A))
    # 将修改后的结果转换回 Python 列表
    return list(A_arr)

# 示例输入
# A = [2,4,6,8,1,7,9,13]
# A = [1,4,6,7,9,8,11,4,2,3]

# 从键盘输入
A = list(map(int, input("输入顺序表A的元素（以空格分隔）：").split()))
print("原始:", A)
# 调用函数
result = typeTrans(A)
print("交换:", result)

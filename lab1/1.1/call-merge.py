import ctypes

# 加载C库
lib = ctypes.CDLL("./merge.so")

# 定义C函数的参数类型
lib.merge_api.argtypes = [
    ctypes.POINTER(ctypes.c_int),  # A数组指针
    ctypes.c_int,                   # A数组大小
    ctypes.POINTER(ctypes.c_int),  # B数组指针
    ctypes.c_int,                   # B数组大小
    ctypes.POINTER(ctypes.c_int)   # C数组指针
]

# 定义C函数的返回类型
lib.merge_api.restype = None
# 定义一个Python包装器（封装函数），进行类型转换
def typeTrans(A, B):
    # 将Python列表转换为C数组
    A_arr = (ctypes.c_int * len(A))(*A)
    B_arr = (ctypes.c_int * len(B))(*B)

    # 初始化C数组，C的大小为A和B的大小之和
    C_size = len(A) + len(B)
    C_arr = (ctypes.c_int * C_size)()

    # 调用C函数进行合并
    lib.merge_api(A_arr, len(A), B_arr, len(B), C_arr)

    # 将合并结果转换为Python列表
    C = [C_arr[i] for i in range(C_size)]
    return C

# 从键盘输入两个顺序表A和B
A = list(map(int, input("输入顺序表A的元素（以空格分隔）：").split()))
B = list(map(int, input("输入顺序表B的元素（以空格分隔）：").split()))


# 调用函数合并数组并输出结果
print("顺序表A的元素:", A)
print("顺序表B的元素:", B)
C = typeTrans(A, B)
print("合并后的顺序表C的元素:", C)

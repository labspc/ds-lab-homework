import ctypes

# 加载共享库
lib = ctypes.CDLL('./HalfRead.so')

# 设置参数和返回值类型
lib.HalfRead.argtypes = [ctypes.c_char_p]
lib.HalfRead.restype = ctypes.c_int  # cBool 实际上是 C.int

def HalfRead(s):
    return lib.HalfRead(s.encode('utf-8')) == 1

# 获取用户输入
input_str = input("请输入一个字符串: ").strip()

if HalfRead(input_str):
    print("是回文")
else:
    print("不是回文")

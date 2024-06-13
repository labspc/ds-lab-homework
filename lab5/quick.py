def quick_sort(arr):
    # 递归结束条件
    if len(arr) <= 1: # 单个元素直接返回
        return arr

    pivot = arr[0] # 选择第一个关键字作为第一个pivot

    # 小的放在左边，大的放在右边
    left = [x for x in arr[1:] if x < pivot]
    right = [x for x in arr[1:] if x >= pivot]

    # 将每次的列表连接起来 left+pivot+right
    return quick_sort(left) + [pivot] + quick_sort(right)

# 关键字序列
keywords = [21, 25, 49, 25, 16, 8]


# 快速排序
sorted_keywords_quick = quick_sort(keywords.copy())
print("快速排序结果:", sorted_keywords_quick)


def shell_sort(arr):
    # 获取数组的长度
    n = len(arr)

    # 初始的间隔（gap）为数组长度的一半
    gap = n // 2

    # 只要间隔（gap）大于0，就继续排序
    while gap > 0:
        # 从间隔（gap）开始，到数组结束
        for i in range(gap, n):
            # 暂存当前元素
            temp = arr[i]
            j = i

            # 插入排序的内部循环
            # 如果当前元素小于前面的元素，就进行交换
            while j >= gap and arr[j - gap] > temp:
                arr[j] = arr[j - gap]
                j -= gap

            # 将暂存的元素放到正确的位置
            arr[j] = temp

        # 缩小间隔（gap）
        gap //= 2

    # 返回排序后的数组
    return arr


# 关键字序列
keywords = [81, 94, 11, 96, 12, 35, 17, 95, 28, 58, 41, 75, 15]

# 希尔排序
sorted_keywords_shell = shell_sort(keywords.copy())
print("希尔排序结果:", sorted_keywords_shell)

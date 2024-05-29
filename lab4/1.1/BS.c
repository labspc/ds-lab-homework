#include <stdio.h>
#include <stdlib.h>

// 折半查找函数
int BS(int array[], int size, int k)
{
    int left = 0;
    int right = size - 1;
    while (left <= right)
    {
        int mid = left + (right - left) / 2;
        if (array[mid] == k)
        {
            return mid;
        }
        else if (array[mid] < k)
        {
            left = mid + 1;
        }
        else
        {
            right = mid - 1;
        }
    }
    return -1;
}
int main()
{
    int size, k;
    printf("请输入有序序列的大小: ");
    scanf("%d", &size);
    int array[size];
    printf("请输入有序序列的元素: ");
    // 将输入读取到数组array中
    for (int i = 0; i < size; i++)
    {
        scanf("%d", &array[i]);
    }
    printf("请输入要查找的关键字: ");
    scanf("%d", &k);
    
    int result = BS(array, size, k);
    if (result != -1)
    {
        printf("关键字 %d 的索引为 %d\n", k, result);
    }
    else
    {
        printf("关键字 %d 未找到\n", k);
    }
    return 0;
}

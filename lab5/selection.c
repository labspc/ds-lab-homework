#include <stdio.h>

void selection_sort(int arr[], int n) {
    int i, j, k, temp;

    // 遍历数组
    for (i = 0; i < n - 1; i++) { // 书中 i=1; i<n 
        // 总之就是1～n 最后要走 n-1 趟  n 的值通过 sizeof 来求
        // 设定最小元素的索引k
        k = i;
        
        // 找到剩余数组中最小的元素，就是 i+1 ～ n 下一个都和当前的比较
        for (j = i + 1; j < n; j++) {
            if (arr[j] < arr[k]) {
                k = j; // 保存此时的位置索引
            }
        }

        // 交换找到的最小元素和当前元素
        temp = arr[k];
        arr[k] = arr[i];
        arr[i] = temp;
    }
}

int main() {
    // 关键字序列
    int keywords[] = {21, 25, 49, 25, 16, 8};
    int n = sizeof(keywords) / sizeof(keywords[0]);

    // 直接选择排序
    selection_sort(keywords, n);

    // 打印排序后的结果
    printf("直接选择排序结果:");
    for (int i = 0; i < n; i++) {
        printf(" %d", keywords[i]);
    }
    printf("\n");

    return 0;
}

#include <stdio.h>
#include <stdlib.h>

// 合并两个有序数组
void merge(int *A, int sizeA, int *B, int sizeB, int *C) {
    // 初始化三个索引用于遍历数组A、B和C，k用于填充C
    int i = 0, j = 0, k = 0;
    while (i < sizeA && j < sizeB) {
        if (A[i] < B[j]) {
            C[k++] = A[i++];
        } else {
            C[k++] = B[j++];
        }
    }
    while (i < sizeA) {
        C[k++] = A[i++];
    }
    while (j < sizeB) {
        C[k++] = B[j++];
    }
}

// 接口函数，供外部调用
void merge_api(int *A, int sizeA, int *B, int sizeB, int *C) {
    merge(A, sizeA, B, sizeB, C);
}

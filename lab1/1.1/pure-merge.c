#include <stdio.h>
#include <stdlib.h>

// 合并两个有序数组
void merge(int *A, int sizeA, int *B, int sizeB, int *C) {
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

int main() {
    int A[] = {1, 3, 5, 7};
    int B[] = {2, 4, 6, 8};
    int sizeA = sizeof(A) / sizeof(A[0]);
    int sizeB = sizeof(B) / sizeof(B[0]);
    int sizeC = sizeA + sizeB;
    int *C = (int *)malloc(sizeC * sizeof(int));

    if (C == NULL) {
        printf("内存分配失败\n");
        return 1;
    }

    merge(A, sizeA, B, sizeB, C);

    printf("合并后的数组: ");
    for (int i = 0; i < sizeC; i++) {
        printf("%d ", C[i]);
    }
    printf("\n");

    free(C);
    return 0;
}

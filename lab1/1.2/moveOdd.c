#include <stdio.h>

void swap(int *a, int *b) {
    int temp = *a;
    *a = *b;
    *b = temp;
}

void moveOdd(int A[], int size) {
    int left = 0, right = size - 1; // right要空出末尾的结束符号
    while (left < right) {
        if (A[left] % 2 == 0) {
            swap(&A[left], &A[right]);
            right--;
        } else {
            left++;
        }
    }
}

void moveOdd_api(int A[], int size) {
    moveOdd(A, size);
}

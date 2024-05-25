#include <iostream>
#include <vector>
#include <algorithm> // 提供 swap 函数、copy 函数等

void moveOdd(std::vector<int>& A) {
    int left = 0;
    int right = A.size() - 1;
    while (left < right) {
        if (A[left] % 2 == 0) {
            std::swap(A[left], A[right]);
            right--;
        } else {
            left++;
        }
    }
}

// 导出 API 函数
extern "C" void moveOdd_api(int* A, int size) {
    std::vector<int> vec(A, A + size);  // 将数组 A 转换为 std::vector（处理前）
    moveOdd(vec);                       
    std::copy(vec.begin(), vec.end(), A); // 将 std::vector 中的元素复制回数组 A（处理后）
}

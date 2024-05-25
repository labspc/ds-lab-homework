#include <stdio.h>
#include <stdlib.h>

// 单链表结点定义
struct ListNode {
    int val;
    struct ListNode* next;
};

// 创建链表
struct ListNode* createList(int size) {
    // 初始化指针
    struct ListNode* head = NULL;
    struct ListNode* tail = NULL;

    // 读取数组大小
    int* arr = (int*)malloc(size * sizeof(int));
    if (!arr) {
        printf("内存分配失败\n");
        exit(1);
    }

    printf("请输入 %d 个整数：\n", size);
    for (int i = 0; i < size; ++i) {
        scanf("%d", &arr[i]);
    }

    // 根据数组创建链表
    for (int i = 0; i < size; ++i) {
        // 遍历创建新的链表结点 newNode，新节点的 val 属性被设置为当前数组元素的值，next 属性被设置为 NULL
        struct ListNode* newNode = (struct ListNode*)malloc(sizeof(struct ListNode));
        newNode->val = arr[i];
        newNode->next = NULL;

        if (!head) {
            head = newNode;
            tail = newNode;
        } else { // 不为空时，将新节点连接到链表尾部
            tail->next = newNode;
            tail = newNode; // 更新 tail 为新的尾巴
        }
    }

    free(arr); // 释放动态分配的数组内存
    return head;
}

// 统计链表中不及格人数
int less60(struct ListNode* head) {
    int count = 0;
    struct ListNode* curr = head;
    while (curr) {
        if (curr->val < 60) { // 判断成绩是否不及格
            count++;
        }
        curr = curr->next;
    }
    return count;
}

int main() {
    int size;
    printf("请输入数组的大小：\n");
    scanf("%d", &size);

    // 创建单链表
    struct ListNode* head = createList(size);

    // 统计不及格人数并输出
    int failCount = less60(head);
    printf("不及格人数：%d\n", failCount);

    return 0;
}

#include <stdio.h>
#include <stdlib.h>

// 单链表结点定义
struct ListNode {
    int val;
    struct ListNode* next;
};

// 创建链表
struct ListNode* createList(int arr[], int size) {
    // 初始化指针
    struct ListNode* head = NULL;
    struct ListNode* tail = NULL;

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
    return head;
}

// 遍历单链表并找到最大值的节点
struct ListNode* findMax(struct ListNode* head) {
    if (!head) {
        return NULL;
    }
    // 查找时临时需要的当前最大和下一个
    struct ListNode* maxNode = head;
    struct ListNode* curr = head->next;

    while (curr) {
        if (curr->val > maxNode->val) { // 现在的比上一个还大，最大的就是当前
            maxNode = curr;
        }
        curr = curr->next; // 向后移动继续遍历
    }

    return maxNode;
}


int main() {
    int arr[8];
    printf("请输入8个整数：\n");
    for (int i = 0; i < 8; ++i) {
        scanf("%d", &arr[i]);
    }

    // 创建单链表
    struct ListNode* head = createList(arr, 8);

    // 遍历单链表找到最大值的节点
    struct ListNode* maxNode = findMax(head);

    if (maxNode) {
        printf("单链表中元素值最大的节点值为：%d\n", maxNode->val);
    } else {
        printf("单链表为空\n");
    }

    return 0;
}

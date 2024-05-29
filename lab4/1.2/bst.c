#include <stdio.h>
#include <stdlib.h>

// 二叉树节点定义
typedef struct Node {
    int data;
    struct Node* left;
    struct Node* right;
} Node;

// 创建新节点
Node* create_node(int data) {
    Node* new_node = (Node*)malloc(sizeof(Node));
    new_node->data = data;
    new_node->left = new_node->right = NULL;
    return new_node;
}

// 插入新节点到二叉排序树
Node* insert(Node* root, int data) {
    if (root == NULL) {
        return create_node(data); // 树空，直接插入
    }
    // 非空，插入值比根节点小则插入左子树，否则插入右子树
    if (data < root->data) {
        root->left = insert(root->left, data);
    } else {
        root->right = insert(root->right, data);
    }
    return root;
}

// 中序遍历 LTR
void inorder_traversal(Node* root) {
    if (root != NULL) {
        inorder_traversal(root->left);
        printf("%d ", root->data);
        inorder_traversal(root->right);
    }
}

// 对外的API函数
Node* initialize_tree(int* keys, int n) {
    Node* root = NULL;
    for (int i = 0; i < n; i++) {
        root = insert(root, keys[i]);
    }
    return root;
}

void inorder_traversal_api(Node* root) {
    inorder_traversal(root);
    printf("\n");
}

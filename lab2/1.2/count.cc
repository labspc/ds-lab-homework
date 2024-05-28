
#include <ios>
#include <iostream>
#include <vector>

using namespace std;

struct TreeNode {
    int val;
    TreeNode *left, *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

// TLR
void preOrder(TreeNode* root) {
    if (!root) return;
    cout << root->val << " "; // print root 
    preOrder(root->left);
    preOrder(root->right);
}

// LTR
void inOrder(TreeNode* root) {
    if (!root) return;
    inOrder(root->left);
    cout << root->val << " ";
    inOrder(root->right);
}

// LRT
void postOrder(TreeNode* root) {
    if (!root) return;
    postOrder(root->left);
    postOrder(root->right);
    cout << root->val << " ";
}

// Function to count the number of leaf nodes in the binary tree
int countLeaves(TreeNode* root) {
    if (!root) return 0;
    if (!root->left && !root->right) return 1;
    return countLeaves(root->left) + countLeaves(root->right);
}

int main() {
    // 创建结点5、8、3
    TreeNode *node5 = new TreeNode(5);
    node5->val = 5;
    node5->left = NULL;
    node5->right = NULL;

    TreeNode *node8 = new TreeNode(8);
    node8->val = 8;
    node8->left = NULL;
    node8->right = NULL;

    TreeNode *node3 = new TreeNode(3);
    node3->val = 3;
    node3->left = node5;
    node3->right = node8;

    // 创建结点2、1
    TreeNode *node2 = new TreeNode(2);
    node2->val = 2;
    node2->left = NULL;
    node2->right = NULL;

    TreeNode *node1 = new TreeNode(1);
    node1->val = 1;
    node1->left = node3;
    node1->right = node2;

    preOrder(node1);
    cout << endl;

    inOrder(node1);
    cout << endl;

    postOrder(node1);
    cout << endl;

    // Count and print the number of leaf nodes
    int leafCount = countLeaves(node1);
    cout << "Number of leaf nodes: " << leafCount << endl;

    return 0;
}

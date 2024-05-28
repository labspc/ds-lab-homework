#include <ios>
#include <iostream>
#include <vector>

using namespace std;

struct TreeNode {
    int val;
    TreeNode *left, *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

// Function to build the tree from the pre-order sequence
TreeNode* buildTree(const vector<int>& preOrderVec, int& index) {
    if (index >= preOrderVec.size() || preOrderVec[index] == -1) {
        index++;
        return NULL; // return NULL if the current node is -1
    }
    // 调用构造函数，新建一个TreeNode对象，
    // 然后传入参数这个对象的val的值是一组vector，主程序会具体化这组vector的内容；
    // 最后将root指针指向这个创建的对象
    TreeNode* root = new TreeNode(preOrderVec[index++]); // call the constructor

    root->left = buildTree(preOrderVec, index);
    root->right = buildTree(preOrderVec, index);
    return root;
}

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

int main() {
    // vector<int> preOrderVec = {1, 2, -1, -1, 3, 5, -1, -1, 8, -1, -1};
    // int index = 0;
    // TreeNode* root = buildTree(preOrderVec, index);

    // cout << "Pre-order: ";
    // preOrder(root);
    // cout << endl;

    // cout << "In-order: ";
    // inOrder(root);
    // cout << endl;

    // cout << "Post-order: ";
    // postOrder(root);
    // cout << endl;

    /* 练习二叉树的写法 ****************/

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
    node3->left = NULL;
    node3->right = NULL;

    // 建立联系
    node3->left = node5;
    node3->right = node8;

    // 创建结点2、1
    TreeNode *node2 = new TreeNode(2);
    node2->val = 2;
    node2->left = NULL;
    node2->right = NULL;

    TreeNode *node1 = new TreeNode(1);
    node1->val = 1;
    node1->left = NULL;
    node1->right = NULL;

    // 建立联系
    node1->left = node3;
    node1->right = node2;

    preOrder(node1);
    cout << endl;

    inOrder(node1);
    cout << endl;

    postOrder(node1);
    cout << endl;

    return 0;
}

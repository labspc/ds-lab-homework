#include <iostream>
using namespace std;

// 定义边表结构
struct AdjListNode {
    int vertex;
    struct AdjListNode* next;
};

struct Graph {
    int numVertices; // 图中顶点的数量
    struct AdjListNode** adjLists; // 由表头结点构成的数组
    bool* visited; // 访问标记数组
};

// 创建新的邻接表节点
AdjListNode* createNode(int v) {
    AdjListNode* newNode = new AdjListNode;
    newNode->vertex = v;
    newNode->next = nullptr;
    return newNode;
}

// 创建图
Graph* createGraph(int vertices) {
    Graph* graph = new Graph;
    graph->numVertices = vertices;

    // 创建邻接表
    graph->adjLists = new AdjListNode*[vertices]; // 赋值给 adjLists 是一串地址，用于存储邻接链表
    graph->visited = new bool[vertices];
    // 初始化邻接表和访问标记数组
    for (int i = 0; i < vertices; i++) {
        graph->adjLists[i] = nullptr;
        graph->visited[i] = false;
    }
    return graph;
}

// 添加边
void addEdge(Graph* graph, int src, int dest) {
    // 添加从 src 到 dest 的边
    AdjListNode* newNode = createNode(dest);
    newNode->next = graph->adjLists[src];
    graph->adjLists[src] = newNode;

    // 添加从 dest 到 src 的边（因为是无向图）
    newNode = createNode(src);
    newNode->next = graph->adjLists[dest];
    graph->adjLists[dest] = newNode;
}



// 深度优先搜索函数（仿照二叉树先根遍历）
void DFS(Graph* graph, int vertex) {
    // 访问当前节点
    graph->visited[vertex] = true; // 标记当前顶点为已访问
    cout << vertex << " "; // 输出访问过的顶点

    // 获取当前顶点的邻接表头节点
    AdjListNode* adjList = graph->adjLists[vertex];
    while (adjList != nullptr) {
        int connectedVertex = adjList->vertex; // 从当前邻接表头结点处获取链表的顶点编号
        if (!graph->visited[connectedVertex]) {
            DFS(graph, connectedVertex);
        }
        adjList = adjList->next;
    }
}


int main() {
    int numVertices, numEdges;
    cout << "Enter the number of vertices: ";
    cin >> numVertices;
    cout << "Enter the number of edges: ";
    cin >> numEdges;

    Graph* graph = createGraph(numVertices);

    cout << "Enter edges (start vertex and end vertex):" << endl;
    for (int i = 0; i < numEdges; i++) {
        int u, v;
        cin >> u >> v;
        addEdge(graph, u, v);
    }


    int startVertex;
    cout << "Enter the start vertex for DFS: ";
    cin >> startVertex;

    cout << "DFS starting from vertex " << startVertex << ": ";
    DFS(graph, startVertex);
    cout << endl;

    // 释放内存
    for (int i = 0; i < numVertices; i++) {
        AdjListNode* temp = graph->adjLists[i];
        while (temp) {
            AdjListNode* toDelete = temp;
            temp = temp->next;
            delete toDelete;
        }
    }
    delete[] graph->adjLists;
    delete[] graph->visited;
    delete graph;

    return 0;
}

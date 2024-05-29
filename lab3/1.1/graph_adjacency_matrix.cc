#include <iostream>
#include <vector>
#include <iomanip>

using namespace std;

const int INF = 99999; // 定义一个表示无穷大的常量（用于表示顶点间无边）

// 图的结构定义
struct Graph {
    int numVertices; // 顶点数
    int numEdges;    // 边数
    vector<vector<int>> adjMatrix; // 邻接矩阵

    // 构造函数
    Graph(int vertices, int edges) {
        numVertices = vertices;
        numEdges = edges;
        // 初始化邻接矩阵为全为无穷大，resize()用来调整矩阵大小
        adjMatrix.resize(vertices, vector<int>(vertices, INF)); // 初始化邻接矩阵
    }

    // 添加边的方法
    void addEdge(int u, int v, int weight) {
        adjMatrix[u][v] = weight;
        // 如果是无向图，取消注释下面一行
        // adjMatrix[v][u] = weight;
    }

    // 输出邻接矩阵的方法
    // const不会修改类的成员变量，因此可以在const成员函数中访问类的成员变量
    void printMatrix() const {
        for (int i = 0; i < numVertices; ++i) { // 遍历每一行
            for (int j = 0; j < numVertices; ++j) { // 遍历每一列
                if (adjMatrix[i][j] == INF) {
                    cout << setw(5) << "INF"; // 输出宽度为5，使得矩阵美观
                } else {
                    cout << setw(5) << adjMatrix[i][j];
                }
            }
            cout << endl; // 换行
        }
    }

    // 计算并输出每个顶点的度的方法（有向图出度）
    void calculateDegrees() const {
        for (int i = 0; i < numVertices; ++i) {
            int degree = 0;
            for (int j = 0; j < numVertices; ++j) {
                if (adjMatrix[i][j] != 0 && adjMatrix[i][j] != INF) { // 不为0和不为INF时
                    degree++;
                }
            }
            cout << "Vertex " << i << " degree: " << degree << endl;
        }
    }
};

int main() {
    int n, e;
    cout << "Enter the number of vertices: ";
    cin >> n;
    cout << "Enter the number of edges: ";
    cin >> e;

    // 初始化图
    Graph graph(n, e);

    // 输入边的信息并填充邻接矩阵
    cout << "Enter edges (start vertex, end vertex, weight):" << endl;
    for (int i = 0; i < e; ++i) {
        int u, v, w;
        cin >> u >> v >> w;
        graph.addEdge(u, v, w);
    }

    // 分行输出邻接矩阵
    cout << "Adjacency Matrix:" << endl;
    graph.printMatrix();

    // 求出各顶点的度并输出
    cout << "Degrees of each vertex:" << endl;
    graph.calculateDegrees();

    return 0;
}

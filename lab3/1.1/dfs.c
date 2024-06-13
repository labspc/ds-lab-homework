#include <stdio.h>
#include <stdlib.h>

#define MAX 100 // 最大顶点数
#define INF 99999 // 定义无穷大

// 邻接矩阵
int adjMatrix[MAX][MAX];
int visited[MAX]; // 访问标记数组

// 深度优先搜索函数
void DFS(int vertex, int numVertices) {
    visited[vertex] = 1; // 标记当前顶点为已访问
    printf("Visited %d\n", vertex); // 输出访问过的顶点

    // 遍历所有顶点，检查是否有边并且未访问
    for (int i = 0; i < numVertices; ++i) {
        // 1.从当前顶点到这个顶点是否有边
        // 2.这个顶点是否已经访问过 visited[i] == 0
        if (adjMatrix[vertex][i] != 0 && adjMatrix[vertex][i] != INF && !visited[i]) {
            DFS(i, numVertices); // 递归调用 DFS
        }
    }
}


// 主函数
int main() {
    int numVertices, numEdges;

    // 输入顶点数和边数
    printf("Enter the number of vertices: ");
    scanf("%d", &numVertices);
    printf("Enter the number of edges: ");
    scanf("%d", &numEdges);

    // 初始化邻接矩阵和访问标记数组
    for (int i = 0; i < numVertices; ++i) {
        for (int j = 0; j < numVertices; ++j) {
            adjMatrix[i][j] = (i == j) ? 0 : INF; // 对角线元素为0，其它为无穷大
        }
        visited[i] = 0; // 初始化所有顶点为未访问
    }

    // 输入边的信息并填充邻接矩阵
    printf("Enter edges (start vertex, end vertex):\n");
    for (int i = 0; i < numEdges; ++i) {
        int u, v, w;
        printf("Edge %d: ", i+1);
        scanf("%d %d", &u, &v);
        adjMatrix[u][v] = 1;
        // 如果是无向图，添加以下一行代码
        adjMatrix[v][u] = 1;
    }

    // 输入起始顶点并调用 DFS
    int startVertex;
    printf("Enter the start vertex: ");
    scanf("%d", &startVertex);

    printf("DFS starting from vertex %d:\n", startVertex);
    DFS(startVertex, numVertices);

    return 0;
}

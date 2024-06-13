// 最早开始EST == 最晚开始LST
// 结点Node表示事件Event，边Edge表示活动Activity

#include <stdio.h>
#include <limits.h>

#define MAX_NODES 100

typedef struct {
    int v, w;
} Edge;

typedef struct {
    Edge edges[MAX_NODES];
    int edge_count;
} Node;

Node graph[MAX_NODES];
int in_degree[MAX_NODES];
int earliest[MAX_NODES];
int latest[MAX_NODES];
int top_order[MAX_NODES];
int queue[MAX_NODES];
int front = 0, rear = 0;

void add_edge(int u, int v, int w) {
    graph[u].edges[graph[u].edge_count].v = v;
    graph[u].edges[graph[u].edge_count].w = w;
    graph[u].edge_count++;
    in_degree[v]++;
}

// 顶点排序
void topological_sort(int num_nodes) {
    for (int i = 0; i < num_nodes; i++) {
        if (in_degree[i] == 0) {
            queue[rear++] = i;
        }
    }

    int idx = 0;
    while (front != rear) {
        int u = queue[front++];
        top_order[idx++] = u;

        for (int i = 0; i < graph[u].edge_count; i++) {
            int v = graph[u].edges[i].v;
            if (--in_degree[v] == 0) {
                queue[rear++] = v;
            }
        }
    }
}

// 关键路径
void critical_path_aoe(int num_nodes) {
    topological_sort(num_nodes);

    for (int i = 0; i < num_nodes; i++) {
        earliest[i] = 0;
    }

    for (int i = 0; i < num_nodes; i++) {
        int u = top_order[i];
        for (int j = 0; j < graph[u].edge_count; j++) {
            int v = graph[u].edges[j].v;
            int w = graph[u].edges[j].w;
            if (earliest[v] < earliest[u] + w) {
                earliest[v] = earliest[u] + w;
            }
        }
    }
    int end_time = earliest[top_order[num_nodes - 1]];
    
    for (int i = 0; i < num_nodes; i++) {
        latest[i] = end_time;
    }

    for (int i = num_nodes - 1; i >= 0; i--) {
        int u = top_order[i];
        for (int j = 0; j < graph[u].edge_count; j++) {
            int v = graph[u].edges[j].v;
            int w = graph[u].edges[j].w;
            if (latest[u] > latest[v] - w) {
                latest[u] = latest[v] - w;
            }
        }
    }

    printf("Critical Path:\n");
    for (int u = 0; u < num_nodes; u++) {
        for (int i = 0; i < graph[u].edge_count; i++) {
            int v = graph[u].edges[i].v;
            if (earliest[u] == latest[u] && earliest[v] == latest[v]) {
                printf("%d -> %d\n", u, v);
            }
        }
    }
    printf("Project Duration: %d\n", end_time);
}

int main() {
    int num_nodes = 5;
    add_edge(0, 1, 3);
    add_edge(0, 2, 2);
    add_edge(1, 3, 2);
    add_edge(2, 3, 4);
    add_edge(3, 4, 3);

    critical_path_aoe(num_nodes);

    return 0;
}

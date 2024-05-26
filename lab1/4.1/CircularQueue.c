#include <stdio.h>
#include <stdlib.h>

#define MAXSIZE 100

typedef struct {
    int data[MAXSIZE];
    int front;
    int rear;
    int num; // 队列中元素个数
} CircularQueue;

void initQueue(CircularQueue* q);
int isFull(CircularQueue* q);
int isEmpty(CircularQueue* q);
void enqueue(CircularQueue* q, int element);
int dequeue(CircularQueue* q);


void initQueue(CircularQueue* q) {
    q->front = 0;
    q->rear = 0;
    q->num = 0;
}

int isFull(CircularQueue* q) {
    return q->num == MAXSIZE;
}

int isEmpty(CircularQueue* q) {
    return q->num == 0;
}



void enqueue(CircularQueue* q, int element) {
    if (isFull(q)) {
        printf("Queue is full\n");
        return;
    }
    // 规定入队就是这样的（rear最后会指向当前的下一个）
    q->data[q->rear] = element;
    q->rear = (q->rear + 1) % MAXSIZE;
    q->num++;
}

int dequeue(CircularQueue* q) {
    if (isEmpty(q)) {
        printf("Queue is empty\n");
        return -1;
    }
    // 规定出队就是这样的（front指向目前的）
    int element = q->data[q->front];
    q->front = (q->front + 1) % MAXSIZE;
    q->num--;
    return element;
}



// 处理输入的判断函数
void processInput(CircularQueue *q) {
    int value; // 输入的值
    while (1) {

        printf("Enter an integer (0 to stop): ");
        if (scanf("%d", &value) != 1) {
            // 注意 scanf 的返回值
            printf("Invalid input. Please enter a valid integer.\n");
            // scanf 的用法，清空输入缓冲区
            while (getchar() != '\n');
            continue;
        }

        if (value == 0) {
            break;
        } else if (value < 0) {
            enqueue(q, value);
        } else {
            dequeue(q);
        }
    }
}

int main() {
    CircularQueue q;
    initQueue(&q);
    processInput(&q);
    return 0;
}

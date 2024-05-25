#include <stdio.h>
#include <stdlib.h>

/* 栈模板 *************************************************/

#define MAXSIZE 100

typedef struct {
    int top;
    char data[MAXSIZE];
} Stack;

// 初始化栈
void InitStack(Stack *stack)
{
    // 栈空，栈顶指针
    stack->top = -1;
}

// 判断栈是否空
int IsEmpty(Stack *stack)
{
    return stack->top == -1;
}

// 判断栈是否满
int IsFull(Stack *stack)
{
    return stack->top == MAXSIZE - 1; // top 初始化为 -1
}
// 入栈
void Push(Stack *stack, int val)
{
    // 栈满无法入栈
    if (IsFull(stack))
    {
        fprintf(stderr, "Stack overflow 满\n");
        exit(1);
    }

    stack->top++;
    stack->data[stack->top] = val; // 此时索引 stack->top 处的值是 val
}

// 出栈Top指针向下移动
int Pop(Stack *stack)
{
    // 栈空无法弹出
    if (IsEmpty(stack))
    {
        fprintf(stderr, "Stack underflow 空\n");
        exit(1);
        // return 0;
    }
    // 弹出之后要恢复栈顶指针就是要向下移动
    return stack->data[stack->top--];
}

/* 转换函数 *************************************************/
void decimal(int number, int base, char* result) {
    char digits[] = "0123456789ABCDEF";
    // 创建一个栈
    Stack stack;
    InitStack(&stack);

    if (number == 0) {
        Push(&stack, '0');
    } else {
        while (number > 0) {
            int remainder = number % base; // 余
            Push(&stack, digits[remainder]);
            number = number / base; // 商
        }
    }
    // 取出
    int i = 0;
    while (!IsEmpty(&stack)) {  // 判断创建的栈不为空
        result[i++] = Pop(&stack);
    }
    result[i] = '\0'; // 相当于一个 flag
}

int main() {

    int number; printf("请输入一个非负的十进制整数: "); scanf("%d", &number);

    char oct[MAXSIZE];
    char hex[MAXSIZE];

    // 转换为八进制
    decimal(number, 8, oct);
    printf("八进制: %s\n", oct);

    // 转换为十六进制
    decimal(number, 16, hex);
    printf("十六进制: %s\n", hex);

    return 0;
}
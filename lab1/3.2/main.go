package main

// #include <stdlib.h>
import "C"

type cBool C.int

/* 栈模板 *************************************************************/
type Node struct {
	data byte
	next *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) Push(value byte) {
	newNode := &Node{data: value, next: s.top}
	s.top = newNode
}

func (s *Stack) Pop() byte {
	if s.top == nil {
		return 0
	}
	value := s.top.data
	s.top = s.top.next
	return value
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

/* 回文算法 *************************************************************/
//export HalfRead
func HalfRead(s *C.char) cBool {
	str := C.GoString(s)
	stack := &Stack{}
	length := len(str)
	
	// 一进来不用管，直接拿到商数，将前半部分字符压入栈中
	mid := length / 2
	for i := 0; i < mid; i++ {
		stack.Push(str[i])
	}


	// 如果长度为奇数，mid++ 跳过中间字符；偶数直接跳过，进入后面的 for 遍历比较
	if length%2 != 0 {
		mid++
	}

	// 比较后半部分字符和栈顶字符
	for i := mid; i < length; i++ {
		if stack.Pop() != str[i] {
			return cBool(C.int(0)) // false
		}
	}
	return cBool(C.int(1)) // true
}

func main() {}

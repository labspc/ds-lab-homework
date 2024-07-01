package main

import (
	"fmt"
)

// 旅客节点结构体
type Passenger struct {
	name     string
	position int
	next     *Passenger
}

// 单循环链表结构体
type CircularLink struct {
	head *Passenger
	size int
	tail *Passenger
}

// 添加旅客节点方法
// Add 方法用于向循环链表中添加一个新的节点。
// 参数 name 表示新节点的名字，position 表示新节点的位置。
// 方法会根据链表当前是否为空来决定是创建一个新的循环链表还是在已有链表的末尾添加新节点。
func (cl *CircularLink) Add(name string, position int) {
	// 创建一个新的节点，带有名字和位置信息。
	newNode := &Passenger{name: name, position: position}

	// 如果链表当前为空，即头节点为空，则将新节点设置为头节点，并且自己指向自己形成一个循环。
	if cl.head == nil {
		cl.head = newNode
		newNode.next = newNode
	} else {
		// 如果链表不为空，找到当前链表的最后一个节点。
		temp := cl.head
		// 循环直到找到最后一个节点，即下一个节点指向头节点。
		for temp.next != cl.head {
			temp = temp.next
		}
		// 将最后一个节点的下一个节点指向新节点，然后新节点的下一个节点指向头节点，形成新的循环。
		temp.next = newNode
		newNode.next = cl.head
	}
	// 增加链表的节点数量。
	cl.size++
}

// 移除旅客节点方法
// Remove 方法从循环链表中移除指定的乘客节点。
// 如果链表或节点为空，则直接返回，不做任何操作。
// 参数:
//
//	node - 需要被移除的乘客节点。
func (cl *CircularLink) Remove(node *Passenger) {
	if cl.head == nil || node == nil {
		return
	}
	// 从头节点开始遍历链表，寻找待移除的节点。
	current := cl.head
	// 用于记录当前节点的前一个节点，以便后续移除操作。
	var previous *Passenger
	// 循环直到找到待移除的节点。
	for current.next != node {
		// 更新previous指向当前节点。
		previous = current
		// 移动到下一个节点。
		current = current.next
	}

	// 如果找到了待移除的节点。
	if current == node {
		// 如果待移除节点不是头节点。
		if previous != nil {
			// 将前一个节点的next指向当前节点的next，从而跳过待移除节点。
			previous.next = node.next
		} else {
			// 如果待移除节点是头节点，则更新头节点为下一个节点。
			cl.head = node.next
		}

		// 移除循环链表中指定的节点
		// 如果待移除节点同时也是头节点，则需要特殊处理，以保持循环链表的连续性。
		if node == cl.head {
			// 使用临时变量保存当前头节点
			temp := cl.head
			// 遍历链表直到找到待移除节点的前一个节点
			for temp.next != node {
				temp = temp.next
			}
			// 重新连接链表，绕过待移除的头节点
			// 重新连接链表，绕过原头节点。
			temp.next = cl.head.next
			// 更新头节点为下一个节点
			cl.head = cl.head.next
		}

		// 如果待移除节点同时也是尾节点，则需要特殊处理，以保持循环链表的连续性。
		if node == cl.tail {
			// 使用临时变量遍历至尾节点的前一个节点
			temp := cl.head
			for temp.next != cl.tail {
				temp = temp.next
			}
			// 将尾节点的前一个节点的next指针指向头节点，实现循环
			temp.next = cl.head
			// 更新尾节点为新的最后一个元素
			cl.tail = temp
		}
	}
	// 移除节点后，更新链表大小。
	cl.size--
}

// 获取链表大小方法
func (cl *CircularLink) Size() int {
	return cl.size
}

// 按指定步长遍历链表方法
// Travel 在循环链表中从指定节点开始移动指定步数，并返回移动后的节点。
// @param startNode 起始节点。
// @param step 移动的步数。
// @return 返回移动后的节点。
func (cl *CircularLink) Travel(startNode *Passenger, step int) *Passenger {
	// 如果链表头节点为空，则返回nil，表示链表为空。
	if cl.head == nil {
		return nil
	}
	// 将当前节点设置为起始节点。
	current := startNode
	// 从起始节点开始，向前移动指定步数。
	for i := 1; i < step; i++ {
		// 每次移动到下一个节点。
		current = current.next
	}
	// 返回移动后的节点。
	return current //= current.next
}

// JProblem 结构体
type JProblem struct {
	list      *CircularLink
	interval  int // M
	totalSize int // N
}

// 初始化链表方法
func (jp *JProblem) Initialize(names []string) {
	jp.list = &CircularLink{}
	jp.totalSize = len(names)
	for i, name := range names {
		jp.list.Add(name, i+1) // 旅客从1开始报数
	}
}

// 执行问题规则方法
// Exe 方法用于解决特定问题，即根据设定的条件从乘客列表中移除一部分乘客。
// 它返回被移除的乘客列表。
func (jp *JProblem) Exe() []Passenger {
	// 初始化一个切片用于存储被移除的乘客。
	var removedPg []Passenger
	// 从乘客列表的头部开始遍历。
	current := jp.list.head
	// 当列表中的乘客数量超过总容量的一半时，执行移除操作。
	for jp.list.Size() > jp.totalSize/2 {
		// 根据设定的间隔前进到下一个乘客。
		current = jp.list.Travel(current, jp.interval)
		// 将当前乘客添加到被移除的乘客列表中。
		removedPg = append(removedPg, *current)
		// 从乘客列表中移除当前乘客。
		jp.list.Remove(current)
		// 继续遍历下一个乘客。
		current = current.next
	}
	// 返回被移除的乘客列表。
	return removedPg
}

// 主函数
func main() {
	//names := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O"}
	//names := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	//names := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	names := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	//names := []string{"A", "B", "C", "D"}
	//names := []string{"A", "B", "C", "D", "E"}

	// 给定 M 值
	interval := 9
	jp := &JProblem{interval: interval}

	// 给定 N 值
	jp.Initialize(names)

	removedPg := jp.Exe()

	fmt.Println("被投入大海的旅客:")
	for _, passenger := range removedPg {
		fmt.Printf("姓名: %s, 位置: %d\n", passenger.name, passenger.position)
	}
}

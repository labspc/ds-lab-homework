package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
)

// Node 结构体，表示景点信息
type Node struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Info string `json:"info"`
}

// Graph结构体，表示图
type Graph struct {
	Nodes     map[string]Node // 将图中节点名称映射到Node结构体
	AdjMatrix [][]float64     // 邻接矩阵
}

// 初始化图
// NewGraph 初始化一个图结构。
// 参数size指定了图中节点的数量，该图将被初始化为无向完全图，即任意两个节点之间都有边。
// 这个方法主要负责分配存储节点和邻接矩阵的空间，并将邻接矩阵中的所有元素初始化为正无穷大，
// 表示初始时所有边的权重都是无限大（尚未定义）。
func (g *Graph) NewGraph(size int) {
	// 初始化节点映射，用于存储图中的所有节点。
	g.Nodes = make(map[string]Node)

	// 初始化邻接矩阵，大小为size×size，用于存储图中节点之间的边权重。
	g.AdjMatrix = make([][]float64, size)
	for i := range g.AdjMatrix {
		// 每一行都初始化为长度为size的切片，用于存储与当前节点直接相连的其他节点的边权重。
		g.AdjMatrix[i] = make([]float64, size)
		for j := range g.AdjMatrix[i] {
			// 将所有边的初始权重设置为正无穷大，表示尚未定义。
			g.AdjMatrix[i][j] = math.Inf(1) // +∞
		}
	}
}

// 添加景点
func (g *Graph) AddNode(code, name, info string) {
	g.Nodes[code] = Node{Code: code, Name: name, Info: info} // code 作为键，Node结构体作为值
}

// 添加边
// AddEdge 向图中添加一条边。
// 该方法通过指定边的起始顶点和结束顶点以及边的权重来添加边。
// 参数 from 表示边的起始顶点，to 表示边的结束顶点，weight 表示边的权重。
// 该方法适用于无向图，因为对于无向图中的每条边，从顶点A到顶点B的边和从顶点B到顶点A的边是等价的。
// 因此，方法中同时更新了两个方向的边的权重，以确保图的对称性。
func (g *Graph) AddEdge(from, to string, weight float64) {
	// 计算起始顶点和结束顶点在邻接矩阵中的索引。
	// 这里假设顶点的名称是大写字母A到Z，索引从0到25对应。
	fromIndex := int(from[0] - 'A')
	toIndex := int(to[0] - 'A')

	// 在邻接矩阵中添加边，同时考虑到无向图的特性，需要在两个方向上都添加边。
	g.AdjMatrix[fromIndex][toIndex] = weight
	g.AdjMatrix[toIndex][fromIndex] = weight
}

// 格式化输出为JSON的方法
func Format(data any) string {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("JSON编码错误: %s", err)
	}
	return string(jsonData)
}

// // DFS遍历
// func (g *Graph) DFS(start string) []string {
// 	startIndex := int(start[0] - 'A')
// 	visited := make([]bool, len(g.AdjMatrix))
// 	var result []string
// 	g.dfsHelper(startIndex, visited, &result)
// 	return result
// }

//	func (g *Graph) dfsHelper(v int, visited []bool, result *[]string) {
//		visited[v] = true
//		*result = append(*result, string('A'+v))
//		for i, val := range g.AdjMatrix[v] {
//			if val != math.Inf(1) && !visited[i] {
//				g.dfsHelper(i, visited, result)
//			}
//		}
//	}
// func (g *Graph) DFS(start string) []string {

// 	dfs := func(v int, visited []bool, result *[]string) {
// 		// 标记当前节点为已访问。
// 		visited[v] = true
// 		// 将当前节点添加到遍历结果中。
// 		*result = append(*result, string('A'+v))
// 		// 遍历当前节点的邻接矩阵，寻找未访问的邻接节点。
// 		for i, val := range g.AdjMatrix[v] {
// 			// 如果邻接节点未被访问且存在连接（非无穷大），则递归访问该邻接节点。
// 			if val != math.Inf(1) && !visited[i] {
// 				dfs(i, visited, result)
// 			}
// 		}
// 	}

// 	startIndex := int(start[0] - 'A')
// 	visited := make([]bool, len(g.AdjMatrix))
// 	result := make([]string, 0)
// 	dfs(startIndex, visited, &result)

//		return result
//	}
//
// // DFS遍历
// func (g *Graph) DFS(start string) []string {
// 	startIndex := int(start[0] - 'A')
// 	visited := make([]bool, len(g.AdjMatrix))
// 	var result []string

// 	dfs := func(v int) {
// 		visited[v] = true
// 		result = append(result, string('A'+v))
// 		for i, val := range g.AdjMatrix[v] {
// 			if val != math.Inf(1) && !visited[i] {
// 				dfs(i)
// 			}
// 		}
// 	}

// 	dfs(startIndex)
// 	return result
// }

// DFS遍历
func (g *Graph) DFS(start string) []string {

	var dfs func(int, []bool, *[]string)
	dfs = func(v int, visited []bool, result *[]string) {
		visited[v] = true
		*result = append(*result, string('A'+v))
		for i, val := range g.AdjMatrix[v] {
			if val != math.Inf(1) && !visited[i] {
				dfs(i, visited, result)
			}
		}
	}
	startIndex := int(start[0] - 'A')
	visited := make([]bool, len(g.AdjMatrix))
	result := make([]string, 0)
	dfs(startIndex, visited, &result)

	return result
}

// BFS遍历
func (g *Graph) BFS(start string) []string {
	startIndex := int(start[0] - 'A')
	visited := make([]bool, len(g.AdjMatrix))
	var queue []int
	var result []string

	queue = append(queue, startIndex)
	visited[startIndex] = true

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		result = append(result, string('A'+v))

		for i, val := range g.AdjMatrix[v] {
			if val != math.Inf(1) && !visited[i] {
				queue = append(queue, i)
				visited[i] = true
			}
		}
	}

	return result
}

// Dijkstra算法求最短路径
func (g *Graph) Dijkstra(start, target string) (float64, []string) {
	size := len(g.AdjMatrix)
	dist := make([]float64, size)
	prev := make([]int, size)
	visited := make([]bool, size)

	// 初始化距离和前驱
	for i := range dist {
		dist[i] = math.Inf(1) // +∞
		prev[i] = -1          // 为空
	}
	startIndex := int(start[0] - 'A')
	targetIndex := int(target[0] - 'A')
	dist[startIndex] = 0

	for {
		// 初始化最小距离为正无穷大，用于比较当前节点到源点的最小距离
		// 找到距离起点最近的未处理节点
		u := -1
		minDist := math.Inf(1) // +∞

		// 遍历所有节点，寻找未访问且距离源点最近的节点
		for i := 0; i < size; i++ {
			// 如果节点未被访问且距离小于当前最小距离，则更新最小距离和对应的节点
			if !visited[i] && dist[i] < minDist {
				minDist = dist[i]
				u = i
			}
		}

		// 如果所有节点都被访问或找到目标节点，则结束循环
		if u == -1 || u == targetIndex {
			break
		}

		// 标记当前节点为已访问
		visited[u] = true

		// 遍历当前节点的所有邻接节点
		// 更新邻接节点的距离
		for v := 0; v < size; v++ {
			// 如果当前节点到邻接节点的边存在且未被访问
			if g.AdjMatrix[u][v] != math.Inf(1) && !visited[v] {
				// 计算从源点经当前节点到邻接节点的新的距离
				alt := dist[u] + g.AdjMatrix[u][v]
				// 如果新的距离小于当前记录的最短距离，则更新最短距离和前驱节点
				if alt < dist[v] {
					dist[v] = alt
					prev[v] = u
				}
			}
		}
	}

	// 根据给定的目标索引和前置节点数组，计算从起点到目标节点的最短距离和路径。
	// targetIndex 是目标节点的索引，prev 是一个数组，记录了每个节点的前置节点。
	// 返回值是目标节点的距离和从起点到目标节点的路径。
	path := []string{}
	for currentNode := targetIndex; currentNode != -1; currentNode = prev[currentNode] {
		// 将当前节点添加到路径的起始位置。
		// 这里通过字符运算将节点索引转换为字母，以构建路径字符串。
		path = append([]string{string('A' + currentNode)}, path...)
	}

	// 返回目标节点的距离和路径。
	return dist[targetIndex], path
}

// CLI交互函数
func CLI(g *Graph) {
	for {
		fmt.Println("\n选择操作:")
		fmt.Println("1: 查询景点信息")
		fmt.Println("2: 查询景点之间的最短路径")
		fmt.Println("3: DFS遍历")
		fmt.Println("4: BFS遍历")
		fmt.Println("5: 退出")
		fmt.Print("输入选项: ")

		var choice int
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Println("无效输入，请输入一个数字")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("输入景点代码: ")
			var code string
			fmt.Scan(&code)
			if node, found := g.Nodes[code]; found {
				fmt.Println("景点信息:", Format(node))
			} else {
				fmt.Println("未找到景点:", code)
			}
		case 2:
			fmt.Print("输入起点和终点代码 (例如: A J): ")
			var start, target string
			fmt.Scan(&start, &target)
			if _, found := g.Nodes[start]; !found {
				fmt.Println("未找到起点景点:", start)
				continue
			}
			if _, found := g.Nodes[target]; !found {
				fmt.Println("未找到终点景点:", target)
				continue
			}
			dist, path := g.Dijkstra(start, target)
			fmt.Printf("从 %s 到 %s 的最短路径距离为: %f\n", start, target, dist)
			fmt.Printf("路径为: %v\n", path)
		case 3:
			fmt.Print("输入起点代码: ")
			var start string
			fmt.Scan(&start)
			if _, found := g.Nodes[start]; !found {
				fmt.Println("未找到景点:", start)
				continue
			}
			dfsResult := g.DFS(start)
			fmt.Println("DFS遍历结果:", dfsResult)
		case 4:
			fmt.Print("输入起点代码: ")
			var start string
			fmt.Scan(&start)
			if _, found := g.Nodes[start]; !found {
				fmt.Println("未找到景点:", start)
				continue
			}
			bfsResult := g.BFS(start)
			fmt.Println("BFS遍历结果:", bfsResult)
		case 5:
			fmt.Println("退出程序")
			os.Exit(0)
		default:
			fmt.Println("无效选项，请重新输入")
		}
	}
}

// 显示景点编号及其名称
func (g *Graph) Display() {
	for code, node := range g.Nodes {
		if code >= "A" && code <= "J" { // 确保只展示A-J的景点
			fmt.Printf("%s: %s\n", code, node.Name)
		}
	}
}

func main() {

	// 初始化图
	g := &Graph{}
	g.NewGraph(10)

	// 添加景点信息（根据表格）
	g.AddNode("A", "景点A:大操场（可以踢足球）", "景点A的简介:可以用来上体育课")
	g.AddNode("B", "景点B:大体育馆", "景点B的简介:可以举行毕业典礼")
	g.AddNode("C", "景点C:航飞楼大飞机", "景点C的简介:可以拍照")
	g.AddNode("D", "景点D:图书馆", "景点D的简介:可以看书")
	g.AddNode("E", "景点E:图书馆门前的大草坪", "景点E的简介:可以拍照")
	g.AddNode("F", "景点F:行政楼", "景点F的简介:可以去交材料")
	g.AddNode("G", "景点G:交通楼", "景点G的简介:上计算机实验课")
	g.AddNode("H", "景点H:图文二阅", "景点H的简介:补办校园卡")
	g.AddNode("I", "景点I:教学楼旁的莲花池", "景点I的简介:可以看鱼儿游泳")
	g.AddNode("J", "景点J:工程实训中心", "景点J的简介:上物理实验课")

	// 按照“景点简化图”添加边
	g.AddEdge("A", "B", 10)
	g.AddEdge("A", "C", 20)
	g.AddEdge("C", "H", 5)
	g.AddEdge("C", "D", 5)
	g.AddEdge("C", "J", 9)
	g.AddEdge("C", "G", 30)
	g.AddEdge("H", "G", 5)
	g.AddEdge("G", "F", 60)
	g.AddEdge("J", "F", 7)
	g.AddEdge("J", "D", 10)
	g.AddEdge("D", "E", 20)
	g.AddEdge("E", "F", 5)
	g.AddEdge("E", "I", 15)
	g.AddEdge("D", "I", 15)
	g.AddEdge("B", "G", 50)

	// 显示所有景点信息供用户参考
	fmt.Println("景点列表：")
	g.Display()

	// 调用CLI交互函数
	CLI(g)
}

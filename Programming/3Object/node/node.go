package node

import (
	"fmt"
)

// Node 是一个自定义的节点
type Node struct {
	Data        string
	Left, Right *Node
}

// SetData 设值
func (node *Node) SetData(data string) {
	if node == nil {
		fmt.Println("Setting value to nil...ignore")
		return
	}
	node.Data = data
}

// Print 打印
func (node Node) Print() {
	fmt.Println(node.Data)
}

// Traverse 遍历所有节点
func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

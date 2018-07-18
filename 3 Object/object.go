package main

import (
	"fmt"
)

// Node 是一个自定义的节点
type Node struct {
	data        string
	left, right *Node
}

func main() {
	//	结构体和方法
	// structInit()
	// structSlice()
	node := createNode("")
	// node.print()
	// nilPointerOperate()
	node.traverse()
}

func (node *Node) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func nilPointerOperate() {
	var nilNode *Node
	nilNode.print()
	nilNode.setData("我会报错哦！")
}

func (node *Node) setData(data string) {
	if node == nil {
		fmt.Println("Setting value to nil...ignore")
		return
	}
	node.data = data
}

func (node Node) print() {
	fmt.Println(node.data)
}

func createNode(data string) *Node {
	return &Node{data: data}
}

func structSlice() {
	nodes := []Node{
		{data: "第一个"},
		{data: "第二个", left: nil, right: nil},
		{},
	}
	fmt.Println(nodes)
}

func structInit() {
	var root Node
	root = Node{data: "root中"}

	root.left = &Node{data: "left"}
	root.left.left = new(Node)

	root.right = &Node{}
	root.right.left = &Node{"right-left", nil, nil}
}

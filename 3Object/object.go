package main

import (
	"fmt"

	"daker.wang/Azen/Go-execise/3Object/node"
)

func main() {
	//	结构体和方法
	structInit()
	structSlice()
	node := createNode("")
	node.Print()
	// nilPointerOperate()
	traverseTest()
}

func traverseTest() {
	root := createNode("root")
	root.Left = &node.Node{Data: "left"}
	root.Left.Right = &node.Node{Data: "left-right"}

	root.Right = &node.Node{Data: "right"}
	root.Right.Left = &node.Node{Data: "right-left"}
	root.Right.Right = &node.Node{Data: "right-right"}

	root.Traverse()
}

func nilPointerOperate() {
	var nilNode *node.Node
	nilNode.Print()
	nilNode.SetData("我会报错哦！")
}

func createNode(data string) *node.Node {
	return &node.Node{Data: data}
}

func structSlice() {
	nodes := []node.Node{
		{Data: "第一个"},
		{Data: "第二个", Left: nil, Right: nil},
		{},
	}
	fmt.Println(nodes)
}

func structInit() {
	var root node.Node
	root = node.Node{Data: "root中"}

	root.Left = &node.Node{Data: "left"}
	root.Left.Left = new(node.Node)

	root.Right = &node.Node{}
	root.Right.Left = &node.Node{"right-left", nil, nil}
}

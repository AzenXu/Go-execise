package something

import (
	"daker.wang/Azen/Go-execise/3Object/node"
)

type myNode struct {
	node *node.Node
}

func (m *myNode) testFunction() {
	if m == nil || m.node == nil {
		return
	}
	m.node.Print()
}
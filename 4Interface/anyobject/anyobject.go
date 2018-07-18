package anyobject

// Queue 可以放任意元素的队列
type Queue []interface{}

// Push 为Queue追加元素
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

// Pop 如果首元素类型为int，则pop出来
func (q *Queue) Pop() int {
	head, ok := (*q)[0].(int)
	if !ok {
		panic("首元素不是int")
	}
	*q = (*q)[1:]
	return head
}

package queue

type queue []int

func (q *queue) pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *queue) push(v int) {
	*q = append(*q, v)
}

func (q *queue) isEmpyt() bool {
	return len(*q) <= 0
}

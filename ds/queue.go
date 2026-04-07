package ds

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() int {
	v := q.items[0]
	q.items = q.items[1:]
	return v
}

func (q *Queue) Print() {
	fmt.Println(q.items)
}

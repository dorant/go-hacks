package queue

import (
	"container/list"
)

func NewString() *String {
	return &String{list.New()}
}

type String struct {
	list *list.List
}

func (q *String) Len() int {
	return q.list.Len()
}

func (q *String) Enqueue(i string) {
	q.list.PushBack(i)
}

func (q *String) Dequeue() string {
	if q.list.Len() == 0 {
		panic(ErrEmptyQueue)
	}
	raw := q.list.Remove(q.list.Front())
	if typed, ok := raw.(string); ok {
		return typed
	}
	panic(ErrInvalidType)
}

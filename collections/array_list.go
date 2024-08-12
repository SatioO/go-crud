package collections

import "fmt"

type ArrayList[T any] struct {
	Items    []T
	len      uint
	capacity uint
}

func New[T any](initialCapacity uint) *ArrayList[T] {
	return &ArrayList[T]{
		Items:    make([]T, 0),
		capacity: initialCapacity,
		len:      0,
	}
}

func (t *ArrayList[T]) Add(item T) {
	t.Items = append(t.Items, item)
	t.len++
}

func (t *ArrayList[T]) Print() {
	for i := 0; i < len(t.Items); i++ {
		fmt.Printf("item: %v\n", t.Items[i])
	}
}

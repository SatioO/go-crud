package concurrent

import (
	"fmt"
	"sync"
)

type FileReader struct {
	ByteSize int
	mu       sync.Mutex
}

func (h *FileReader) ReadBytes() {
	h.mu.Lock()
	h.ByteSize += 1
	h.mu.Unlock()
}

func Map[T any](arr []T, fn func(T, int) T) []T {
	output := make([]T, 0, len(arr))

	for i, v := range arr {
		output = append(output, fn(v, i))
	}

	return output
}

func Filter[T any](arr []T, fn func(T, int) bool) []T {
	output := make([]T, 0, len(arr))

	for i, v := range arr {
		if fn(v, i) {
			output = append(output, v)
		}
	}

	return output
}

func TryMap() {
	items := []int{1, 2, 3, 4, 5}
	output := Map(items, func(item int, i int) int {
		return item << 2
	})

	fmt.Println("output: ", output)
}

func TryFilter() {
	items := []int{1, 2, 3, 4, 5}
	output := Filter(items, func(item int, i int) bool {
		return item%2 == 0
	})

	fmt.Println("output: ", output)
}

func TryRemove[T any](arr []T, index int) []T {
	return append(arr[:index], arr[index+1:]...)
}

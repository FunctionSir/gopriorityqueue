/*
 * @Author: FunctionSir
 * @License: AGPLv3
 * @Date: 2025-04-22 18:08:21
 * @LastEditTime: 2025-04-22 20:17:02
 * @LastEditors: FunctionSir
 * @Description: -
 * @FilePath: /gopriorityqueue/priorityqueue.go
 */

package gopriorityqueue

type PriorityQueue[T any] struct {
	data []T
	Cmp  func(root, leaf *T) bool
}

func swap[T any](a, b *T) {
	tmp := *a
	*a = *b
	*b = tmp
}

func NewPriorityQueue[T any](cmp func(root, leaf *T) bool) PriorityQueue[T] {
	return PriorityQueue[T]{
		data: make([]T, 0),
		Cmp:  cmp,
	}
}

func (h *PriorityQueue[T]) Size() int {
	return len(h.data)
}

func (h *PriorityQueue[T]) Push(x T) {
	h.data = append(h.data, x)
	cur := h.Size() - 1
	for cur >= 1 && !h.Cmp(&h.data[(cur-1)/2], &h.data[cur]) {
		swap(&h.data[cur], &h.data[(cur-1)/2])
		cur = (cur - 1) / 2
	}
}

func (h *PriorityQueue[T]) Pop() T {
	poped := h.data[0]
	swap(&h.data[0], &h.data[h.Size()-1])
	h.data = h.data[:h.Size()-1]
	father := 0
	for father*2+1 < h.Size() {
		child := father*2 + 1
		if child+1 < h.Size() && h.Cmp(&h.data[child+1], &h.data[child]) {
			child++
		}
		if h.Cmp(&h.data[father], &h.data[child]) {
			break
		}
		swap(&h.data[father], &h.data[child])
		father = child
	}
	return poped
}

func (h *PriorityQueue[T]) Top() T {
	return h.data[0]
}

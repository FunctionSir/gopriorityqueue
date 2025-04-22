/*
 * @Author: FunctionSir
 * @License: AGPLv3
 * @Date: 2025-04-22 18:48:45
 * @LastEditTime: 2025-04-22 22:37:28
 * @LastEditors: FunctionSir
 * @Description: -
 * @FilePath: /gopriorityqueue/priorityqueue_test.go
 */

package gopriorityqueue

import (
	"math/rand"
	"slices"
	"sort"
	"testing"
)

const TEST_SCALE int = 10000000

func TestPriorityQueueRandMinHeap(t *testing.T) {
	pq := NewPriorityQueue(func(a, b *int) bool { return *a <= *b })
	testData := make([]int, TEST_SCALE)
	for i := 0; i < TEST_SCALE; i++ {
		testData[i] = rand.Intn(0x3f3f3f3f)
	}
	correct := make([]int, TEST_SCALE)
	copy(correct, testData)
	sort.Ints(correct)
	for i := 0; i < TEST_SCALE; i++ {
		pq.Push(testData[i])
	}
	for i := 0; i < TEST_SCALE; i++ {
		top := pq.Top()
		if top != correct[i] {
			t.Errorf("#%d: Top should be %d, but got %d", i, correct[i], top)
		}
		poped := pq.Pop()
		if poped != correct[i] {
			t.Errorf("#%d: Should pop %d, but got %d", i, correct[i], poped)
		}
	}
}

func TestPriorityQueueRandMaxHeap(t *testing.T) {
	pq := NewPriorityQueue(func(a, b *int) bool { return *a >= *b })
	testData := make([]int, TEST_SCALE)
	for i := 0; i < TEST_SCALE; i++ {
		testData[i] = rand.Intn(0x3f3f3f3f)
	}
	correct := make([]int, TEST_SCALE)
	copy(correct, testData)
	sort.Ints(correct)
	slices.Reverse(correct)
	for i := 0; i < TEST_SCALE; i++ {
		pq.Push(testData[i])
	}
	for i := 0; i < TEST_SCALE; i++ {
		top := pq.Top()
		if top != correct[i] {
			t.Errorf("#%d: Top should be %d, but got %d", i, correct[i], top)
		}
		poped := pq.Pop()
		if poped != correct[i] {
			t.Errorf("#%d: Should pop %d, but got %d", i, correct[i], poped)
		}
	}
}

func TestPriorityQueueRandMinHeapWithNeg(t *testing.T) {
	pq := NewPriorityQueue(func(a, b *int) bool { return *a <= *b })
	testData := make([]int, TEST_SCALE)
	for i := 0; i < TEST_SCALE; i++ {
		testData[i] = rand.Intn(0x3f3f3f3f)
		if i%2 == 0 {
			testData[i] = -testData[i]
		}
	}
	correct := make([]int, TEST_SCALE)
	copy(correct, testData)
	sort.Ints(correct)
	for i := 0; i < TEST_SCALE; i++ {
		pq.Push(testData[i])
	}
	for i := 0; i < TEST_SCALE; i++ {
		top := pq.Top()
		if top != correct[i] {
			t.Errorf("#%d: Top should be %d, but got %d", i, correct[i], top)
		}
		poped := pq.Pop()
		if poped != correct[i] {
			t.Errorf("#%d: Should pop %d, but got %d", i, correct[i], poped)
		}
	}
}

func TestPriorityQueueRandMaxHeapWithNeg(t *testing.T) {
	pq := NewPriorityQueue(func(a, b *int) bool { return *a >= *b })
	testData := make([]int, TEST_SCALE)
	for i := 0; i < TEST_SCALE; i++ {
		testData[i] = rand.Intn(0x3f3f3f3f)
		if i%2 == 0 {
			testData[i] = -testData[i]
		}
	}
	correct := make([]int, TEST_SCALE)
	copy(correct, testData)
	sort.Ints(correct)
	slices.Reverse(correct)
	for i := 0; i < TEST_SCALE; i++ {
		pq.Push(testData[i])
	}
	for i := 0; i < TEST_SCALE; i++ {
		top := pq.Top()
		if top != correct[i] {
			t.Errorf("#%d: Top should be %d, but got %d", i, correct[i], top)
		}
		poped := pq.Pop()
		if poped != correct[i] {
			t.Errorf("#%d: Should pop %d, but got %d", i, correct[i], poped)
		}
	}
}

func TestPriorityQueueOrdered(t *testing.T) {
	pq := NewPriorityQueue(func(a, b *int) bool { return *a <= *b })
	testData := make([]int, TEST_SCALE)
	for i := 0; i < TEST_SCALE; i++ {
		testData[i] = i + 1
	}
	correct := make([]int, TEST_SCALE)
	copy(correct, testData)
	sort.Ints(correct)
	for i := 0; i < TEST_SCALE; i++ {
		pq.Push(testData[i])
	}
	for i := 0; i < TEST_SCALE; i++ {
		top := pq.Top()
		if top != correct[i] {
			t.Errorf("#%d: Top should be %d, but got %d", i, correct[i], top)
		}
		poped := pq.Pop()
		if poped != correct[i] {
			t.Errorf("#%d: Should pop %d, but got %d", i, correct[i], poped)
		}
	}
}

func TestPriorityQueueReversed(t *testing.T) {
	pq := NewPriorityQueue(func(a, b *int) bool { return *a <= *b })
	testData := make([]int, TEST_SCALE)
	for i := 0; i < TEST_SCALE; i++ {
		testData[i] = TEST_SCALE - i
	}
	correct := make([]int, TEST_SCALE)
	copy(correct, testData)
	sort.Ints(correct)
	for i := 0; i < TEST_SCALE; i++ {
		pq.Push(testData[i])
	}
	for i := 0; i < TEST_SCALE; i++ {
		top := pq.Top()
		if top != correct[i] {
			t.Errorf("#%d: Top should be %d, but got %d", i, correct[i], top)
		}
		poped := pq.Pop()
		if poped != correct[i] {
			t.Errorf("#%d: Should pop %d, but got %d", i, correct[i], poped)
		}
	}
}

func TestPriorityQueueEqual(t *testing.T) {
	pq := NewPriorityQueue(func(a, b *int) bool { return *a <= *b })
	testData := make([]int, TEST_SCALE)
	fixedNumber := rand.Intn(0x3f3f3f3f)
	if rand.Intn(2) == 1 {
		fixedNumber = -fixedNumber
	}
	for i := 0; i < TEST_SCALE; i++ {
		testData[i] = fixedNumber
	}
	correct := make([]int, TEST_SCALE)
	copy(correct, testData)
	sort.Ints(correct)
	for i := 0; i < TEST_SCALE; i++ {
		pq.Push(testData[i])
	}
	for i := 0; i < TEST_SCALE; i++ {
		top := pq.Top()
		if top != correct[i] {
			t.Errorf("#%d: Top should be %d, but got %d", i, correct[i], top)
		}
		poped := pq.Pop()
		if poped != correct[i] {
			t.Errorf("#%d: Should pop %d, but got %d", i, correct[i], poped)
		}
	}
}

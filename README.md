<!--
 * @Author: FunctionSir
 * @License: AGPLv3
 * @Date: 2025-04-22 16:22:12
 * @LastEditTime: 2025-04-22 22:36:51
 * @LastEditors: FunctionSir
 * @Description: -
 * @FilePath: /gopriorityqueue/README.md
-->

# gopriorityqueue

Priority queue in Golang.

Well tested!

## Usage

You need Go version 1.8+

### Import

``` go
import gpq "github.com/FunctionSir/gopriorityqueue"
```

### Create New Priority Queue

``` go
// Note: Type in Cmp func determined the type of elements in the priority queue.
priorityQueueA := gpq.NewPriorityQueue(func(a, b *int) bool { return *a <= *b }) // A min heap, roots <= leaves
priorityQueueB := gpq.NewPriorityQueue(func(a, b *int) bool { return *a >= *b }) // A max heap, roots >= leaves
```

### Push (Add an Element)

``` go
somePriorityQueue.Push(someValue)
```

### Use Top (Get the Max / Min)

``` go
somePriorityQueue.Top() // If it's a min heap, will get min value, if it's a max heap, max value.
```

### Pop (Remove a Max / Min)

``` go
poped := somePriorityQueue.Pop() // Poped thing will be returned.
somePriorityQueue.Pop() // Ignore returned value is OK.
```

### Count Elements

``` go
size := somePriorityQueue.Size() // To know how many elements in the priority queue.
```

## About testing

Used 10000000 (1e7) of: random data(with or with no negative numbers), ordered data, reversed ordered data, equal values to test and passed.

package main

import "fmt"

type Priority byte

const (
	// NOTE: modify the String method, if you add new values or change the order

	Low Priority = iota
	Medium
	High
)

func (p Priority) String() string {
	return []string{"Low", "Medium", "High"}[p] // NOTE: order matters
}

type PriorityQueue[T comparable, U any] struct {
	items      map[T][]U
	priorities []T
}

func (pq *PriorityQueue[T, U]) Add(priority T, value U) {
	pq.items[priority] = append(pq.items[priority], value)
}

func (pq *PriorityQueue[T, U]) Next() (U, bool) {
	for i := 0; i < len(pq.priorities); i++ {
		priority := pq.priorities[i]
		items := pq.items[priority]

		if len(items) > 0 {
			next := items[0]
			pq.items[priority] = items[1:]
			return next, true
		}
	}

	var empty U
	return empty, false
}

func (pq *PriorityQueue[T, U]) Clear(priority T) {
	pq.items[priority] = []U{}
}

func (pq *PriorityQueue[T, U]) ClearAll() {
	for k := range pq.items {
		pq.items[k] = []U{}
	}
}

func NewPriorityQueue[T comparable, U any](priorities []T) PriorityQueue[T, U] {
	return PriorityQueue[T, U]{
		items:      make(map[T][]U),
		priorities: priorities,
	}
}

func main() {
	queue := NewPriorityQueue[Priority, string]([]Priority{High, Medium, Low})

	queue.Add(Low, "L-1")
	queue.Add(High, "H-1")

	fmt.Println(queue.Next()) // H-1 true

	queue.Add(Medium, "M-1")
	queue.Add(High, "H-2")
	queue.Add(High, "H-3")
	queue.Add(Medium, "M-2")

	fmt.Println(queue.Next()) // H-2 true
	fmt.Println(queue.Next()) // H-3 true
	fmt.Println(queue.Next()) // M-1 true
	fmt.Println(queue.Next()) // M-2 true
	fmt.Println(queue.Next()) // L-1 true
	fmt.Println(queue.Next()) // false
	fmt.Println(queue.Next()) // false

	// NOTE: no need to clear these items, they are already empty
	queue.Clear(High)
	queue.Clear(Low)
	queue.Clear(Medium)
	fmt.Println(queue)

	queue.Add(Low, "Workout")
	queue.Add(High, "Get a job")
	queue.Add(Medium, "Finish learning TDD")
	queue.Add(High, "Create a personal website")
	queue.Add(Medium, "Finish reading the book")

	for {
		todo, ok := queue.Next()
		if !ok {
			break
		}

		fmt.Println("🚀", todo)
	}

	// NOTE: no need to clear these items, they are already empty
	queue.ClearAll()
	fmt.Println(queue)
}

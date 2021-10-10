package structure

import "sync"

type queue struct {
	data []int
	size int
	lock sync.RWMutex
}

func InitQueue() *queue {
	queue := new(queue)
	queue.data = []int{}
	queue.size = 0
	return queue
}

func (q *queue) GetSize() int {
	return q.size
}

func (q *queue) Push(value int) bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.data = append([]int{value}, q.data...)
	q.size++
	return true
}

func (q *queue) Pop() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.size == 0 {
		return -1
	}
	queueLastValue := q.data[q.size-1]
	q.size--
	q.data = q.data[:q.size]
	return queueLastValue
}

func (q *queue) Front() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.size == 0 {
		return -1
	}
	return q.data[0]
}

func (q *queue) Back() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.size == 0 {
		return -1
	}
	return q.data[q.size-1]
}

func (q *queue) Clear() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.size = 0
	q.data = []int{}
	return true
}

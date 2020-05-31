package main

import (
	"fmt"
	"sync"
)

type Queue struct {
	len   int32
	queue []interface{}
	mx    sync.RWMutex
}

func (q *Queue) Enqueue(i interface{}) {
	if i != nil {
		q.mx.Lock()
		q.queue = append(q.queue, i)
		q.mx.Unlock()
	}
}
func (q *Queue) Dequeue() interface{} {
	if len(q.queue) > 0 {
		item := q.queue[0]
		q.mx.Lock()
		q.queue = q.queue[1:]
		q.mx.Unlock()
		return item
	}
	return nil
}

func main() {
	q := &Queue{queue: make([]interface{}, 0)}
	q.Enqueue("selam")
	q.Enqueue("selam1")
	q.Enqueue("selam2")
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}

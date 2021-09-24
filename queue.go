package ds

import "fmt"

type Queue struct {
	c        chan interface{}
	capacity int
}

// NewQueue name says it all
func NewQueue(capacity int) *Queue {
	return &Queue{
		capacity: capacity,
		c:        make(chan interface{}, capacity),
	}
}

// Put a value in the queue
func (q *Queue) Put(val interface{}) error {
	select {
	case q.c <- val:
		return nil
	default:
		return fmt.Errorf("queue is full")
	}
}

// Gets a value from the queue return nil if the queue is empty
func (q *Queue) Get() interface{} {
	select {
	case val := <-q.c:
		return val
	default:
		return nil
	}
}

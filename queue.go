package ds

import "fmt"

type Queue struct {
	c        chan interface{}
	capacity int
	closed   bool
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
	if q.closed {
		return fmt.Errorf("queue is closed")
	}

	select {
	case q.c <- val:
		return nil
	default:
		return fmt.Errorf("queue is full")
	}
}

// Gets a value from the queue return nil if the queue is empty
func (q *Queue) Get() interface{} {
	if q.closed {
		return nil
	}

	select {
	case val := <-q.c:
		return val
	default:
		return nil
	}
}

func (q *Queue) Close() {
	if q.closed {
		return
	}
	q.closed = true
	close(q.c)
}

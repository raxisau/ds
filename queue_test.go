package ds

import (
	"fmt"
	"testing"
)

func TestQueue_FullTest(t *testing.T) {
	t.Run("Testing the Queue Get Function", func(t *testing.T) {
		capacity := 10

		q := NewQueue(capacity)
		valI := q.Get()
		if valI != nil {
			t.Errorf("queue.Get() = %v; want %v", valI, nil)
		}

		for i := 0; i < capacity; i++ {
			val := fmt.Sprint("Test ", i)
			err := q.Put(val)
			if err != nil {
				t.Errorf("queue.Put(%s) = %v; want %v", val, err, nil)
			}
		}

		valFull := "Test full"
		err := q.Put(valFull)
		if err == nil {
			t.Errorf("queue.Put(%s) = %v; want error", valFull, err)
		}

		for i := 0; i < capacity; i++ {
			val := fmt.Sprint("Test ", i)
			valI := q.Get()
			if valI != val {
				t.Errorf("queue.Get() = %v; want %v", valI, val)
			}
		}
		valI = q.Get()
		if valI != nil {
			t.Errorf("queue.Get() = %v; want %v", valI, nil)
		}
	})
}

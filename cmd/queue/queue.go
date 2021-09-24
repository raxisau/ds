package main

import (
	"fmt"

	"github.com/raxisau/ds"
)

func main() {
	capacity := 10

	q := ds.NewQueue(capacity)

	for i := 0; i < capacity; i++ {
		val := fmt.Sprint("Test ", i)
		q.Put(val)
	}

	for i := 0; i < capacity; i++ {
		valI := q.Get()
		val, ok := valI.(string)
		if ok {
			fmt.Println(val)
		} else {
			fmt.Println("Have Problems")
		}
	}
}

# Data Structures and Algorithms

Data structures and algorithms that I need for software

## ScrollBuffer

Implements a Limited size buffer 

```go
    scrollBuff := ds.NewScrollBuffer(500)
    scrollBuff.Put("Item 1")
    scrollBuff.Put("Item 2")
    fmt.Println( scrollBuff.GetBuffer() )
```

## Queue

Implements a Queue using a channel so thread safe

```go
    q := ds.NewQueue(5)
    q.Put("Item 1")
    q.Put("Item 2")
    fmt.Println( q.Get() )
```

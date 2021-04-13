# Data Structures and Algorithms

Data structures and algorithms that I need for software

## ScrollBuffer

Implements a circular queue

```go
    scrollBuff := ds.NewScrollBuffer(500)
    scrollBuff.Put("Item 1")
    scrollBuff.Put("Item 2")
    fmt.Println( scrollBuff.GetBuffer() )
```

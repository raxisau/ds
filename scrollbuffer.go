package ds

// ScrollBuffer is Circular Queue. Add elements to it and at any time
// can dump ot the most recent size elements
type ScrollBuffer struct {
	writePtr  int
	capacity  int
	itemCount int
	buffer    []string
}

// NewScrollBuffer name says it all
func NewScrollBuffer(capacity int) *ScrollBuffer {
	return &ScrollBuffer{
		capacity: capacity,
		buffer:   make([]string, capacity),
	}
}

// Put Name says it all
func (sb *ScrollBuffer) Put(line string) {
	sb.buffer[sb.writePtr] = line

	sb.writePtr++
	if sb.writePtr >= sb.capacity {
		sb.writePtr = 0
	}

	sb.itemCount++
	if sb.itemCount > sb.capacity {
		sb.itemCount = sb.capacity
	}
}

// GetBuffer Name says it all
func (sb *ScrollBuffer) GetBuffer() []string {
	buffer := make([]string, sb.itemCount)
	idx := 0
	if sb.itemCount >= sb.capacity {
		idx = sb.writePtr
	}
	for i := 0; i < sb.itemCount; i++ {
		if idx >= sb.capacity {
			idx = 0
		}
		buffer[i] = sb.buffer[idx]
		idx++
	}
	return buffer
}

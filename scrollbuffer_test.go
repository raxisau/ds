package ds

import (
	"testing"
)

func TestScrollBuffer_GetBuffer(t *testing.T) {
	t.Run("Testing the ScrollBuffer", func(t *testing.T) {
		scrollBuff := NewScrollBuffer(10)
		if scrollBuff.writePtr != 0 {
			t.Errorf("ScrollBuffer.writePtr = %v, want %v", scrollBuff.writePtr, 0)
		}
		if scrollBuff.capacity != 10 {
			t.Errorf("ScrollBuffer.capacity = %v, want %v", scrollBuff.capacity, 10)
		}
		if scrollBuff.itemCount != 0 {
			t.Errorf("ScrollBuffer.itemCount = %v, want %v", scrollBuff.itemCount, 0)
		}

		scrollBuff.Put("Item 1") //0
		scrollBuff.Put("Item 2") //1

		if scrollBuff.writePtr != 2 {
			t.Errorf("ScrollBuffer.writePtr = %v, want %v", scrollBuff.writePtr, 2)
		}
		if scrollBuff.capacity != 10 {
			t.Errorf("ScrollBuffer.capacity = %v, want %v", scrollBuff.capacity, 10)
		}
		if scrollBuff.itemCount != 2 {
			t.Errorf("ScrollBuffer.itemCount = %v, want %v", scrollBuff.itemCount, 2)
		}

		sb := scrollBuff.GetBuffer()
		if len(sb) != 2 {
			t.Errorf("count(ScrollBuffer.GetBuffer()) = %v, want %v", len(sb), 2)
		}
		if sb[0] != "Item 1" {
			t.Errorf("ScrollBuffer.GetBuffer()[0] = %v, want %v", sb[0], "Item 1")
		}
		if sb[1] != "Item 2" {
			t.Errorf("ScrollBuffer.GetBuffer()[1] = %v, want %v", sb[0], "Item 2")
		}

		scrollBuff.Put("Item 1.1")  //2
		scrollBuff.Put("Item 1.2")  //3
		scrollBuff.Put("Item 1.3")  //4
		scrollBuff.Put("Item 1.4")  //5
		scrollBuff.Put("Item 1.5")  //6
		scrollBuff.Put("Item 1.6")  //7
		scrollBuff.Put("Item 1.7")  //8
		scrollBuff.Put("Item 1.8")  //9
		scrollBuff.Put("Item 1.9")  //0
		scrollBuff.Put("Item 1.10") //1
		scrollBuff.Put("Item 1.11") //2

		if scrollBuff.writePtr != 3 {
			t.Errorf("ScrollBuffer.writePtr = %v, want %v", scrollBuff.writePtr, 3)
		}
		if scrollBuff.capacity != 10 {
			t.Errorf("ScrollBuffer.capacity = %v, want %v", scrollBuff.capacity, 10)
		}
		if scrollBuff.itemCount != 10 {
			t.Errorf("ScrollBuffer.itemCount = %v, want %v", scrollBuff.itemCount, 10)
		}
		sb = scrollBuff.GetBuffer()
		if len(sb) != 10 {
			t.Errorf("count(ScrollBuffer.GetBuffer()) = %v, want %v", len(sb), 10)
		}
		if sb[0] != "Item 1.2" {
			t.Errorf("ScrollBuffer.GetBuffer()[0] = %v, want %v", sb[0], "Item 1.2")
		}
		if sb[9] != "Item 1.11" {
			t.Errorf("ScrollBuffer.GetBuffer()[9] = %v, want %v", sb[9], "Item 1.11")
		}
	})
}

package ds

import (
	"testing"
)

func TestInArray_FullTest(t *testing.T) {
	t.Run("Testing the inArray Function for string", func(t *testing.T) {
		strArray := [5]string{"India", "Canada", "Japan", "Germany", "Italy"}

		valB1 := InArray(strArray, "Canada")
		if valB1 != true {
			t.Errorf("inArray(strArray,'Canada') = %v; want %v", valB1, true)
		}

		valB2 := InArray(strArray, "Africa")
		if valB2 != false {
			t.Errorf("inArray(strArray,'Africa') = %v; want %v", valB2, false)
		}
	})
	t.Run("Testing the inArray Function for int", func(t *testing.T) {
		intArray := [7]int{8, 6, 7, 5, 3, 0, 9}

		valB1 := InArray(intArray, 6)
		if valB1 != true {
			t.Errorf("inArray(intArray, 6) = %v; want %v", valB1, true)
		}

		valB2 := InArray(intArray, 27)
		if valB2 != false {
			t.Errorf("inArray(intArray, 27) = %v; want %v", valB2, false)
		}

	})
}

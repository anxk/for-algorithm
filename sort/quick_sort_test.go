package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	cases := NewTestCaseGenerator(maxArrayLength, maxValueInArray)
	for i := 0; i < numOfCases; i++ {
		c := cases.Next()
		input := make([]int, len(c.input))
		copy(input, c.input)
		if QuickSort(c.input); !reflect.DeepEqual(c.input, c.expect) {
			t.Errorf("input: %v, expect: %v, output: %v", input, c.expect, c.input)
		}
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(arrayForBenchmark)
	}
}

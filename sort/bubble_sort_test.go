package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	cases := NewTestCaseGenerator(maxArrayLength, maxValueInArray)
	for i := 0; i < numOfCases; i++ {
		c := cases.Next()
		output := make([]int, len(c.input))
		copy(output, c.input)
		if BubbleSort(output); !reflect.DeepEqual(output, c.expect) {
			t.Errorf("input: %v, expect: %v, output: %v", c.input, c.expect, output)
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(arrayForBenchmark)
	}
}

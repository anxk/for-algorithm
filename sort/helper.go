package sort

import (
	"math/rand"
	"sort"
	"time"
)

const (
	maxArrayLength  = 100
	maxValueInArray = 100
	numOfCases      = 1000
)

var arrayForBenchmark = []int{
	23, 22, -40, 33, -25, 35, 35, 14, 27, -8, 27, -22,
	-38, -16, 13, 24, -44, -4, -28, 0, -4, 1, 6, -8,
	-7, 32, 34, 45, -30, -46, -50, -22, 30, 42, -13,
	30, -17, 36, -44, -34, -40, 20, 42, 12, -3, -41,
	-32, 14, 31, -42, -43, -40, 29, -33, 15, -27, 28, -49,
}

type testCase struct {
	input  []int
	expect []int
}

type TestCaseGenerator struct {
	maxArrayLength  int
	maxValueInArray int
}

func NewTestCaseGenerator(maxArrayLength, maxValueInArray int) *TestCaseGenerator {
	rand.NewSource(time.Now().UnixNano())
	return &TestCaseGenerator{
		maxArrayLength:  maxArrayLength,
		maxValueInArray: maxValueInArray,
	}
}

func (g *TestCaseGenerator) Next() testCase {
	length := rand.Int() % g.maxArrayLength
	input := make([]int, 0)
	expect := make([]int, length)
	for i := 0; i < length; i++ {
		input = append(input, rand.Intn(g.maxValueInArray)-g.maxValueInArray/2)
	}
	copy(expect, input)
	sort.Ints(expect)
	return testCase{
		input:  input,
		expect: expect,
	}
}

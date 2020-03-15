package sort

func MergeSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	refIndex := len(arr) / 2
	arr1, arr2 := arr[0:refIndex], arr[refIndex:]

	MergeSort(arr1)
	MergeSort(arr2)

	merge(arr1, arr2)

	return
}

func merge(arr1, arr2 []int) {
	temp := make([]int, 0)
	s1 := NewStack(arr1)
	s2 := NewStack(arr2)
	for {
		if s1.IsEmpty() {
			temp = append(temp, s2.Left()...)
			break
		}
		if s2.IsEmpty() {
			temp = append(temp, s1.Left()...)
			break
		}
		x := s1.Pop()
		y := s2.Pop()
		if x < y {
			temp = append(temp, x)
			s2.UnPop()
		} else {
			temp = append(temp, y)
			s1.UnPop()
		}
	}
	copy(arr1, temp[0:len(arr1)])
	copy(arr2, temp[len(arr1):])
}

// Stack is a simple stack
type Stack struct {
	Top   int
	Array []int
}

func NewStack(arr []int) *Stack {
	return &Stack{
		Top:   0,
		Array: arr,
	}
}

func (s *Stack) Pop() int {
	s.Top++
	return s.Array[s.Top-1]
}

func (s *Stack) UnPop() {
	s.Top--
}

func (s *Stack) Left() []int {
	return s.Array[s.Top:]
}

func (s *Stack) IsEmpty() bool {
	return s.Top == len(s.Array)
}

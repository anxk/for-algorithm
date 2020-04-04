package main

import (
	"fmt"
)

var count int = 0

func main() {
	arr := []int{0, 0, 0, 0, 0, 0, 0, 0}
	foo(arr, 0)
	fmt.Println(count)
}

func foo(arr []int, curCol int) {
	if curCol == len(arr) {
		fmt.Println()
		display(arr)
		count++
		return
	}
	for i := 0; i < len(arr); i++ {
		arr[curCol] = i
		if isOk(arr, curCol) {
			foo(arr, curCol+1)

		}
	}
}

func isOk(arr []int, col int) bool {
	for i := 0; i < col; i++ {
		if arr[i] == arr[col] || col-i == arr[col]-arr[i] || col-i == arr[i]-arr[col] {
			return false
		}
	}
	return true
}

func display(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] == j {
				fmt.Printf("%s", "x  ")
			} else {
				fmt.Printf("%s", "-  ")
			}
		}
		fmt.Println()
	}
}

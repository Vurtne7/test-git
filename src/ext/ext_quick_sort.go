package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 5, 2, 5, 7, 1, 3, 5, 25, 2, 15}
	fmt.Printf("before arr: %v\n", arr)

	result := quick_sort(arr)
	fmt.Printf("after arr: %v\n", result)
}

func quick_sort(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}
	if len(arr) == 1 {
		return arr
	}

	var left, right, leftArr, rightArr []int
	first := arr[0]
	arr = arr[1:]
	for _, value := range arr {
		if value > first {
			rightArr = append(rightArr, value)
		} else {
			leftArr = append(leftArr, value)
		}
	}

	if len(leftArr) > 0 {
		left = quick_sort(leftArr)
	}
	if len(rightArr) > 0 {
		right = quick_sort(rightArr)
	}

	result := append(left, first)
	return append(result, right...)
}

package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 5, 2, 5, 7, 1, 7, 3, 3, 5, 25, 2, 15}
	fmt.Printf("before arr: %v \n", arr)

	quick_sort2(arr, 0, len(arr)-1)
	fmt.Printf("after arr: %v \n", arr)
}

func quick_sort2(arr []int, l, r int) {
	if l >= r {
		return
	}
	x := arr[(l+r)/2]

	left, right := l, r
	for left < right {
		for arr[left] < x {
			left++
		}
		for arr[right] > x {
			right--
		}

		if left < right {
			if arr[left] == arr[right] {
				left++
				right--
				continue
			}
			arr[left], arr[right] = arr[right], arr[left]
		}
	}

	quick_sort2(arr, l, right)
	quick_sort2(arr, right+1, r)
}

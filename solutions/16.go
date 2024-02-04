package solutions

import "fmt"

func Example16() {
	l := []int{3, 14, 1, 7, 9, 8, 11, 6, 4, 2}
	fmt.Println(quickSortStart(l))
}

func partition(arr []int, left, right int) int {
	pivot := arr[(left+right)/2]
	i := left
	j := right
	for i <= j {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	return j
}

func quickSort(arr []int, left, right int) []int {
	if left < right {
		p := partition(arr, left, right)
		quickSort(arr, left, p-1)
		quickSort(arr, p+1, right)
	}
	return arr
}

func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

package main

import "fmt"

func subSlice(slice []int, begin int, length int, capacity int) []int {
	if capacity < length {
		capacity = length
	}
	array := make([]int, length, capacity)
	slice_len := len(slice)
	j := 0
	if begin >= 0 && begin < slice_len {
		j = begin
	} else {
		j = -1
	}
	for i := 0; i < length; i++ {
		if j != -1 && j < slice_len {
			array[i] = slice[j]
			j++
		} else {
			array[i] = 0
		}
	}
	return array
}

func main() {
	var orig = []int{0, 1, 2, 3, 4, 5}
	var ret []int
	ret = subSlice(orig, 0, 3, 3)
	fmt.Printf("ret = %v, len = %d, cap = %d\n", ret, len(ret), cap(ret))
	ret = subSlice(orig, 2, 7, 10)
	fmt.Printf("ret = %v, len = %d, cap = %d\n", ret, len(ret), cap(ret))

	ret = subSlice(orig, 2, 7, 3)
	fmt.Printf("ret = %v, len = %d, cap = %d\n", ret, len(ret), cap(ret))
}

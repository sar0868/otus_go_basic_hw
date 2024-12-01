package main

import "sort"

func main() {
}

func BinarySearch(data []int, el int) (int, bool) {
	k := len(data) - 1
	p := 0
	for p < k {
		ind := (p + k) / 2
		if data[ind] == el {
			return ind, true
		}
		if el < data[ind] {
			k = ind - 1
		} else {
			p = ind + 1
		}
	}
	return -1, false
}

func SortSlice(array []int) []int {
	sort.Slice(array, func(i, j int) bool {
		return array[i] < array[j]
	})
	return array
}

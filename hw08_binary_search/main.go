package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	ind, res := BinarySearch(arr, 3)
	fmt.Println(ind, res)
}

func BinarySearch(data []int, el int) (int, bool) {
	k := len(data) - 1
	p := 0
	for p < k {
		ind := p + (k-p)/2
		if data[ind] == el {
			return ind, true
		}
		if el < data[ind] {
			k = ind
		} else {
			p = ind
		}
	}
	return -1, false
}

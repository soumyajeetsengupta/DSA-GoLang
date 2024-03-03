package main

import "fmt"

func max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	var maxElem int = arr[0]

	for i := 0; i < len(arr)-1; i++ {
		maxElem = max(arr[i], arr[i+1])
	}

	fmt.Println(maxElem)
}
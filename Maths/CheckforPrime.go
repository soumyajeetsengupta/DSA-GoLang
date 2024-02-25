package main

import (
	"fmt"
	"math"
)

func main() {
	var inputNum int = 7
	var flag bool = true

	for i := 2; i <= int(math.Sqrt(float64(inputNum))); i++ {
		if inputNum%i == 0 {
			flag = false
			break
		}
	}

	fmt.Println(flag)
}

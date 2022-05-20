//https://codeforces.com/gym/103714/problem/A
package main

import (
	"fmt"
	"math"
)

func main() {
	var sheets int
	fmt.Scanf("%d", &sheets)

	//var minCost int
	var price int
	var paperPrice [5]int
	for i := 0; i < 5; i++ {
		fmt.Scanf("%d", &price)
		paperPrice[i] = price
	}

	//get the lowest
	var low float64
	var exp float64
	exp = 8
	for i := 0; i < 5; i++ {
		result := math.Ceil(float64(sheets)/exp) * float64(paperPrice[i])
		if i == 0 {
			low = result
		} else {
			if result < low {
				low = result
			}
		}
		exp = exp - 2
	}

	fmt.Println(int(low))
}

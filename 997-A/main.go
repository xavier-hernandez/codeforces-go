// https://codeforces.com/problemset/problem/977/A
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var k int
	fmt.Scanf("%d", &k)

	for !(k == 0) {
		if n%10 == 0 && n > 0 {
			n = n / 10
		} else {
			n--
		}
		k--
	}

	fmt.Println(n)
}

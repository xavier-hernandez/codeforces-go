package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	argN, _ := strconv.Atoi(os.Args[1])
	argK, _ := strconv.Atoi(os.Args[2])

	for !(argK == 0) {
		if argN%10 == 0 && argN > 0 {
			argN = argN / 10
		} else {
			argN--
		}
		argK--
	}

	fmt.Println(argN)
}

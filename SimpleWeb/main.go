package main

import (
	"fmt"
)

func main() {
	var n, x int
	fmt.Scan(&n, &x)
	count := 0
	for i := 0; i < n; i++ {
		var c string
		var d int
		fmt.Scan(&c, &d)
		if c == "+" {
			x += d
		} else {
			if x < d {
				count++
			} else {
				x -= d
			}
		}
	}
	fmt.Print(x, " ", count)

}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var t int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)
		f := 1
		s := int(math.Ceil((float64(n)*float64(n))/2)) + 1
		if n == 2 {
			fmt.Println(-1)
			continue
		}
		fmt.Println(f, s)
		for j := 1; j <= n; j++ {
			for k := 1; k <= n; k++ {
				if (k+j)%2 == 0 {
					fmt.Print(f, " ")
					f++
				} else {
					fmt.Print(s, " ")
					s++
				}
			}
			fmt.Println()
		}
	}
}

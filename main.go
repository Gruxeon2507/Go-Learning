package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var a, b, c int
		fmt.Scan(in, &a, &b, &c)

	}
}

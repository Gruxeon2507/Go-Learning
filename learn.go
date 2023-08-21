package main

import (
	"fmt"
)

// chi dc khai bao full
var test int = 27
var (
	actorName     string = "Gruxeon"
	actorNickName string = "Duckm"
	i             int    = 42
)
func learn() {
	fmt.Println(i)
	//shadowing: bien trong scope dc khai bao cung ten se dc uu tien
	var i int = 27
	fmt.Println(i)
	//This is declare a new varable
	j := 21.
	var k int
	k = 20
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Printf("%v, %T\n", j, j)
	//Naming Conversion
	//Short name for instance use of variable
	//acronym: keep it as it is (HTTP, URL, etc...)

	//convert
	var integer int = 42
	fmt.Printf("%v,%T\n", integer, integer)

	var convertInteger float32
	convertInteger = float32(integer)
	fmt.Printf("%v, %T\n", convertInteger, convertInteger)

	var convertString string
	convertString = string(integer)
	// value assign to * because 42 in ASCII = "*", to convert and keep the original value, use pakage "strconv"
	fmt.Printf("%v, %T\n", convertString, convertString)

}



package main

import "fmt"

func main() {
	fmt.Println("Slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	//[Apple Tomato Peach Mango Banana]
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	// //[Tomato Peach Mango Banana]
	// fruitList = append(fruitList[1:])
	// fmt.Println(fruitList)

	// //[Peach Mango] element no 1 and 2
	// fruitList = append(fruitList[1:3])
	// fmt.Println(fruitList)

	// //[Apple Tomato Peach]
	// fruitList = append(fruitList[:3])
	// fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 345
	highScore[2] = 456
	highScore[3] = 678

	highScore = append(highScore, 555, 666, 777)
	fmt.Println(highScore)
}

package main

import (
	"fmt"
)

var staticArray [5]int

func main() {
	staticArray[0] = 1
	staticArray[1] = 22
	staticArray[2] = 5
	staticArray[3] = 4
	staticArray[4] = 3

	staticArray[2] = 4
	fmt.Println(staticArray, staticArray[2])
}

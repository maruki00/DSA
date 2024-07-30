package main

var staticArray []int

func remove(index int) {
	staticArray = append(staticArray[:index], staticArray[index:]...)
}

func add(item int) {
	staticArray = append(staticArray, item)
}

func getAt(index int) int {
	return staticArray[index]
}

func main() {

}

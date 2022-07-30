package main

import "fmt"

func main() {

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range numbers {
		var category string
		if num%2 == 0 {
			category = "Even"
		} else {
			category = "Odd"
		}
		fmt.Println(num, "is", category)
	}

}

package main

import "fmt"

func main() {
	mySlice := getSlice(2, 33)
	checkSlice(mySlice)
}

func getSlice(from int, to int) []int {
	var intSlice []int
	diff := to - from

	for i := 0; i <= diff; i++ {
		intSlice = append(intSlice, from+i)
	}

	return intSlice
}

func checkSlice(s []int) {
	for _, value := range s {
		if value%2 == 0 {
			fmt.Println(value, "is even")
		} else {
			fmt.Println(value, "is odd")
		}
	}
}

package main

import "fmt"

func fSlice() {
	mySlice := []string{"Hi", "there", "are", "you?"}

	fmt.Println(mySlice)

	updateSlice(mySlice)

	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}

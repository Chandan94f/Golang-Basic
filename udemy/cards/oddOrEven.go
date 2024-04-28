package main

import "fmt"

func printOddOrEven() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, i := range nums {
		if i%2 == 0 {
			fmt.Println(i, " is a even number")
		} else {
			fmt.Println(i, " is a odd number")
		}
	}

}

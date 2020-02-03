package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, ele := range a {
		if ele%2 == 0 {
			fmt.Println(ele, "is even")
		} else {
			fmt.Println(ele, "is odd")
		}
	}
}

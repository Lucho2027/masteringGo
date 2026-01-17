package main

import "fmt"

func foo(s []int) int {
	return s[0] + s[1] + s[2] + s[3]
}

func bar(slice []int) int {
	a := (*[3]int)(slice)
	return a[0] + a[1] + a[2] + a[3]
}

func main() {
	car := bar([]int{1, 2, 3, 4, 5})
	fmt.Println("car", car)
}

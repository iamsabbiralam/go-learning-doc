package main

import "fmt"

func main() {
	
	a := []int{1,3,4,5}
	res := sum(a)
	fmt.Println("slice: ", res)
	res = multi(a)
	fmt.Println("slice: ", res)
}

func sum(number []int) int {
	total := 0
	for _, num:= range number {
		total += num
	}
	return total
}

func multi(number []int) int {
	total := 1
	for _, num:= range number {
		total *= num
	}
	return total
}
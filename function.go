package main

import "fmt"

func main() {

	// a := []int{1,2,3,4}
	res := data(sum)
	fmt.Println("Additional Result: ", res)
	res = data(multi)
	fmt.Println("Multiple Result: ", res)
	fmt.Println("Fatal: ", fact(8))
	nxtInt := seqInt()
	fmt.Println("Function: ", nxtInt())
	fmt.Println("Function: ", nxtInt())
}

func sum(number []int) int {
	total := 0
	for _, num := range number {
		total += num
	}
	return total
}

func multi(number []int) int {
	total := 1
	for _, num := range number {
		total *= num
	}
	return total
}

func data(a func(number []int) int) int {

	b := []int{1,2,3,4}

	return a(b)
}

func fact(a int) int {
	if a == 0 {
		return 1
	}
	return a * fact(a - 1) // 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1 * 1
}

func seqInt() func() int {
	i := 0
	return func () int {
		i++
		return i
	}
}
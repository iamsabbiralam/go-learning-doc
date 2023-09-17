package main

import "fmt"

func main() {
	n := 7
	for i := 0; i < n; i++ {
		for j := 0; j < n+i; j++ {
			if j < n-i-1 {
				fmt.Printf(" ")
			} else {
				fmt.Printf("*")
			}
		}
		fmt.Printf("\n")
	}
}

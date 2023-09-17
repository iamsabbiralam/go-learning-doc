package main

import "fmt"

// type Vertex struct {
// 	Lat, Long float64
// }

// var m map[string]Vertex

// func main() {
// 	m = make(map[string]Vertex)
// 	m["Bell Labs"] = Vertex{
// 		40.68433, -74.39967,
// 	}
// 	fmt.Println(m["Bell Labs"])
// }

var a string = "codemen"

var m map[string]string

func main() {
	m = make(map[string]string)
	m["Hello String"] = a
	fmt.Println(m["Hello String"])
}

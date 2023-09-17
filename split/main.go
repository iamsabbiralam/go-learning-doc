package main

import (
	"fmt"
	"strings"
)

func main () {
	str := `[{"value":"Tate Beard"},{"value":"Laravel"},{"value":"go"}]`
	str1 := strings.Split(str, `"}]`)
	fmt.Println(str1)
	str2 := strings.Split(string(str1[0][11:]), `"},{"value":"`)
	
	fmt.Println(str2)
}
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hello World")
	fmt.Println("Hello" + " World")
	fmt.Println("Hello World")

	s := "l"
	fmt.Println(strings.Contains("Hello world!", s))

	x := 61
	y := float64(x)
	fmt.Println(y)

	var d string = "143"
	i, _ := strconv.Atoi(d)
	fmt.Println(i + 1)

	a := []int{2, 3, 4, 5, 6}
	fmt.Printf("%T /n¥n", a)

	n := make([]int, 4, 6)
	fmt.Printf("len=%d cap=%d value=%v ¥n ", len(n), cap(n), n)

	m := map[string]int{"apple": 100, "orange": 200}
	fmt.Println(m)
	v := m["apple"]
	fmt.Println(v)
	v2, ok := m["orange"]
	fmt.Println(v2, ok)

	fmt.Println(add(2, 3))
}

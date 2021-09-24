package main

import "fmt"

func main() {
	defer cmd()
	fmt.Println("UMMMMM")
}

func cmd() {
	fmt.Println("CMD:Hello, World!")
}

package util

import (
	"bufio"
	"fmt"
	"os"
)

func MakeArray() []string {
	fmt.Println("* ******* MakeArray() is called ******** *")
	sc := bufio.NewScanner(os.Stdin)
	var s string
	var slice []string
	for sc.Scan() {
		s = sc.Text()
		if s == "Q" {
			break
		}
		slice = append(slice, s)
	}
	fmt.Println("* ******** Returns ********************* *")
	fmt.Println(slice)
	return slice
}

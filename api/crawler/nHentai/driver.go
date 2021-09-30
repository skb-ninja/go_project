package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sclevine/agouti"
)

func getDriver() *agouti.WebDriver {
	options := agouti.ChromeOptions(
		"args", []string{
			"--disable-gpu",
		})
	driver := agouti.ChromeDriver(options)

	return driver
}

func getScanner() *bufio.Scanner {
	return bufio.NewScanner(os.Stdin)
}

func makeArray() []string {
	var s []string
	sc := getScanner()
	fmt.Println("makeArray() is Called!")
	for {
		sc.Scan()
		var p string = sc.Text()
		if p == "Q" {
			break
		}

		p = strings.TrimSpace(p)
		s = append(s, p)
	}
	return s
}

func main() {
	driver := getDriver()
	err := driver.Start()
	if err != nil {
		fmt.Println(err)
	}

	page, err := driver.NewPage()
	if err != nil {
		fmt.Println(err)
	}
	page.Navigate("https://google.com")
	fmt.Println(page.Title())

	defer driver.Stop()

}

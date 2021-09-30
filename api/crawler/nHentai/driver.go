package main

import (
	"fmt"

	"github.com/sclevine/agouti"
	"github.com/skb-ninja/go_project/util"
)

func getDriver() *agouti.WebDriver {
	options := agouti.ChromeOptions(
		"args", []string{
			"--disable-gpu",
		})
	driver := agouti.ChromeDriver(options)

	return driver
}

func DriverMain() {

	//srcList Array
	s := util.MakeArray()

	//accesing each site &&
	driver := getDriver()
	err := driver.Start()
	if err != nil {
		fmt.Println(err)
	}
	page, err := driver.NewPage()
	if err != nil {
		fmt.Println(err)
	}
	for _, k := range s {
		page.Navigate(k)
		fmt.Println(page.Title())
	}

	defer driver.Stop()

}

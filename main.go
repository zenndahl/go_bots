package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
)

func main() {
	fmt.Println("Hello World")

	service, err := selenium.NewChromeDriverService("./chromedriver", 4444)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{}
	// caps.AddChrome(chrome.Capabilities{Args: []string{
	// 	"--headless-new", // comment out this line for testing
	// }})

	// create a new remote client with the specified options
	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Fatal("Error:", err)
	}

	waitTime := 2 * time.Second
	selenium.WebDriver.SetImplicitWaitTimeout(driver, waitTime)

	// maximize the current window to avoid responsive rendering
	err = driver.MaximizeWindow("")
	if err != nil {
		log.Fatal("Error:", err)
	}

	err = driver.Get("https://www.google.com/")
	if err != nil {
		log.Fatal("Error:", err)
	}

	// time.Sleep(5 * time.Second)

	elem, err := driver.FindElement(selenium.ByCSSSelector, "#APjFqb")
	if err != nil {
		log.Fatal("Error:", err)
	}

	//fmt.Print(elem.Location())

	err = elem.Click()
	if err != nil {
		log.Fatal("Error:", err)
	}

	err = elem.SendKeys("golang tutorial")
	if err != nil {
		log.Fatal("Error:", err)
	}

	driver.KeyDown(selenium.EnterKey)

	time.Sleep(10 * time.Second)

	driver.Quit()
}

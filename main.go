package main

import (
	"fmt"
	"time"

	"github.com/kaduh15/TempWiFi-Creator/driver"
	"github.com/playwright-community/playwright-go"
)

func main() {
	if err := playwright.Install(); err != nil {
		fmt.Println("could not install playwright")
		return
	}

	pw, err := playwright.Run()
	if err != nil {
		fmt.Println("could not run playwright")
		return
	}

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Args:     []string{"--no-sandbox --headless=new"},
		Headless: playwright.Bool(false),
	})

	if err != nil {
		fmt.Println("could not launch browser")
		return
	}

	page, err := browser.NewPage()
	if err != nil {
		fmt.Println("could not create page")
		return
	}

	driver.Login(page, "multipro", "multipro")
	time.Sleep(5 * time.Second)

	// Close the browser
	if err := browser.Close(); err != nil {
		fmt.Println("could not close browser")
		return
	}

	// Close Playwright
	if err := pw.Stop(); err != nil {
		fmt.Println("could not close playwright")
		return
	}
}

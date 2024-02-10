package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/kaduh15/TempWiFi-Creator/driver"
	"github.com/mdp/qrterminal/v3"
	"github.com/playwright-community/playwright-go"
)

const version = "0.1.0"

func main() {
	fmt.Println("TempWiFi-Creator v" + version)

	if err := playwright.Install(); err != nil {
		fmt.Println("could not install playwright")
		return
	}

	pw, err := playwright.Run()
	if err != nil {
		fmt.Println("could not run playwright")
		return
	}

	var opt playwright.BrowserTypeLaunchOptions

	if os.Getenv("HEADLESS") == "true" {
		opt.Args = []string{"--no-sandbox --headless=new"}
		opt.Headless = playwright.Bool(true)
	}

	browser, err := pw.Chromium.Launch(opt)

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
	time.Sleep(1 * time.Second)

	network := driver.GenerateNetwork(page, "Loja")
	time.Sleep(1 * time.Second)

	ShowQrCodeWifi(network)
	fmt.Println("Aguarde o QR Code ser escaneado...")
	fmt.Println("=====================================")
	fmt.Println("Nome da rede: ", network.Name)
	fmt.Println("Senha da rede: ", network.Password)
	fmt.Println("=====================================")
	esperarEnter()
	driver.DisableWifi(page)

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

func ShowQrCodeWifi(network driver.Network) {
	name := network.Name
	password := network.Password

	config := qrterminal.Config{
		Level:     qrterminal.L,
		Writer:    os.Stdout,
		BlackChar: qrterminal.WHITE,
		WhiteChar: qrterminal.BLACK,
		QuietZone: 1,
	}
	qrterminal.GenerateWithConfig("WIFI:T:WPA;S:"+name+";P:"+password+";;", config)
}

func esperarEnter() {
	fmt.Println("Pressione enter para sair...")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}

package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	qrcode "github.com/skip2/go-qrcode"

	"github.com/kaduh15/TempWiFi-Creator/config"
	"github.com/kaduh15/TempWiFi-Creator/driver"
	"github.com/playwright-community/playwright-go"
)

func main() {

	config.InitDotEnv()

	if err := playwright.Install(); err != nil {
		fmt.Println("Erro ao instalar o Playwright")
		panic(err)
	}

	pw, err := playwright.Run()
	if err != nil {
		fmt.Println("Erro ao executar o Playwright")
		panic(err)
	}
	defer pw.Stop()

	var arg []string
	var headless *bool

	if os.Getenv("GO_DEV") == "true" {
		arg = []string{"--sandbox --headless=new"}
		headless = playwright.Bool(false)
	}

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Args:     arg,
		Headless: headless,
	})
	if err != nil {
		fmt.Println("Erro ao abrir o navegador")
		panic(err)
	}

	page, err := browser.NewPage()
	if err != nil {
		fmt.Println("Erro ao abrir a página")
		panic(err)
	}

	a := app.New()
	w := a.NewWindow("Hello World")
	w.SetCloseIntercept(func() {
		fmt.Println("Fechando a aplicação")
		driver.DisableWifi(page)

		browser.Close()
		pw.Stop()
		w.Close()
	})

	var network driver.Network
	var image *canvas.Image
	var buttonDisable *widget.Button
	var buttonEnable *widget.Button
	titleText := widget.NewLabel("Clique no botão para criar a rede temporária")

	boxV := container.NewVBox()

	buttonDisable = widget.NewButton("Desativar Rede Temporária", func() {
		titleText.Text = "Aguarde enquanto a rede é desativada"
		titleText.Refresh()
		driver.DisableWifi(page)
		image.Hide()
		buttonEnable.Enable()
		buttonEnable.Show()
		buttonDisable.Hide()
		boxV.Remove(image)
		titleText.Text = "Clique no botão para criar a rede temporária"
		titleText.Refresh()
	})

	buttonDisable.Hide()

	buttonEnable = widget.NewButton("Criar Rede Temporária", func() {
		buttonEnable.Disable()
		titleText.Text = "Aguarde enquanto a rede é criada"
		titleText.Refresh()
		driver.Login(page)
		network = driver.GenerateNetwork(page)
		name := network.Name
		password := network.Password

		err := qrcode.WriteFile(
			"WIFI:T:WPA;S:"+name+";P:"+password+";;", qrcode.Medium,
			256,
			"qr.png",
		)
		if err != nil {
			fmt.Println("Erro ao gerar o QRCode")
			panic(err)
		}
		image = canvas.NewImageFromFile("qr.png")
		image.FillMode = canvas.ImageFillOriginal
		image.Show()
		boxV.Add(image)
		buttonEnable.Hide()
		buttonDisable.Show()
		titleText.Text = "Rede criada com sucesso \nNome: " + name + "\nSenha: " + password
		titleText.Refresh()
	})

	boxV.Add(titleText)
	boxV.Add(buttonEnable)
	boxV.Add(buttonDisable)

	w.SetContent(boxV)
	w.ShowAndRun()
}

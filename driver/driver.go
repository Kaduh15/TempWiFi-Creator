package driver

import (
	"github.com/lucsky/cuid"
	"github.com/playwright-community/playwright-go"
)

func Login(page playwright.Page, username, password string) {
	// Navigate to the login page
	page.Goto("http://192.168.1.1")

	inputUsername := page.Locator("input[id='Frm_Username']")
	inputUsername.Fill(username)

	inputPassword := page.Locator("input[id='Frm_Password']")
	inputPassword.Fill(password)

	// Click the login button
	page.Locator("input[id='LoginId']").Click()
}

type Network struct {
	Name     string
	Password string
}

func GenerateNetwork(page playwright.Page, ssid string) Network {
	slug := cuid.Slug()
	nameNet := ssid + " - " + slug
	password := cuid.Slug()

	btnlocalNetwork := page.Locator("a[id='localnet']")
	btnlocalNetwork.Click()

	btnWlanConfig := page.Locator("a[id='wlanConfig']")
	btnWlanConfig.Click()

	btnWlanSSID := page.Locator("h1[id='WLANSSIDConfBar']")
	btnWlanSSID.Click()

	btnSSID := page.Locator("a[id='instName_WLANSSIDConf:1']")
	btnSSID.Click()

	inputSSIDName := page.Locator("input[id='ESSID:1']")
	inputSSIDName.Clear()
	inputSSIDName.PressSequentially(nameNet)

	inputSSIDPassword := page.Locator("input[id='KeyPassphrase:1'][type='password']")
	inputSSIDPassword.Clear()
	inputSSIDPassword.PressSequentially(password)

	btnEnableNet := page.Locator("input[id='Enable1:1']")
	btnEnableNet.Click()

	btnApply := page.Locator("input[id='Btn_apply_WLANSSIDConf:1']")
	btnApply.Click()

	btnConfirm := page.Locator("input[id='confirmOK'][value='OK']")
	btnConfirm.Click()

	return Network{
		Name:     nameNet,
		Password: password,
	}
}

func DisableWifi(page playwright.Page) {
	btnEnableNet := page.Locator("input[id='Enable0:1']")
	btnEnableNet.Click()

	btnApply := page.Locator("input[id='Btn_apply_WLANSSIDConf:1']")
	btnApply.Click()
}

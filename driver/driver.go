package driver

import "github.com/playwright-community/playwright-go"

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



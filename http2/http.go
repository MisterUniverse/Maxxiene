package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mxschmitt/playwright-go"
)

func assertError(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func startBrowser(br *playwright.Browser, ip string, w, h int) (pWr playwright.Page, e error) {
	// Setup pw
	pw, err := playwright.Run()
	assertError("could not launch playwright: %w", err)

	//Browser
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	assertError("could not launch Chromium: %w", err)
	br = &browser

	vp := playwright.BrowserNewContextOptionsViewport{
		Width:  &w,
		Height: &h,
	}
	ctx, err := browser.NewContext(playwright.BrowserNewContextOptions{
		Viewport: &vp,
	})
	assertError("could not create context: %w", err)

	// Page
	page, err := ctx.NewPage()
	assertError("could not create page: %w", err)
	_, err = page.Goto("http://" + ip + "/login")
	assertError("could not goto: %w", err)
	return page, nil
}

func logIn(page playwright.Page, ip string, pw string) error {
	entries, err := page.QuerySelectorAll(".form-input")
	assertError("failed getting login input boxes", err)

	entries[0].Type("admin")
	entries[1].Type(pw)

	sub, err := page.QuerySelectorAll("button.form-submit")
	assertError("failed to acquire button elements", err)

	sub[0].Press("Enter")
	_, err = page.Goto("http://" + ip + "/administration#/firewall")
	assertError("failed loading firewall page %w", err)
	return nil
}

func screenShot(page playwright.Page, fName string, b bool) error {
	if _, err := page.Screenshot(playwright.PageScreenshotOptions{
		Path:     playwright.String(fName),
		FullPage: &b,
	}); err != nil {
		return err
	}
	return nil
}

func setIp(page playwright.Page, newIp string, sub string, curIp string) error {
	page.Goto("http://" + curIp + "/administration#/network-settings")
	time.Sleep(3 * time.Second)
	net, errr := page.QuerySelectorAll("[name=eth-mode]")
	assertError("failed to grab radio buttons", errr)

	for i := range net {
		fmt.Println(i)
	}

	net[1].Check()

	ip, aerr := page.QuerySelectorAll("[type=text]")
	assertError("failed to grab input boxes %w", aerr)

	//Need to change how gate way is
	ip[0].Fill(newIp)
	ip[1].Fill(sub)
	ip[2].Fill("192.168.45.1")

	btn, e := page.QuerySelector(".flir-button.large")
	assertError("failed to grab apply button", e)

	btn.Press("Enter")

	fmt.Println("Ip set")
	return nil
}

func main() {
	var browser playwright.Browser
	//715 pw = "IRmRlSkt"
	ipAddress := "192.168.45.101"
	sub := "255.255.255.0"
	//gate := "192.168.45.1"
	// username := "
	password := "XJCUunsi"
	b := true
	w := 1920
	h := 1080

	page, err := startBrowser(&browser, ipAddress, w, h)
	assertError("failed to init browser %w", err)

	logIn(page, ipAddress, password)

	time.Sleep(7 * time.Second)

	screenShot(page, "firewall.png", b)

	page.Goto("http://192.168.45.101/administration#/system")

	time.Sleep(7 * time.Second)

	screenShot(page, "firmware.png", b)

	setIp(page, ipAddress, sub, ipAddress)

	_, err = page.Reload()
	assertError("could not reload: %w", err)

	//assertErrorToNilf("could not close browser: %w", browser.Close())
	//assertErrorToNilf("could not stop Playwright: %w", pw.Stop())
}

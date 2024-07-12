package test

import (
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/playwright-community/playwright-go"
)

var Playwright *playwright.Playwright
var Browser playwright.Browser
var Page playwright.Page
var BaseUrl = "http://localhost:5173"

func SetupTest() {
	err := godotenv.Load("../.env", "../.env.test")
	if err != nil {
		slog.Error("could not load env: %v", "error", err)
	}

	Playwright, err = playwright.Run(&playwright.RunOptions{
		Verbose: true,
	})
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	Browser, err = Playwright.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Timeout:  playwright.Float(10000),
		Headless: playwright.Bool(os.Getenv("CI") == "true"),
	})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	Page, err = Browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
}

func TeardownTest() {
	if err := Browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err := Playwright.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}

func Goto(path string) {
	if _, err := Page.Goto(BaseUrl + path); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	if err := Page.Locator("body.hydrated").WaitFor(); err != nil {
		log.Fatalf("could not wait for body.hydrated: %v", err)
	}
}

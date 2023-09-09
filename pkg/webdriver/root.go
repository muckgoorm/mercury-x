package webdriver

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type wd struct {
	Driver selenium.WebDriver
}

func NewWebDriver(driverUrl string) (*wd, error) {
	caps := selenium.Capabilities{"browserName": "chrome"}
	caps.AddChrome(chrome.Capabilities{
		Args: []string{
			"--no-sandbox",
			"--disable-dev-shm-usage",
			"disable-gpu",
			// "--headless", // comment out this line to see the browser
		},
	})

	driver, err := selenium.NewRemote(caps, driverUrl)
	if err != nil {
		return nil, err
	}

	return &wd{Driver: driver}, nil
}

func (w *wd) Get(url string) error {
	if err := w.Driver.Get(url); err != nil {
		return err
	}

	return nil
}

func (w *wd) WaitForDivClass(className string) error {
	divXPath := fmt.Sprintf("//div[@class='%s']", className)
	if err := w.Driver.WaitWithTimeoutAndInterval(selenium.Condition(func(wd selenium.WebDriver) (bool, error) {
		_, err := wd.FindElement(selenium.ByXPATH, divXPath)
		if err != nil {
			return false, nil
		}
		return true, nil
	}), 5*time.Second, 1*time.Second); err != nil {
		return err
	}

	return nil
}

func (w *wd) FulfillInput(className, value string) error {
	inputXPath := fmt.Sprintf("//div[@class='%s']//input", className)
	input, err := w.Driver.FindElement(selenium.ByXPATH, inputXPath)
	if err != nil {
		return err
	}
	if err := input.SendKeys(value); err != nil {
		return err
	}
	time.Sleep(1 * time.Second)
	if err := input.SendKeys(selenium.EnterKey); err != nil {
		return err
	}

	return nil
}

func (w *wd) ScrollToBottom() error {
	previousHeight := 0.0
	for {
		currentHeight, err := w.Driver.ExecuteScript("return document.body.scrollHeight", nil)
		if err != nil {
			return err
		}
		if previousHeight == currentHeight {
			break
		}
		previousHeight = currentHeight.(float64)
		if _, err := w.Driver.ExecuteScript("window.scrollTo(0, document.body.scrollHeight)", nil); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}

	return nil
}

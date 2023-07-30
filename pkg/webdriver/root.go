package webdriver

import (
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

func (w *wd) ScrollToBottom(count int) error {
	for i := 0; i < count; i++ {
		if _, err := w.Driver.ExecuteScript("window.scrollTo(0, document.body.scrollHeight)", nil); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}

	return nil
}

func (w *wd) FindByClassName(className string) ([]string, error) {
	elements, err := w.Driver.FindElements(selenium.ByClassName, className)
	if err != nil {
		return nil, err
	}

	texts := make([]string, 0)
	for _, element := range elements {
		text, err := element.Text()
		if err != nil {
			return nil, err
		}
		texts = append(texts, text)
	}

	return texts, nil
}

func (w *wd) FindLinks(className string) ([]string, error) {
	elements, err := w.Driver.FindElements(selenium.ByClassName, className)
	if err != nil {
		return nil, err
	}

	links := make([]string, 0)
	for _, element := range elements {
		aTag, err := element.FindElement(selenium.ByTagName, "a")
		if err != nil {
			return nil, err
		}
		href, err := aTag.GetAttribute("href")
		if err != nil {
			return nil, err
		}
		links = append(links, href)
	}

	return links, nil
}

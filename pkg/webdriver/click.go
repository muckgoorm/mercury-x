package webdriver

import (
	"fmt"

	"github.com/tebeka/selenium"
)

func (w *wd) ClickButton(className string) error {
	buttonXPath := fmt.Sprintf("//button[@class='%s']", className)
	button, err := w.Driver.FindElement(selenium.ByXPATH, buttonXPath)
	if err != nil {
		return err
	}
	if err := button.Click(); err != nil {
		return err
	}

	return nil
}

func (w *wd) ClickButtonByDivClass(className string) error {
	buttonXPath := fmt.Sprintf("//div[@class='%s']//button", className)
	button, err := w.Driver.FindElement(selenium.ByXPATH, buttonXPath)
	if err != nil {
		return err
	}
	if err := button.Click(); err != nil {
		return err
	}

	return nil
}

func (w *wd) ClickButtonByDataAttributeId(id string) error {
	buttonXPath := fmt.Sprintf("//button[@data-attribute-id='%s']", id)
	button, err := w.Driver.FindElement(selenium.ByXPATH, buttonXPath)
	if err != nil {
		return err
	}
	if err := button.Click(); err != nil {
		return err
	}

	return nil
}

func (w *wd) ClickButtonByDataTag(tag string) error {
	buttonXPath := fmt.Sprintf("//button[@data-tag='%s']", tag)
	button, err := w.Driver.FindElement(selenium.ByXPATH, buttonXPath)
	if err != nil {
		return err
	}
	if err := button.Click(); err != nil {
		return err
	}

	return nil
}

package webdriver

import "github.com/tebeka/selenium"

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

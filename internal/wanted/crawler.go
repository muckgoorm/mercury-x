package wanted

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mercury-x/internal"
	"mercury-x/pkg/webdriver"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
)

type WantedCrawler struct {
	internal.Crawler
}

func NewWantedCrawler() internal.Crawler {
	return &WantedCrawler{}
}

func (c *WantedCrawler) SearchJobPostings(j internal.JobSearchPayload) ([]internal.JobPosting, error) {
	role := mapRole(j.Role)
	exp := mapExperience(j.Experience)

	ub := strings.Builder{}
	ub.WriteString("https://www.wanted.co.kr/wdlist/")
	ub.WriteString(dev)
	ub.WriteString(role)
	ub.WriteString("?country=kr&job_sort=job.latest_order&years=")
	ub.WriteString(exp)
	ub.WriteString("&locations=all")
	url := ub.String()
	fmt.Println(url)

	wd, err := webdriver.NewWebDriver("http://localhost:9515")
	if err != nil {
		return nil, err
	}
	defer wd.Driver.Quit()

	if err := wd.Get(url); err != nil {
		return nil, err
	}

	if len(j.Stacks) > 0 {
		if err := wd.ClickButtonByDivClass(skillAddButton); err != nil {
			return nil, err
		}

		for _, stack := range j.Stacks {
			if err := wd.FulfillInput(skillInput, stack); err != nil {
				return nil, err
			}
		}

		if err := wd.ClickButtonByDataAttributeId(skillApplyButton); err != nil {
			return nil, err
		}
		wd.WaitForDivClass(jobList)
	}

	if len(j.Benefits) > 0 {
		benefitTags := mapBenefitsToTags(j.Benefits)
		for _, tag := range benefitTags {
			switch tag {
			case workFromHomeTag:
				if err := wd.ClickButtonByDataTag(workFromHomeTag); err != nil {
					return nil, err
				}
			case flexibleTag:
				if err := wd.ClickButton(nextButton); err != nil {
					return nil, err
				}
				time.Sleep(1 * time.Second)
				if err := wd.ClickButtonByDataTag(flexibleTag); err != nil {
					return nil, err
				}
			case flatTag:
				if err := wd.ClickButtonByDataTag(flatTag); err != nil {
					return nil, err
				}
			case snackTag:
				if err := wd.ClickButton(nextButton); err != nil {
					return nil, err
				}
				time.Sleep(1 * time.Second)
				if err := wd.ClickButtonByDataTag(snackTag); err != nil {
					return nil, err
				}
			}
			time.Sleep(1 * time.Second)
		}
	}

	time.Sleep(1 * time.Second)

	scrollCount := (j.Count / 20) - 1
	if scrollCount > 0 {
		if err := wd.ScrollToBottom(scrollCount); err != nil {
			return nil, err
		}
	}

	companies, err := wd.FindByClassName(company)
	if err != nil {
		return nil, err
	}

	roles, err := wd.FindByClassName(position)
	if err != nil {
		return nil, err
	}

	urls, err := wd.FindLinks(card)
	if err != nil {
		return nil, err
	}

	var postings []internal.JobPosting
	for i := 0; i < len(companies); i++ {
		postings = append(postings, internal.JobPosting{
			Platform: "wanted",
			Company:  companies[i],
			Role:     roles[i],
			URL:      urls[i],
		})
	}

	return postings, nil
}

func (c *WantedCrawler) ParseJobDescription(url string) (internal.JobDescription, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return internal.JobDescription{}, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0")
	resp, err := client.Do(req)
	if err != nil {
		return internal.JobDescription{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return internal.JobDescription{}, err
	}

	html, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		return internal.JobDescription{}, err
	}

	script := findNextDataScript(html)
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(script), &result); err != nil {
		return internal.JobDescription{}, err
	}

	props, ok := result["props"].(map[string]interface{})
	if !ok {
		return internal.JobDescription{}, errors.New("props object not found")
	}
	pageProps, ok := props["pageProps"].(map[string]interface{})
	if !ok {
		return internal.JobDescription{}, errors.New("pageProps object not found")
	}
	head, ok := pageProps["head"].(map[string]interface{})
	if !ok {
		return internal.JobDescription{}, errors.New("head object not found")
	}
	jdId := strings.Split(url, "/")[4]
	id, ok := head[jdId].(map[string]interface{})
	if !ok {
		return internal.JobDescription{}, fmt.Errorf("id object not found: %s", jdId)
	}
	jdContents, ok := id["jd"].(string)
	if !ok {
		return internal.JobDescription{}, errors.New("jd object not found")
	}

	var step uint8
	var jd internal.JobDescription
	lines := strings.Split(jdContents, "\n")
	for _, line := range lines {
		switch line {
		case "주요업무", "자격요건", "우대사항", "혜택 및 복지":
			step++
		case "":
			continue
		default:
			switch step {
			case 1:
				jd.MainTasks = append(jd.MainTasks, line)
			case 2:
				jd.Required = append(jd.Required, line)
			case 3:
				jd.Preferred = append(jd.Preferred, line)
			case 4:
				jd.Benefits = append(jd.Benefits, line)
			}
		}
	}

	return jd, nil
}

func findNextDataScript(node *html.Node) string {
	if node.Type == html.ElementNode && node.Data == "script" {
		for _, attr := range node.Attr {
			if attr.Key == "id" && attr.Val == "__NEXT_DATA__" {
				if node.FirstChild != nil {
					return node.FirstChild.Data
				}
			}
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if data := findNextDataScript(c); data != "" {
			return data
		}
	}

	return ""
}

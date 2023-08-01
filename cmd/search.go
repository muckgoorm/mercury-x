package cmd

import (
	"fmt"
	"mercury-x/internal"
	"mercury-x/internal/wanted"
	"mercury-x/pkg/survey"
	"mercury-x/pkg/tui"
	"strconv"
)

func search() {
	platform := survey.SingleChoice("검색할 플랫폼을 골라주세요", internal.Platforms)
	role := survey.SingleChoice("직군을 골라주세요", internal.Roles)
	experience := survey.SingleChoice("경력을 골라주세요", internal.Experiences)
	stacks := survey.MultipleChoice("기술 스택을 골라주세요", internal.Stacks, 5)
	benefits := survey.MultipleChoice("복지 여건을 선택해주세요", internal.Benefits, 3)
	count, _ := strconv.Atoi(survey.SingleChoice("보고 싶은 채용 공고의 개수를 입력해주세요", internal.Counts))

	payload := internal.JobSearchPayload{
		Role:       role,
		Experience: experience,
		Stacks:     stacks,
		Benefits:   benefits,
		Count:      count,
	}

	var postings []internal.JobPosting

	switch platform {
	case "wanted":
		wc := wanted.NewWantedCrawler()
		postings, _ = wc.SearchJobPostings(payload)
	}

	firstTitle := "채용 회사 / 직무"
	secondTitle := "JD"
	var items []string
	for _, p := range postings {
		items = append(items, fmt.Sprintf("%s / %s", p.Company, p.Role))
	}
	tui.GenerateFlexTable(firstTitle, secondTitle, items)
}

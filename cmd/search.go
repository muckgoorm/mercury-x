package cmd

import (
	"fmt"
	"mercury-x/internal"
	"mercury-x/internal/wanted"
	"mercury-x/pkg/survey"
	"mercury-x/pkg/tui"

	"github.com/urfave/cli/v2"
)

func searchCommand() *cli.Command {
	return &cli.Command{
		Name:        "search",
		Usage:       "채용 공고 검색",
		Description: "원하는 조건에 맞는 채용 공고들을 조회합니다.",
		Aliases:     []string{"s"},
		Action:      searchFunction,
	}
}

func searchFunction(c *cli.Context) error {
	platform := survey.SingleChoice("검색할 플랫폼을 골라주세요", internal.Platforms)
	role := survey.SingleChoice("직군을 골라주세요", internal.Roles)
	experience := survey.SingleChoice("경력을 골라주세요", internal.Experiences)
	stacks := survey.MultipleChoice("기술 스택을 골라주세요", internal.Stacks)
	benefits := survey.MultipleChoice("복지 여건을 선택해주세요", internal.Benefits)

	payload := internal.JobSearchPayload{
		Role:       role,
		Experience: experience,
		Stacks:     stacks,
		Benefits:   benefits,
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

	return nil
}

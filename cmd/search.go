package cmd

import (
	"fmt"
	"mercury-x/internal"
	"mercury-x/internal/crawler"
	"mercury-x/pkg"

	"github.com/urfave/cli/v2"
)

func searchFunction(c *cli.Context) error {
	platform := pkg.SingleChoice("검색할 플랫폼을 골라주세요", internal.Platforms)
	role := pkg.SingleChoice("직군을 골라주세요", internal.Roles)
	experience := pkg.SingleChoice("경력을 골라주세요", internal.Experiences)
	stacks := pkg.MultipleChoice("기술 스택을 골라주세요", internal.Stacks)
	benefits := pkg.MultipleChoice("복지 여건을 선택해주세요", internal.Benefits)

	payload := crawler.JobSearchPayload{
		Platform:   platform,
		Role:       role,
		Experience: experience,
		Stacks:     stacks,
		Benefits:   benefits,
	}

	fmt.Println(&payload)

	pkg.JobList()

	return nil
}

func searchCommand() *cli.Command {
	return &cli.Command{
		Name:        "search",
		Usage:       "채용 공고 검색",
		Description: "원하는 조건에 맞는 채용 공고들을 조회합니다.",
		Aliases:     []string{"s"},
		Action:      searchFunction,
	}
}

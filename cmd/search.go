package cmd

import (
	"fmt"
	"mercury-x/internal"

	"github.com/urfave/cli/v2"
)

func searchFunction(c *cli.Context) error {
	platform := internal.SingleChoice("검색할 플랫폼을 골라주세요", internal.Platforms)
	role := internal.SingleChoice("직군을 골라주세요", internal.Roles)
	experience := internal.SingleChoice("경력을 골라주세요", internal.Experiences)
	stacks := internal.MultipleChoice("기술 스택을 골라주세요", internal.Stacks)
	benefits := internal.MultipleChoice("복지 여건을 선택해주세요", internal.Benefits)

	fmt.Println(platform)
	fmt.Println(role)
	fmt.Println(experience)
	fmt.Println(stacks)
	fmt.Println(benefits)

	internal.JobList()

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

package cmd

import (
	"fmt"
	"mercury-x/pkg"
	"strings"

	"github.com/urfave/cli/v2"
)

func searchFunction(c *cli.Context) error {
	platform := pkg.MultipleChoice(
		"검색할 플랫폼을 골라주세요",
		[]string{"wanted", "rallit"},
	)

	if len(platform) == 0 {
		fmt.Println("플랫폼을 선택해주세요.")
		return nil
	}
	platforms := strings.Join(platform, ", ")

	experience := pkg.SingleChoice(
		"경력을 골라주세요",
		[]string{"신입", "1 년차", "2 년차", "3 년차", "4 년차", "5 년차", "6 년차", "7 년차", "8 년차", "9 년차", "10 년차 이상"},
	)

	fmt.Println(platforms)
	fmt.Println(experience)

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

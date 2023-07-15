package cmd

import "github.com/urfave/cli/v2"

func searchFunction(c *cli.Context) error {
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

package cmd

import "github.com/urfave/cli/v2"

func analyzeFunction(c *cli.Context) error {
	return nil
}

func analyzeCommand() *cli.Command {
	return &cli.Command{
		Name:        "analyze",
		Usage:       "채용 공고 분석",
		Description: "내 직무 역량에 맞는 채용 공고들이 어떤 세부적인 JD를 요구하는지 분석합니다.",
		Aliases:     []string{"a"},
		Action:      analyzeFunction,
	}
}

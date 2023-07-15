package cmd

import "github.com/urfave/cli/v2"

var commands []*cli.Command = []*cli.Command{
	checkCommand(),
	searchCommand(),
	analyzeCommand(),
}

func GetCommands() []*cli.Command {
	return commands
}

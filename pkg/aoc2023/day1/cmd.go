package day1

import (
	_ "embed"

	"github.com/spf13/cobra"
)

var (
	//go:embed sample.txt
	sample string
	//go:embed input.txt
	input string
	test  bool
)

func init() {
	Cmd.AddCommand(
		part1Cmd,
		part2Cmd,
	)
	Cmd.PersistentFlags().BoolVar(&test, "test", false, "test")
}

var Cmd = &cobra.Command{
	Use:   "day1",
	Short: "",
}

func getInput() string {
	if test {
		return sample
	}
	return input
}

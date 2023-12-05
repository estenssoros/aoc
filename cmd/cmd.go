package cmd

import (
	"github.com/estenssoros/aoc/pkg/aoc2023"
	"github.com/spf13/cobra"
)

func init() {
	cmd.AddCommand(
		aoc2023.Cmd,
	)
}

var cmd = &cobra.Command{
	Use:   "aoc",
	Short: "",
}

func Execute() error {
	return cmd.Execute()
}

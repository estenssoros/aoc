package aoc2023

import (
	"github.com/estenssoros/aoc/pkg/aoc2023/day1"
	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(
		day1.Cmd,
	)
}

var Cmd = &cobra.Command{
	Use:   "2023",
	Short: "",
}
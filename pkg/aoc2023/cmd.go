package aoc2023

import (
	"github.com/estenssoros/aoc/pkg/aoc2023/day1"
	"github.com/estenssoros/aoc/pkg/aoc2023/day2"
	"github.com/estenssoros/aoc/pkg/aoc2023/day3"
	"github.com/estenssoros/aoc/pkg/aoc2023/day4"
	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(
		day1.Cmd,
		day2.Cmd,
		day3.Cmd,
		day4.Cmd,
	)
}

var Cmd = &cobra.Command{
	Use:   "2023",
	Short: "",
}

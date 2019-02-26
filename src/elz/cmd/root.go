package cmd

import (
	"github.com/elz-lang/elz/src/elz/cmd/compile"

	"github.com/spf13/cobra"
)

var (
	Compiler = &cobra.Command{
		Use:   "elz",
		Short: "elz compiler",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
)

func init() {
	Compiler.AddCommand(compile.Cmd)
}

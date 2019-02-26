package compile

import (
	"fmt"

	"github.com/elz-lang/elz/src/elz/builder"
	"github.com/elz-lang/elz/src/elz/codegen"

	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:   "compile",
		Short: "compile input file",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("compile expect a file input")
				return
			}
			filename := args[0]
			builder := builder.New()
			err := builder.BuildFromFile(filename)
			if err != nil {
				fmt.Printf("failed at compile file, error: %s\n", err)
				return
			}
			bindMap := builder.GetBindMap()
			bindType := builder.GetBindTypes()
			g := codegen.New(bindMap, bindType)
			g.Generate()
			fmt.Printf("%s", g)
		},
	}
)

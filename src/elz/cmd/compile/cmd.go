package compile

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/elz-lang/elz/src/elz/builder"
	"github.com/elz-lang/elz/src/elz/codegen"

	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:   "compile",
		Short: "compile input file",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("compile expect a file input")
			}
			entryFile := args[0]
			rootDirOfProject := filepath.Dir(entryFile)
			c := newCacheAgent()
			err := c.compile(rootDirOfProject, entryFile)
			if err != nil {
				return fmt.Errorf("failed at compile file, error: %s\n", err)
			}
			entry := c.caches["main"]
			// make sure won't reference to entry module in codegen
			c.caches["main"] = nil
			g := codegen.New(entry, c.caches)
			g.Generate()
			fmt.Printf("%s", g)
			return nil
		},
	}
)

type cacheAgent struct {
	caches map[string]*codegen.Tree
}

func newCacheAgent() *cacheAgent {
	return &cacheAgent{
		caches: make(map[string]*codegen.Tree),
	}
}

func (c *cacheAgent) compile(root, file string) error {
	tree, err := builder.NewFromFile(file)
	if err != nil {
		return err
	}
	for _, dep := range tree.GetDependencies() {
		err := c.compile(root, filepath.Join(root, dep))
		if err != nil {
			return err
		}
	}
	k := strings.TrimPrefix(file, "./")
	k = strings.TrimPrefix(k, root+"/")
	k = strings.TrimSuffix(k, ".elz")
	c.caches[k] = tree
	return nil
}

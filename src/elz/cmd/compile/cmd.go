package compile

import (
	"fmt"
	"os"
	"os/exec"
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
			entry, err := c.compile(rootDirOfProject, entryFile)
			if err != nil {
				return fmt.Errorf("failed at compile file, error: %s\n", err)
			}
			g := codegen.New(entry, c.caches)
			g.Generate()
			tmpIRFile, err := os.OpenFile(
				// main_file_name.ll
				fmt.Sprintf("%s.ll", entryFile),
				// if not exist create, truncate from start if exist, read/write open
				os.O_CREATE|os.O_TRUNC|os.O_RDWR,
				0644,
			)
			if err != nil {
				return err
			}
			defer tmpIRFile.Close()
			_, err = tmpIRFile.WriteString(g.String())
			if err != nil {
				return err
			}
			defer os.Remove(tmpIRFile.Name())

			irFileName := tmpIRFile.Name()
			irFilePrefix := strings.TrimSuffix(irFileName, "ll")
			tmpObject := irFilePrefix + "o"
			compileCmd := exec.Command("clang", "-c", tmpIRFile.Name(), "-o", tmpObject)
			if err := compileCmd.Run(); err != nil {
				return fmt.Errorf("failed at compile ir file: %s", err)
			}
			homeDir, err := os.UserHomeDir()
			if err != nil {
				return err
			}
			linkCmd := exec.Command("clang", tmpObject,
				filepath.Join(homeDir, ".elz/core/list.o"),
				filepath.Join(homeDir, ".elz/core/memory.o"),
			)
			defer os.Remove(tmpObject)
			if err := linkCmd.Run(); err != nil {
				return fmt.Errorf("failed at link object file: %s", err)
			}

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

func (c *cacheAgent) compile(rootDir, entryFile string) (*codegen.Tree, error) {
	tree, err := builder.NewFromFile(entryFile)
	if err != nil {
		return nil, err
	}
	for _, importPath := range tree.GetDependencies() {
		// importPath: lib::sub_lib::sub_lib
		dependentFilePath := filepath.Join(strings.Split(importPath, "::")...)
		tree, err := c.compile(rootDir, filepath.Join(rootDir, dependentFilePath)+".elz")
		if err != nil {
			return nil, err
		}
		c.addTree(importPath, tree)
	}
	return tree, nil
}

func (c *cacheAgent) addTree(importPath string, tree *codegen.Tree) {
	if _, exist := c.caches[importPath]; exist {
		return
	}
	c.caches[importPath] = tree
}

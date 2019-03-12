package builder

import (
	"github.com/elz-lang/elz/src/elz/parser"
)

func (b *Builder) ExitImportStatement(c *parser.ImportStatementContext) {
	b.astTree.InsertImport(c.AccessChain().GetText())
}

package listener

import (
	"github.com/elz-lang/elz/parser"
)

// EnterExportor: + prefix
func (s *ElzListener) EnterExportor(*parser.ExportorContext) {
	s.exportThis = true
}

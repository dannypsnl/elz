package listener

import (
	"github.com/elz-lang/elz/parser"
)

// EnterExporter listen
//
// grammar:
//   + prefix
func (s *ElzListener) EnterExporter(*parser.ExporterContext) {
	s.exportThis = true
}

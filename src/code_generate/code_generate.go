package main

import "C"
import (
	"github.com/elz-lang/elz/src/mir"

	"github.com/gogo/protobuf/proto"
)

//export generate
func generate(program *C.char) {
	m := &mir.MIR{}
	err := proto.Unmarshal([]byte(C.GoString(program)), m)
	if err != nil {
		panic(err)
	}

	generateMIR(m)
}

func generateMIR(m *mir.MIR) {
}

func main() {}

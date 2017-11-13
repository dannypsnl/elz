package lexer

import (
	"fmt"
)

type ItemType int

const (
	ItemError ItemType = iota

	ItemIdentifier // ex. name
	ItemNumber     // ex. 1.24 0.12 .23
	ItemString     // ex. "I am a string"

	// Here are keywords
	ItemKwMatch  // match xxx { ... }
	ItemKwLoop   // loop { statement }
	ItemKwBreak  // break
	ItemKwReturn // return expression
	ItemKwLet    // let a = 1
	ItemKwMut    // let mut a = 1
	ItemKwFn     // fn add(lv: num, rv: num) -> num {return lv+rv}
	ItemKwType   // type
	ItemKwImpl   // impl type { methods }
	ItemKwTrait  // trait xxx { methodTypes }
	ItemKwAct    // act functionCall

	// Here are operator
	ItemLeftParam  // (
	ItemRightParam // )
	ItemAdd        // +
	ItemMinus      // -
	ItemMul        // *
	ItemDiv        // /
	ItemDot        // .

	ItemEOF
)

type Item struct {
	Type ItemType
	Val  string
}

func (i Item) String() string {
	switch i.Type {
	case ItemEOF:
		return "EOF"
	case ItemError:
		return i.Val
	}
	if len(i.Val) > 10 {
		return fmt.Sprintf("%.10q...", i.Val)
	}
	return fmt.Sprintf("%q", i.Val)
}

type Lexer struct{}

type stateFn func(l *Lexer) stateFn

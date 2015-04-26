package main

import "fmt"

// item represents a token or text string returned from the scanner.
type action struct {
	typ actionType // The type of this item.
	//pos Pos      // The starting position, in bytes, of this item in the input string.
	val string // The value of this item.
}

// stateFn represents the state of the game as a function that returns the next state.
type stateFunc func(*game) stateFunc

// lexer holds the state of the scanner.
type game struct {
	name  string // the name of the input; used only for error reports
	input string // the string being scanned
	//leftDelim  string    // start of action
	//rightDelim string    // end of action
	state stateFunc // the next action function to enter
	//pos        Pos       // current position in the input
	//start      Pos       // start position of this action
	//width      Pos       // width of last rune read from input
	//lastPos    Pos       // position of most recent action returned by nextItem
	actions chan action // channel of scanned actions
	//parenDepth int         // nesting depth of ( ) exprs
}

func main() {
	fmt.Println("here...")
}

func (i action) String() string {
	switch {
	case i.typ == itemEOF:
		return "EOF"
	case i.typ == itemError:
		return i.val
	case i.typ > itemKeyword:
		return fmt.Sprintf("<%s>", i.val)
	case len(i.val) > 10:
		return fmt.Sprintf("%.10q...", i.val)
	}
	return fmt.Sprintf("%q", i.val)
}

// actionType identifies the type of lex items.
type actionType int

const (
	itemError        actionType = iota // error occurred; value is text of error
	itemBool                           // boolean constant
	itemChar                           // printable ASCII character; grab bag for comma etc.
	itemCharConstant                   // character constant
	itemComplex                        // complex constant (1+2i); imaginary is just a number
	itemColonEquals                    // colon-equals (':=') introducing a declaration
	itemEOF
	itemField      // alphanumeric identifier starting with '.'
	itemIdentifier // alphanumeric identifier not starting with '.'
	itemLeftDelim  // left action delimiter
	itemLeftParen  // '(' inside action
	itemNumber     // simple number, including imaginary
	itemPipe       // pipe symbol
	itemRawString  // raw quoted string (includes quotes)
	itemRightDelim // right action delimiter
	itemRightParen // ')' inside action
	itemSpace      // run of spaces separating arguments
	itemString     // quoted string (includes quotes)
	itemText       // plain text
	itemVariable   // variable starting with '$', such as '$' or  '$1' or '$hello'
	// Keywords appear after all the rest.
	itemKeyword  // used only to delimit the keywords
	itemDot      // the cursor, spelled '.'
	itemDefine   // define keyword
	itemElse     // else keyword
	itemEnd      // end keyword
	itemIf       // if keyword
	itemNil      // the untyped nil constant, easiest to treat as a keyword
	itemRange    // range keyword
	itemTemplate // template keyword
	itemWith     // with keyword
)

package step

import (
	"bufio"
	"flag"
	"io"
	"os"
	"strings"
)

// Lexer : object for extracting symbols(lexemes)
// from the file in order for them to be added to
// the parse graph by the Parser
type Lexer struct {
	Filename    string
	SymbolsChan chan Symbol
}

// Symbol : async lexing tool
// step file info related to a
// specific symbol
type Symbol struct {
	ID  string
	Def Definition
}

// SymbolTable : sync lexing tool
// step file info related to a
// specific symbol
type SymbolTable map[string]Definition

// Definition : step file definition
// related to a specific identifier.
type Definition struct {
	Type      string
	Arguments []string
}

// define lexing state
const (
	lexing = true
	done   = false
)

func initLexer(filename string) Lexer {
	return Lexer{
		Filename:    filename,
		SymbolsChan: make(chan Symbol, 2),
	}
}

// Lex : parses the step file and fills the symbol
// table for the parser.
func (l *Lexer) Lex() SymbolTable {
	flag.Parse()
	file, err := os.Open(l.Filename)

	check(err)
	defer file.Close()
	reader := bufio.NewReader(file)

	/* symbolData is the data related to the parsed symbol
	---got from the step file */
	symbolText := ""
	symbolTable := make(SymbolTable, 0)
	/* DEBUG : there is a bug with the fact that you didn't the case of that there is only one line in the definition text */
	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		buffer := string(line)

		/* DEBUG BREAKPOINT */
		//fmt.Println("buffer:", buffer, "\n buffer Length:", len(buffer))

		/* skippin empty lines*/
		/*
			make it more robust to support multilines /could be the cause of a bug/
		*/
		if len(buffer) == 0 {
			continue // empty
		}
		/* __--------------- */
		/*	you can optimize this further more by making sure the symbol
			is init to the first value when you start parsing and remvove
			the first check*/
		if strings.Contains(buffer, "=") { // contains =
			/* --sync method
			---fmt.Println("emit") */

			//fmt.Println(buffer, "buffer")
			/*  the symbol data gets filled by the whole step description
			----before getting split into identifier and definition */
			if len(symbolText) > 0 {
				symbolTable.add(symbolText)
			}
			/* ---------------------------------------------------- */
			/*
				---async method
				symbol := Symbol{ID: identifier, Def: definition}
				l.emit(symbol)
			*/
			/*_----------------------------------------------------_*/
			symbolText = buffer // is not empty and does not contain =
		} else {
			symbolText += buffer
		}
	}
	symbolTable.add(symbolText)
	return symbolTable
}

// emit : emits the symbol to the symbols chanel
func (l *Lexer) emit(s Symbol) {
	l.SymbolsChan <- s
}

/* add : adds a symbol to the symbol table from symbol definition text
---got from the step file
*/
func (st *SymbolTable) add(symbolText string) {
	split := strings.SplitN(symbolText, "=", 2)
	identifier := split[0]
	definition := tokenize(split[1])
	(*st)[identifier] = definition

}

// tokenize : gets the definition of an identifier
func tokenize(definitionText string) Definition {
	split := strings.SplitN(definitionText, "(", 2)

	argumentsText := split[1]
	definitionArguments := getArguments(argumentsText)
	definitionType := stripWhiteSpaces(split[0])
	return Definition{
		Type:      definitionType,
		Arguments: definitionArguments,
	}
}

package step

import "testing"

func TestParseUndefined(t *testing.T) {
	filename := "testfiles/in/undefined.stp"
	lexer := initLexer(filename)
	symbolsTable := lexer.Lex()
	Parse(symbolsTable)
}

func TestParseSurface(t *testing.T) {
	filename := "testfiles/in/surface.stp"
	lexer := initLexer(filename)
	symbolsTable := lexer.Lex()
	Parse(symbolsTable)
}
func TestParseBunny(t *testing.T) {
	filename := "testfiles/in/bunny.stp"
	lexer := initLexer(filename)
	symbolsTable := lexer.Lex()
	Parse(symbolsTable)
}

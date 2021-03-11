package step

import (
	"fmt"
	"testing"
)

func TestLex(t *testing.T) {
	filename := "testfiles/in/undefined.stp"
	lexer := initLexer(filename)
	symbolsTable := lexer.Lex()
	for ID, Definition := range symbolsTable {
		fmt.Println("identifier:", ID)
		fmt.Println("Definition:")
		fmt.Println("----Type:", Definition.Type)
		fmt.Println("----Arguments:", Definition.Arguments)

	}

	/* Marshalling to Json */
	/* bytes, err := json.Marshal(symbolsTable)
	if err != nil {
		fmt.Println("can't serialize")
	}
	file, err := os.Create("testfiles/out/test.json")
	check(err)
	defer file.Close()
	file.Write(bytes)
	*/
	/* __----------------- */
}

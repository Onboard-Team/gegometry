package step

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"unicode"
)

// getArguments : gets arguments from arguments STEP file text
func getArguments(argumentsText string) []string {
	arguments := make([]string, 0)
	openBrackets := 0
	argument := make([]rune, 0)
	for _, r := range argumentsText {
		if r == '(' {
			openBrackets++
		}
		if r == ')' {
			openBrackets--
		}
		if openBrackets == 0 && r == ',' {
			arguments = append(arguments, string(argument))
			argument = make([]rune, 0)
			continue
		}
		argument = append(argument, r)
	}
	/* sanitizing last input by remove `);` */
	lastArgument := string(argument)[:len(argument)-2]
	arguments = append(arguments, lastArgument)
	return arguments
}

// getArray : gets an array from arguments refactor
/* TODO: refactor this to be able to return an
---[]interface instead of just a string array
---try casting using reflect to cast and not create a huge
---switch yourself.
*/
/* TODO:handle ERRORS */
func getArray(arrayText string, t reflect.Kind) []string {
	/* removing `(` from start and `)` from end*/
	arrayText = arrayText[1 : len(arrayText)-1]
	openBrackets := 0
	element := make([]rune, 0)
	array := make([]string, 0)
	for _, r := range arrayText {
		if r == '(' {
			openBrackets++
		}
		if r == ')' {
			openBrackets--
		}
		if openBrackets == 0 && r == ',' {
			array = append(array, string(element))
			element = make([]rune, 0)
			continue
		}
		element = append(element, r)
	}
	array = append(array, string(element))
	return array
}

/* TODO:handle ERRORS */
func getMatrix(matrixText string) [][]string {
	cols := getArray(matrixText, reflect.Array)
	matrix := make([][]string, 0)
	for _, col := range cols {
		row := getArray(col, reflect.Float64)
		matrix = append(matrix, row)
	}
	return matrix
}

/* TODO:handle ERRORS */
func getLogical(logicText string) bool {
	switch logicText {
	case ".T.":
		return true
	case ".F.":
		return false
	default:
		return false
	}
}

func stripWhiteSpaces(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

/* TODO : check if the interface passed as a param works fine */
func marshallToJSON(item interface{}) {
	bytes, err := json.Marshal(item)
	if err != nil {
		fmt.Println("can't serialize")
	}
	file, err := os.Create("testfiles/out/test.json")
	check(err)
	defer file.Close()
	file.Write(bytes)

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

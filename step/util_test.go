package step

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetArray(t *testing.T) {
	testString := "(bla,(bla,bla,(bla,bla)),bla,(bla,bla),bla)"
	for _, argument := range getArray(testString, reflect.String) {
		fmt.Println(argument)
	}
}

package step

/* In order to make the parser work properly
---in a concurrent fashion we'll need to add
---a skip stack that you'll get back to when
---some of the points are missing
---something like:
-Skip: list of definitions that are impossible to parse
for now
-Missing: list of definitions that are missing
--complexity of map lookups: log(n) `they are based on red black trees`
--complexity of bitarray is o(1) but we require way more way storage space
for big files
--the big problem is the fact that we get threads blocked so if they are
they'll just have to push the blocking definition to the skip/missing stacks.
--we also need to create a thread pool to be able to manage cpu usage.
*/

/*  step files seen as state machines:
----parsing could be seen as transitions through a deterministic
----state machine.
----our parser uses this terminology to parse the symbols table in
----a very elegant way.
*/

/*	parseFn : implements the transition states for parsing
----components from the table of symbols to create a render
----stack to be excecuted by the converter.
----TODO: make it return a set of parseFn
*/
type parseFn func(string)

/* parseState : describes the current state and the
---transition function
*/
type parseState struct {
	fn  parseFn // state
	key string  // transition
}

// Parser : the state machine
type Parser struct {
	lexemes    SymbolTable
	nextStates []parseState
	visited    map[string]struct{}
}

/* TODO: make it work recursively
---1-march-2021: it works recursively by adding it to the parser's nextStates
-----------------VERTEX_POINT:(save this state)
------------------------------CARTESIAN_POINT
------------------------------recover the state
*/
func (parser *Parser) runInferenceModel(state parseState) {
	parser.nextStates = append(parser.nextStates, state)

	for len(parser.nextStates) > 0 {
		popState := parser.nextStates[0]
		if _, found := parser.visited[popState.key]; !found {

			popState.fn(popState.key)
			parser.visited[popState.key] = struct{}{}
		}
		parser.nextStates = parser.nextStates[1:]

	}
}

//Parse : state machine inference engine to retrieve the
//render stack from the lexer
func Parse(lexemes SymbolTable) {
	parser := Parser{lexemes: lexemes}
	parser.visited = make(map[string]struct{}, 0)

	for key, lexeme := range lexemes {
		fn := parser.bridge(lexeme)
		ps := parseState{fn: fn, key: key}
		parser.runInferenceModel(ps)
	}
}

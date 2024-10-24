package syntax

type Token struct {
	Type string
	Line int
	Literal string
}

func Tokenize(source string) (arr []Token) {
	lineNum := 1

	for _, c := range source {
		switch c {
		case '(':
			token(&arr, "LEFT_PAREN", lineNum, string(c))
		case ')':	
			token(&arr, "RIGHT_PAREN", lineNum, string(c))
		case '=':
			token(&arr, "EQUAL", lineNum, string(c))
		case '\n':
			lineNum++
		default:
		}
	}

	return
}

func token(arr *[]Token, t string, line int, literal string) {
	*arr = append(*arr, Token{Type: t, Line: line, Literal: literal})
}
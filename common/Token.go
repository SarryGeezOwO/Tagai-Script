package common

type Token struct {
	Type string
	Lexeme string
	Line int
	Literal interface{} // Any type
}
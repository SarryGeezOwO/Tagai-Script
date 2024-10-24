package syntax

import (
	er "tagai-script/error"
	com "tagai-script/common"
	"unicode"
	"strings"
	"strconv"
)

var start int = 0
var current int = 0
var line int = 1
var source string = ""
var tokenList []com.Token

// Initialize a set of keywords using a map
var keywords = map[string]bool{
    "AND":    true,
    "CLASS":  true,
    "ELSE":   true,
    "FALSE":  true,
    "TRUE":   true,
    "IF":     true,
    "FOR":    true,
    "FUN":    true,
    "NIL":    true,
    "OR":     true,
    "PRINT":  true,
    "RETURN": true,
    "THIS":   true,
    "SUPER":  true,
    "VAR":    true,
    "WHILE":  true,
}

func Tokenize(src string) []com.Token {
	source = src

	for !isEndFile() {
		start = current
		scanToken()
	}

	tokenList = append(tokenList, com.Token{"EOF", "", line, nil})
	return tokenList
}

func scanToken() {
	c := advance()

	switch c {
	case '(': addToken("LEFT_PAREN")
	case ')': addToken("RIGHT_PAREN")
	case '{': addToken("LEFT_BRACE")
	case '}': addToken("RIGHT_BRACE")
	case ',': addToken("COMMA")
	case '.': addToken("DOT")
	case '-': addToken("MINUS")
	case '+': addToken("PLUS")
	case ';': addToken("SEMICOLON")
	case '*': addToken("STAR")
	case '!':
		if match('=') {
				addToken("BANG_EQUAL")
		}else { addToken("BANG")}
	case '=':
		if match('=') {
				addToken("EQUAL_EQUAL")
		}else { addToken("EQUAL")}
	case '>':
		if match('=') {
				addToken("GREATER_EQUAL")
		}else { addToken("GREATER")}
	case '<':
		if match('=') { 
				addToken("LESS_EQUAL")
		}else { addToken("LESS") }
	case '/':
		if match('/') {
			for peek() != '\n' && !isEndFile() {
				advance()
			}
		}else { addToken("SLASH") }
	case ' ':
	case '\t':
	case '\r':
	case '\n': line++
	case '"' : initString()
	default: 
		if unicode.IsDigit(c) {
			number()
		} else if unicode.IsLetter(c) {
			identifier()
		} else {
			er.ErrorLine(line, "Unexpected Character {"+ string(c) + "}.")
		}
	}
}


func initString() {
	for peek() != '"' && !isEndFile() {
		if peek() == '\n' {
			line++
		}
		advance()
	}

	if isEndFile() {
		er.ErrorLine(line, "unterminated string.")
		return
	}

	advance()
	value := source[start+1 : current-1]
	addTokenLit("STRING", value)
}


func identifier() {
	for isAlphaNumeric(peek()) {
		advance()
	}

	text := source[start : current]
	tokenType := strings.ToUpper(text)

	if !validateKeyword(tokenType) {
		tokenType = "IDENTIFIER"
	}
	addToken(tokenType)
}


func number() {
	for unicode.IsDigit(peek()) {
		advance()
	}

	if peek() == '.' && unicode.IsDigit(peekNext()) {
		advance()
		for unicode.IsDigit(peek()) {
			advance()
		}
	}

	float64Value, err := strconv.ParseFloat(source[start : current], 32)
    if err != nil {
        er.ErrorLine(line, "Error string conversion to float")
        return
    }
    float32Value := float32(float64Value)
	addTokenLit("NUMBER", float32Value)
}	


func match(expected rune) bool {
	if isEndFile() {
		return false
	}

	if rune(source[current]) != expected {
		return false
	}

	current++
	return true
}



func peek() rune {
	if isEndFile() {
		return ' '
	}
	return rune(source[current])
}



func peekNext() rune {
	if current + 1 >= len(source) {
		return ' '
	}
	return rune(source[current + 1])
}



func advance() rune {
	current++
	return rune(source[current-1])
}


func isAlphaNumeric(r rune) bool {
	return unicode.IsDigit(r) || unicode.IsLetter(r)
}

func isEndFile() bool {
	return current >= len(source)
}

func addToken(t string) {
	addTokenLit(t, nil)
}

func addTokenLit(t string, literal interface{}) {
	text := source[start:current]
	tokenList = append(tokenList, com.Token{t, text, line, literal})
}


// Returns true if such keyword does exists in the language
func validateKeyword(text string) bool {
    return keywords[text]
}
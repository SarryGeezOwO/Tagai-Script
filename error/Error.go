package error

import (
	"fmt"
	com "tagai-script/common"
)

var ErrorPresent bool = false
var RuntimeErrorPresent bool = false

func ErrorLine(line int, msg string) {
	report(line, "", msg)
}


func Error(token com.Token, msg string) {
	if token.Type == "EOF" {
		report(token.Line, "at end", msg)
	}else {
		report(token.Line, "at '"+token.Lexeme+"'", msg)
	}
}


func report(line int, where string, msg string) {
	fmt.Printf("[Line %d] Error %s: %s\n", line, where, msg)
	ErrorPresent = true;
}
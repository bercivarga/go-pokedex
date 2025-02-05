package repl

import (
	"bufio"
	"os"
	"strings"
)

func CleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func CreateTerminalScanner() *bufio.Scanner {
	return bufio.NewScanner(os.Stdin)
}

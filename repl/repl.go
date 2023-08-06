// repl/repl.go

package repl

import (
	"bufio"
	"fmt"
	"io"

	"seahorsesyntax.villdyr.dev/lexer"
	"seahorsesyntax.villdyr.dev/token"
)

// hello test
const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan() // Read a line of input from the user
		if !scanned {
			return
		}

		line := scanner.Text() // Get the text of the line
		l := lexer.New(line)   // Create a lexer for the line

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok) // Print the tokens
		}

	}
}

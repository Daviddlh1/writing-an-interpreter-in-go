package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Daviddlh1/writing-an-interpreter-in-go/pkg/lexer"
	"github.com/Daviddlh1/writing-an-interpreter-in-go/pkg/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		out.Write([]byte(PROMPT))

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			out.Write([]byte(fmt.Sprintf("%+v\n", tok)))
		}
	}
}

package shell

import (
	"io"
	"log"

	"github.com/chzyer/readline"
)

func newReader(in io.ReadCloser, out io.Writer) *readline.Instance {
	reader, err := readline.NewEx(&readline.Config{
		Prompt:            blue(prompt),
		HistorySearchFold: true,
		InterruptPrompt:   "^C",
		EOFPrompt:         "Good Bye!",
		Stdin:             in,
		Stdout:            out,
		FuncFilterInputRune: func(r rune) (rune, bool) {
			if r == readline.CharCtrlZ {
				return r, false
			}
			return r, true
		},
	})
	if err != nil {
		log.Fatalln(err)
	}
	return reader
}

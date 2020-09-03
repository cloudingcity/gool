package shell

import (
	"io"
	"log"
	"path/filepath"

	"github.com/chzyer/readline"
)

func newReader(in io.ReadCloser, out io.Writer, historyPath string, completer readline.AutoCompleter) *readline.Instance {
	var historyFile string
	if historyPath != "" {
		historyFile = filepath.Join(historyPath, "gool.tmp")
	}
	reader, err := readline.NewEx(&readline.Config{
		Prompt:            blue(prompt),
		HistoryFile:       historyFile,
		HistorySearchFold: true,
		AutoComplete:      completer,
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

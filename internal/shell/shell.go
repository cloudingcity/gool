package shell

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/chzyer/readline"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const prompt = "=# "

// color definitions
var (
	cyan   = color.RGB(1, 173, 216).SprintfFunc()
	blue   = color.New(color.FgHiBlue).SprintfFunc()
	yellow = color.New(color.FgYellow).SprintfFunc()
	white  = color.RGB(255, 255, 255).SprintfFunc()
)

type cmd struct {
	name string
	desc string
}

var basicCmds = []cmd{
	{name: "/help", desc: "Show help"},
	{name: "/clear", desc: "Clean the screen"},
	{name: "/quit", desc: "Quit the REPL"},
}

type Shell struct {
	in          io.ReadCloser
	out         io.Writer
	reader      *readline.Instance
	historyPath string
	current     string
	mCmds       map[string]*cobra.Command
	cmds        []cmd
	quit        bool
}

func New(in io.ReadCloser, out io.Writer) *Shell {
	return &Shell{
		in:    in,
		out:   out,
		mCmds: make(map[string]*cobra.Command),
	}
}

func (s *Shell) Register(cobraCmd ...*cobra.Command) {
	for _, c := range cobraCmd {
		s.mCmds[c.Name()] = c
		cmd := cmd{
			name: "/" + c.Name(),
			desc: c.Short,
		}
		s.cmds = append(s.cmds, cmd)
	}
}

func (s *Shell) SetHistoryPath(path string) {
	s.historyPath = path
}

func (s *Shell) Run() {
	s.reader = newReader(s.in, s.out, s.historyPath, s.autoCompleter())
	defer s.reader.Close()

	s.welcome()

	for {
		line, err := s.reader.Readline()
		if errors.Is(err, readline.ErrInterrupt) {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		input := strings.TrimSpace(line)
		s.exec(input)

		if s.quit {
			s.println("Good Bye!")
			break
		}
	}
}

func (s *Shell) welcome() {
	s.println(cyan(`
 ██████┐  ██████┐  ██████┐ ██┐     
██┌────┘ ██┌───██┐██┌───██┐██│     
██│  ███┐██│   ██│██│   ██│██│     
██│   ██│██│   ██│██│   ██│██│     
└██████┌┘└██████┌┘└██████┌┘███████┐
 └─────┘  └─────┘  └─────┘ └──────┘
`))
	s.printf(" %s %s\n\n", white("/help"), "for more information")
}

func (s *Shell) exec(input string) {
	if strings.HasPrefix(input, "/") {
		cmd := strings.TrimPrefix(input, "/")
		s.execCommand(cmd)
		return
	}
	s.execCommandWith(input)
}

func (s *Shell) execCommand(cmd string) {
	switch cmd {
	case "help":
		s.helpCmd()
	case "clear":
		s.cleanCmd()
	case "quit":
		s.quitCmd()
	default:
		if _, ok := s.mCmds[cmd]; ok {
			s.switchCmd(cmd)
		} else {
			s.execCommandWith(cmd)
		}
	}
}

func (s *Shell) helpCmd() {
	s.println()
	s.println("Basics:")
	for _, cmd := range basicCmds {
		s.printf(" %s - %s\n", white(cmd.name), cmd.desc)
	}

	s.println()

	s.println("Commands:")
	for _, cmd := range s.cmds {
		s.printf(" %s - %s\n", white(cmd.name), cmd.desc)
	}
	s.println()
}

func (s *Shell) switchCmd(cmd string) {
	s.current = cmd
	s.reader.SetPrompt(yellow(s.current) + blue(prompt))
}

func (s *Shell) cleanCmd() {
	cmd := exec.Command("clear")
	cmd.Stdout = s.out
	_ = cmd.Run()
}

func (s *Shell) quitCmd() {
	s.quit = true
}

func (s *Shell) execCommandWith(input string) {
	if s.current == "" {
		return
	}
	cmd := s.mCmds[s.current]
	cmd.Run(cmd, []string{input})
}

func (s *Shell) println(a ...interface{}) {
	_, _ = fmt.Fprintln(s.out, a...)
}

func (s *Shell) printf(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(s.out, format, a...)
}

func (s *Shell) autoCompleter() readline.AutoCompleter {
	var pcs []readline.PrefixCompleterInterface

	for _, cmd := range basicCmds {
		pcs = append(pcs, readline.PcItem(cmd.name))
	}
	for _, cmd := range s.cmds {
		pcs = append(pcs, readline.PcItem(cmd.name))
	}
	return readline.NewPrefixCompleter(pcs...)
}

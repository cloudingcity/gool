package shell

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const prompt = "=# "

type Shell struct {
	in      io.Reader
	out     io.Writer
	current string
	scripts map[string]*cobra.Command
	names   []string
	quit    bool
}

func New(in io.Reader, out io.Writer) *Shell {
	return &Shell{
		in:      in,
		out:     out,
		scripts: make(map[string]*cobra.Command),
	}
}

func (s *Shell) Register(scripts ...*cobra.Command) {
	for _, script := range scripts {
		s.scripts[script.Name()] = script
		s.names = append(s.names, script.Name())
	}
}

func (s *Shell) Run() {
	color.New(color.FgCyan).Fprintln(s.out, `Gool Shell`)
	fmt.Fprintln(s.out, `\h: show help`)

	reader := bufio.NewReader(s.in)
	for {
		color.New(color.FgHiBlue).Fprint(s.out, s.current+prompt)
		text, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		input := strings.TrimSpace(text)
		s.run(input)

		if s.quit {
			fmt.Fprintln(s.out, `Good Bye!`)
			break
		}
	}
}

func (s *Shell) run(input string) {
	if input == "" {
		return
	}
	if strings.HasPrefix(input, `\`) && len(input) >= 2 {
		s.runCommand(input)
		return
	}
	s.runScript(input)
}

func (s *Shell) runCommand(cmd string) {
	switch cmd[1:2] {
	case "h":
		fmt.Fprintln(s.out, `  \h show help`)
		fmt.Fprintln(s.out, `  \c change current script`)
		fmt.Fprintln(s.out, `  \l list available scripts`)
		fmt.Fprintln(s.out, `  \q to quit`)
	case "c":
		fields := strings.Fields(cmd)
		if len(fields) < 2 {
			fmt.Fprintln(s.out, "Missing script")
			return
		}
		script := fields[1]
		if _, ok := s.scripts[script]; !ok {
			fmt.Fprintf(s.out, "script %q does not exists\n", script)
			return
		}
		s.current = script
	case "l":
		for _, name := range s.names {
			fmt.Fprintln(s.out, name)
		}
	case "q":
		s.quit = true
	}
}

func (s *Shell) runScript(input string) {
	if s.current == "" {
		return
	}
	script := s.scripts[s.current]
	script.Run(script, []string{input})
}

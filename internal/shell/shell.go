package shell

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const prompt = "=# "

var (
	cyan   = color.New(color.FgCyan).SprintfFunc()
	blue   = color.New(color.FgHiBlue).SprintfFunc()
	green  = color.New(color.FgGreen).SprintFunc()
	yellow = color.New(color.FgYellow).SprintfFunc()
)

type cmd struct {
	cmd  string
	desc string
}

var cmds = []cmd{
	{cmd: `\h`, desc: "show help"},
	{cmd: `\l`, desc: "list available scripts"},
	{cmd: `\s`, desc: "switch to the specified script"},
	{cmd: `\c`, desc: "clean the terminal screen"},
	{cmd: `\q`, desc: "to quit"},
}

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
	s.welcome()

	reader := bufio.NewReader(s.in)
	for {
		if s.current != "" {
			s.print(yellow(s.current))
		}
		s.print(blue(prompt))
		text, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		input := strings.TrimSpace(text)
		s.exec(input)

		if s.quit {
			s.println(`Good Bye!`)
			break
		}
	}
}

func (s *Shell) welcome() {
	s.println(cyan("Gool Shell"))
	s.println(cmds[0].cmd, cmds[0].desc)
}

func (s *Shell) exec(input string) {
	if input == "" {
		return
	}
	if strings.HasPrefix(input, `\`) && len(input) >= 2 {
		s.execCommand(input)
		return
	}
	s.execScript(input)
}

func (s *Shell) execCommand(cmd string) {
	switch cmd[1:2] {
	case "h":
		s.helpCmd()
	case "l":
		s.listCmd()
	case "s":
		s.switchCmd(cmd)
	case "c":
		s.cleanCmd()
	case "q":
		s.quitCmd()
	}
}

func (s *Shell) helpCmd() {
	s.println("Available Commands:")
	for _, cmd := range cmds {
		s.printf("  %s %s\n", green(cmd.cmd), yellow(cmd.desc))
	}
}

func (s *Shell) listCmd() {
	s.println("Available Scripts:")
	for _, name := range s.names {
		s.printf("  %s\n", yellow(name))
	}
}

func (s *Shell) switchCmd(cmd string) {
	fields := strings.Fields(cmd)
	if len(fields) < 2 {
		s.println("No script given")
		return
	}
	script := fields[1]
	if _, ok := s.scripts[script]; !ok {
		s.printf("script %q does not exists\n", script)
		return
	}
	s.current = script
}

func (s *Shell) cleanCmd() {
	cmd := exec.Command("clear")
	cmd.Stdout = s.out
	_ = cmd.Run()
}

func (s *Shell) quitCmd() {
	s.quit = true
}

func (s *Shell) execScript(input string) {
	if s.current == "" {
		return
	}
	script := s.scripts[s.current]
	script.Run(script, []string{input})
}

func (s *Shell) print(a ...interface{}) {
	_, _ = fmt.Fprint(s.out, a...)
}

func (s *Shell) println(a ...interface{}) {
	_, _ = fmt.Fprintln(s.out, a...)
}

func (s *Shell) printf(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(s.out, format, a...)
}

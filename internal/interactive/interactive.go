package interactive

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const prompt = "=# "

type Interactive struct {
	in      io.Reader
	out     io.Writer
	current string
	scripts map[string]*cobra.Command
	names   []string
	quit    bool
}

func New(in io.Reader, out io.Writer) *Interactive {
	return &Interactive{
		in:      in,
		out:     out,
		scripts: make(map[string]*cobra.Command),
	}
}

func (i *Interactive) Register(scripts ...*cobra.Command) {
	for _, script := range scripts {
		i.scripts[script.Name()] = script
		i.names = append(i.names, script.Name())
	}
}

func (i *Interactive) Run() {
	color.New(color.FgCyan).Fprintln(i.out, `Gool Shell`)
	fmt.Fprintln(i.out, `\h: show help`)

	reader := bufio.NewReader(i.in)
	for {
		if i.quit {
			fmt.Fprintln(i.out, `Good Bye!`)
			return
		}
		color.New(color.FgHiBlue).Fprint(i.out, i.current+prompt)
		text, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		input := strings.TrimSuffix(text, "\n")
		i.run(input)
	}
}

func (i *Interactive) run(input string) {
	if input == "" {
		return
	}
	if strings.HasPrefix(input, `\`) && len(input) >= 2 {
		i.runCommand(input)
		return
	}
	i.runScript(input)
}

func (i *Interactive) runCommand(cmd string) {
	switch cmd[1:2] {
	case "h":
		fmt.Fprintln(i.out, `  \h show help`)
		fmt.Fprintln(i.out, `  \c change current script`)
		fmt.Fprintln(i.out, `  \l list available scripts`)
		fmt.Fprintln(i.out, `  \q to quit`)
	case "c":
		fields := strings.Fields(cmd)
		if len(fields) < 2 {
			fmt.Fprintln(i.out, "Missing script")
			return
		}
		script := fields[1]
		if _, ok := i.scripts[script]; !ok {
			fmt.Fprintf(i.out, "script %q does not exists\n", script)
			return
		}
		i.current = script
	case "l":
		for _, name := range i.names {
			fmt.Fprintln(i.out, name)
		}
	case "q":
		i.quit = true
	}
}

func (i *Interactive) runScript(input string) {
	if i.current == "" {
		return
	}
	script := i.scripts[i.current]
	script.Run(script, []string{input})
}

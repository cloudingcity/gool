package shell

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

var (
	fakeOut = &bytes.Buffer{}
	fakeCmd = &cobra.Command{
		Use: "fake",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(123)
			fmt.Fprint(fakeOut, "fake exec")
		},
	}
)

func TestInteractive(t *testing.T) {
	t.Run("startup message", func(t *testing.T) {
		in := &bytes.Buffer{}
		out := &bytes.Buffer{}
		New(in, out).Run()
		got := out.String()

		assert.Contains(t, got, "Gool Shell")
		assert.Contains(t, got, `\h: show help`)
	})

	t.Run("enter blank", func(t *testing.T) {
		in := bytes.NewBufferString("\nenter something\n")
		out := &bytes.Buffer{}
		New(in, out).Run()
		got := out.String()

		assert.Contains(t, got, `=#`)
	})

	t.Run("enter \\h", func(t *testing.T) {
		in := bytes.NewBufferString("\\h\n")
		out := &bytes.Buffer{}
		New(in, out).Run()
		got := out.String()

		assert.Contains(t, got, `\h: show help`)
		assert.Contains(t, got, `\c change current script`)
		assert.Contains(t, got, `\l list available scripts`)
		assert.Contains(t, got, `\q to quit`)
	})

	t.Run("enter \\c", func(t *testing.T) {
		t.Run("without script", func(t *testing.T) {
			in := bytes.NewBufferString("\\c\n")
			out := &bytes.Buffer{}
			New(in, out).Run()
			got := out.String()

			assert.Contains(t, got, "No script given")
		})
		t.Run("script not found", func(t *testing.T) {
			in := bytes.NewBufferString("\\c not-exists\n")
			out := &bytes.Buffer{}
			New(in, out).Run()
			got := out.String()

			assert.Contains(t, got, `script "not-exists" does not exists`)
		})
		t.Run("exec script", func(t *testing.T) {
			in := bytes.NewBufferString("\\c fake\nenter something\n")
			out := &bytes.Buffer{}
			shell := New(in, out)
			shell.Register(fakeCmd)
			shell.Run()
			got := out.String()

			assert.Contains(t, got, "fake=#")
			assert.Contains(t, fakeOut.String(), "fake exec")
			fakeOut = &bytes.Buffer{}
		})
	})

	t.Run("enter \\l", func(t *testing.T) {
		in := bytes.NewBufferString("\\l\n")
		out := &bytes.Buffer{}
		shell := New(in, out)
		shell.Register(fakeCmd)
		shell.Run()
		got := out.String()

		assert.Contains(t, got, "fake")
	})

	t.Run("enter \\q", func(t *testing.T) {
		in := bytes.NewBufferString("\\q\n")
		out := &bytes.Buffer{}
		New(in, out).Run()
		got := out.String()

		assert.Contains(t, got, "Good Bye!")
	})
}

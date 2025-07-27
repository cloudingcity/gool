package shell

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

var (
	fakeOut = &bytes.Buffer{}
	fakeCmd = &cobra.Command{
		Use:   "fake",
		Short: "fake usage message",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(fakeOut, "fake exec")
		},
	}
)

func TestInteractive(t *testing.T) {
	t.Run("startup message", func(t *testing.T) {
		in := io.NopCloser(&bytes.Buffer{})
		out := &bytes.Buffer{}
		New(in, out).Run()
		got := out.String()

		assert.Contains(t, got, "/help for more information")
	})

	t.Run("enter /help", func(t *testing.T) {
		in := io.NopCloser(bytes.NewBufferString("/help\n"))
		out := &bytes.Buffer{}
		shell := New(in, out)
		shell.Register(fakeCmd)
		shell.Run()
		got := out.String()

		assert.Contains(t, got, "Basics:")
		assert.Contains(t, got, "/help - Show help")
		assert.Contains(t, got, "/clear - Clean the screen")
		assert.Contains(t, got, "/quit - Quit the REPL")
		assert.Contains(t, got, "Commands:")
		assert.Contains(t, got, "/fake - fake usage message")
	})

	t.Run("enter /quit", func(t *testing.T) {
		in := io.NopCloser(bytes.NewBufferString("/quit\n"))
		out := &bytes.Buffer{}
		New(in, out).Run()
		got := out.String()

		assert.Contains(t, got, "Good Bye!")
	})

	t.Run("enter /fake", func(t *testing.T) {
		t.Run("exec /fake", func(t *testing.T) {
			in := io.NopCloser(bytes.NewBufferString("/fake\nenter something\n"))
			out := &bytes.Buffer{}
			shell := New(in, out)
			shell.Register(fakeCmd)
			shell.Run()

			assert.Contains(t, fakeOut.String(), "fake exec")
			fakeOut = &bytes.Buffer{}
		})
	})
}

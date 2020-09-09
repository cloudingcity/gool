package shell

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

var (
	fakeOut = &bytes.Buffer{}
	fakeCmd = &cobra.Command{
		Use: "fake",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(fakeOut, "fake exec")
		},
	}
)

func TestInteractive(t *testing.T) {
	t.Run("startup message", func(t *testing.T) {
		in := ioutil.NopCloser(&bytes.Buffer{})
		out := &bytes.Buffer{}
		New(in, out).Run()
		got := out.String()

		assert.Contains(t, got, "Gool Shell")
		assert.Contains(t, got, `\h show help`)
	})

	t.Run("enter \\h", func(t *testing.T) {
		in := ioutil.NopCloser(bytes.NewBufferString("\\h\n"))
		out := &bytes.Buffer{}
		New(in, out).Run()
		got := out.String()

		assert.Contains(t, got, "Commands:")
		assert.Contains(t, got, `\h show help`)
		assert.Contains(t, got, `\l list available scripts`)
		assert.Contains(t, got, `\s switch to the specified script`)
		assert.Contains(t, got, `\q to quit`)
	})

	t.Run("enter \\l", func(t *testing.T) {
		in := ioutil.NopCloser(bytes.NewBufferString("\\l\n"))
		out := &bytes.Buffer{}
		shell := New(in, out)
		shell.Register(fakeCmd)
		shell.Run()
		got := out.String()

		assert.Contains(t, got, "Scripts:")
		assert.Contains(t, got, "fake")
	})

	t.Run("enter \\s", func(t *testing.T) {
		t.Run("without script", func(t *testing.T) {
			in := ioutil.NopCloser(bytes.NewBufferString("\\s\n"))
			out := &bytes.Buffer{}
			New(in, out).Run()
			got := out.String()

			assert.Contains(t, got, "no script given")
		})
		t.Run("script not found", func(t *testing.T) {
			in := ioutil.NopCloser(bytes.NewBufferString("\\s not-exists\n"))
			out := &bytes.Buffer{}
			New(in, out).Run()
			got := out.String()

			assert.Contains(t, got, `script "not-exists" does not exists`)
		})
		t.Run("exec script", func(t *testing.T) {
			in := ioutil.NopCloser(bytes.NewBufferString("\\s fake\nenter something\n"))
			out := &bytes.Buffer{}
			shell := New(in, out)
			shell.Register(fakeCmd)
			shell.Run()

			assert.Contains(t, fakeOut.String(), "fake exec")
			fakeOut = &bytes.Buffer{}
		})
	})

	t.Run("enter \\r", func(t *testing.T) {
		t.Run("without script", func(t *testing.T) {
			in := ioutil.NopCloser(bytes.NewBufferString("\\r\n"))
			out := &bytes.Buffer{}
			New(in, out).Run()
			got := out.String()

			assert.Contains(t, got, "run script failed")
		})
		t.Run("script not found", func(t *testing.T) {
			in := ioutil.NopCloser(bytes.NewBufferString("\\r not-exists-script foo\n"))
			out := &bytes.Buffer{}
			New(in, out).Run()
			got := out.String()

			assert.Contains(t, got, `script "not-exists-script" does not exists`)
		})
		t.Run("exec script", func(t *testing.T) {
			in := ioutil.NopCloser(bytes.NewBufferString("\\r fake enter something\n"))
			out := &bytes.Buffer{}
			shell := New(in, out)
			shell.Register(fakeCmd)
			shell.Run()

			assert.Contains(t, fakeOut.String(), "fake exec")
			fakeOut = &bytes.Buffer{}
		})
	})

	t.Run("enter \\q", func(t *testing.T) {
		in := ioutil.NopCloser(bytes.NewBufferString("\\q\n"))
		out := &bytes.Buffer{}
		New(in, out).Run()
		got := out.String()

		assert.Contains(t, got, "Good Bye!")
	})
}

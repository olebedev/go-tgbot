package tgbot

import (
	"fmt"
	"io"
	"time"

	"github.com/fatih/color"
)

func Recover(c *Context) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch r := r.(type) {
			case error:
				err = r
			default:
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	return c.Next()
}

func Logger(out io.Writer) func(*Context) error {
	cyan := color.New(color.FgCyan)
	white := color.New(color.FgWhite)
	red := color.New(color.FgRed)
	yellow := color.New(color.FgYellow)
	green := color.New(color.FgGreen)
	return func(c *Context) error {
		start := time.Now()
		p := c.Path
		err := c.Next()
		pp := c.Path
		indent := "16"

		cyan.Fprintf(out, "%13v ", time.Now().Sub(start))
		fmt.Fprint(out, "<~")
		white.Fprintf(out, " %s", p)
		if c.Text != "" {
			white.Fprintf(out, " (payload: '%s')", c.Text)
		}

		if p != pp {
			white.Fprintf(out, "\n%"+indent+"v └── %s (override)", "", pp)
			indent = "20"
		}
		defer fmt.Fprintln(out)

		white.Fprintf(out, "\n%"+indent+"v └── ", "")
		if err != nil {
			if e, ok := err.(*Error); ok && e.Code == 404 {
				yellow.Fprint(out, err)
			} else {
				red.Fprint(out, err)
			}
		} else {
			green.Fprint(out, "OK")
		}

		return err
	}
}

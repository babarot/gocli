package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
)

const (
	app     = "gocli"
	version = "0.1.0"
)

// CLI represents the attributes for command-line interface
type CLI struct {
	opt  option
	args []string

	stdout io.Writer
	stderr io.Writer
}

type option struct {
	version bool
}

func main() {
	os.Exit(newCLI(os.Args[1:]).run())
}

func newCLI(args []string) CLI {
	var c CLI

	c.stdout = os.Stdout
	c.stderr = os.Stderr

	flag.BoolVar(&c.opt.version, "version", false, "show version")
	flag.BoolVar(&c.opt.version, "v", false, "show version")
	flag.Parse()

	c.args = flag.Args()

	return c
}

func (c CLI) exit(msg interface{}) int {
	switch m := msg.(type) {
	case int:
		return m
	case nil:
		return 0
	case string:
		fmt.Fprintf(c.stdout, "%s\n", m)
		return 0
	case error:
		fmt.Fprintf(c.stderr, "[ERROR] %s: %s\n", app, m.Error())
		return 1
	default:
		panic(msg)
	}
}

func (c CLI) run() int {
	if c.opt.version {
		return c.exit(fmt.Sprintf("%s v%s (runtime: %s)", app, version, runtime.Version()))
	}

	return c.exit(fmt.Sprintf("%#v", c))
}

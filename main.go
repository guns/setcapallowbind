/*
 * Copyright (c) 2016 Sung Pae <self@sungpae.com>
 */

package main

import (
	"fmt"
	"os"
	"path"
	"syscall"

	"github.com/jessevdk/go-flags"
)

const usagesummary = `path

Simple wrapper that calls:

  /usr/bin/setcap cap_net_bind_service=+ep path

Add to sudoers file:

  user ALL = (ALL) NOPASSWD: /path/to/setcapallowbind`

func abort(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	os.Exit(1)
}

func abortWithHelp(parser *flags.Parser) {
	parser.WriteHelp(os.Stderr)
	os.Exit(1)
}

func getopts(arguments []string) (parser *flags.Parser, args []string) {
	var err error

	parser = flags.NewNamedParser(path.Base(arguments[0]), flags.HelpFlag|flags.PassDoubleDash)
	parser.Usage = usagesummary

	if args, err = parser.ParseArgs(arguments[1:]); err != nil {
		abort(err)
	}

	return parser, args
}

func main() {
	parser, args := getopts(os.Args)
	if len(args) != 1 {
		abortWithHelp(parser)
	} else if _, err := os.Stat(args[0]); os.IsNotExist(err) {
		abort(err)
	}

	if err := syscall.Exec("/usr/bin/setcap", []string{"setcap", "cap_net_bind_service=+ep", args[0]}, nil); err != nil {
		abort(err)
	}
}

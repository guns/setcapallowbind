/*
 * Copyright (c) 2016-2017 Sung Pae <self@sungpae.com>
 * Distributed under the MIT license.
 * http://www.opensource.org/licenses/mit-license.php
 */

package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"syscall"

	"github.com/jessevdk/go-flags"
)

const usage = `file …

Simple wrapper that calls:

  /usr/bin/setcap cap_net_bind_service=+ep file …

Add to sudoers file:

  user ALL = (ALL) NOPASSWD: /path/to/setcapallowbind`

type options struct {
	Help bool `short:"h" long:"help"`
}

func validate(opts *options, args []string) error {
	if len(args) == 0 {
		return errors.New("nothing to do; see --help")
	}

	for i := range args {
		if _, err := os.Stat(args[i]); os.IsNotExist(err) {
			return err
		}
	}

	return nil
}

func abort(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	os.Exit(1)
}

func getopts(arguments []string) (opts *options, args []string) {
	opts = new(options)
	var err error

	parser := flags.NewNamedParser(path.Base(arguments[0]), flags.PassDoubleDash)
	parser.Usage = usage

	if _, err = parser.AddGroup("Options", "", opts); err != nil {
		abort(err)
	}

	if args, err = parser.ParseArgs(arguments[1:]); err != nil {
		abort(err)
	}

	if opts.Help {
		parser.WriteHelp(os.Stderr)
		os.Exit(0)
	}

	if err = validate(opts, args); err != nil {
		abort(err)
	}

	return opts, args
}

func main() {
	_, args := getopts(os.Args)
	cmd := []string{"setcap"}

	for i := range args {
		cmd = append(cmd, "cap_net_bind_service=+ep", args[i])
	}

	if err := syscall.Exec("/usr/bin/setcap", cmd, nil); err != nil {
		abort(err)
	}
}

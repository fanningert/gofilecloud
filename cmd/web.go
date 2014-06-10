// Copyright 2014 The goFileCloud Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/codegangsta/cli"
)

var CmdWeb = cli.Command{
	Name:        "web",
	Usage:       "Start goFileCloud web server",
	Description: `bla bla`,
	Action:      runWeb,
	Flags: []cli.Flag{
		cli.StringFlag{"config, w", "config.ini", "Set the path for the config file"},
	},
}

func runWeb(c *cli.Context) {
	fmt.Println("Test")
}

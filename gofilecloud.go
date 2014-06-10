// Copyright 2014 The goFileCloud Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/codegangsta/cli"
	"github.com/fanningert/gofilecloud/cmd"
	"github.com/fanningert/gofilecloud/modules/setting"
	"os"
	"runtime"
)

const APP_VER = "0.0.1"

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	setting.AppVer = APP_VER
}

func main() {
	app := cli.NewApp()
	app.Name = "goFileCloud"
	app.Usage = ""
	app.Version = APP_VER
	app.Commands = []cli.Command{
		cmd.CmdWeb,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}

// Copyright 2025 Innkeeper DanteSu(苏孟君) <mengjunsu@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/DanteSu/miniblog

package main

import (
	"github.com/DanteSu/miniblog/internal/miniblog"
	_ "go.uber.org/automaxprocs"
	"os"
)

func main() {
	command := miniblog.NewMiniBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}

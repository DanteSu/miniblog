// Copyright 2025 Innkeeper DanteSu(苏孟君) <mengjunsu@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/DanteSu/miniblog

package miniblog

import (
	"encoding/json"
	"fmt"
	"github.com/DanteSu/miniblog/internal/pkg/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// NewMiniBlogCommand create a *cobra.Command object. after, we can use Command object's Execute func to startup the service.
func NewMiniBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "miniblog",
		Short: "miniblog",
		Long: `miniblog service, used to create user with basic information.
			Find more miniblog information at: https://github.com/DanteSu/miniblog`,
		SilenceUsage: true,
		// when cmd.Execute(), call this run fuc
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Init(logOptions())
			defer log.Sync()
			return run()
		},
		// when command runs, do not need to indicate command line parameter
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}

			return nil
		},
	}

	cobra.OnInitialize(initConfig)

	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "The path to the miniblog configuration file. Empty string for no configuration file.")

	cmd.Flags().BoolP("toggle", "t", false, "Print the version number.")

	return cmd
}

// real entrypoint for service
func run() error {
	// 打印所有的配置项及其值
	settings, _ := json.Marshal(viper.AllSettings())
	log.Infow(string(settings))
	// 打印 db -> username 配置项的值
	log.Infow(viper.GetString("db.username"))
	return nil
}

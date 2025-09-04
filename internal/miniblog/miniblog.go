// Copyright 2025 Innkeeper DanteSu(苏孟君) <mengjunsu@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/DanteSu/miniblog

package miniblog

import (
	"context"
	"errors"
	"fmt"
	"github.com/DanteSu/miniblog/internal/pkg/log"
	mw "github.com/DanteSu/miniblog/internal/pkg/middleware"
	"github.com/DanteSu/miniblog/pkg/version/verflag"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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
			verflag.PrintAndExitIfRequested()
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

	verflag.AddFlags(cmd.PersistentFlags())

	return cmd
}

// real entrypoint for service
func run() error {
	// 初始化 store 层
	if err := initStore(); err != nil {
		return err
	}

	// 设置gin模式
	gin.SetMode(viper.GetString("runmode"))

	g := gin.New()

	// gin.Recovery() 中间件，用来捕获任何 panic，并恢复
	mws := []gin.HandlerFunc{gin.Recovery(), mw.NoCache, mw.Cors, mw.Secure, mw.RequestID()}

	g.Use(mws...)

	if err := installRouters(g); err != nil {
		return err
	}

	// http server实例
	httpsrv := &http.Server{Addr: viper.GetString("addr"), Handler: g}

	log.Infow("Start to listening the incoming requests on http address", "addr", viper.GetString("addr"))
	go func() {
		if err := httpsrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalw(err.Error())
		}
	}()
	// 等待中断信号优雅关闭服务器
	quit := make(chan os.Signal)
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的 CTRL + C 就是触发系统 SIGINT 信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Infow("Shutting down server ...")
	// 创建 ctx 用于通知服务器 goroutine, 它有 10 秒时间完成当前正在处理的请求
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpsrv.Shutdown(ctx); err != nil {
		log.Errorw("Insecure Server forced to shutdown", "err", err)
		return err
	}

	log.Infow("Server exiting")

	// 打印所有的配置项及其值
	//settings, _ := json.Marshal(viper.AllSettings())
	//log.Infow(string(settings))
	// 打印 db -> username 配置项的值
	//log.Infow(viper.GetString("db.username"))
	return nil
}

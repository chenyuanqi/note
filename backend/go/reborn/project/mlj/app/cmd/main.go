package main

import (
	"fmt"
	"os"

	"mlj/app/cmd/controllers"
	btConfig "mlj/config"
	"mlj/pkg/common/helpers"
	"mlj/pkg/config"
	"mlj/pkg/console"

	"github.com/spf13/cobra"
)

var ROOT string

func init() {
	// 加载 config 目录下的配置信息
	btConfig.Initialize()
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   config.Get("app.name"),
		Short: "mlj project",
		Long:  `Default will no command, you can use "-h" flag to see all subcommands`,
		// rootCmd 的所有子命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {
			// 配置初始化，依赖命令行 --env 参数
			config.InitConfig(Env)

			// 初始化 Logger

			// 初始化数据库

			// 初始化 Redis

			// 初始化缓存
		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		controllers.CmdKey,
		controllers.CmdPlay,
		controllers.CmdUser,
	)

	// 注册全局参数，--env
	RegisterGlobalFlags(rootCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}

// Env 存储全局选项 --env 的值
var Env string

// RegisterGlobalFlags 注册全局选项（flag）
func RegisterGlobalFlags(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().StringVarP(&Env, "env", "e", "", "load .env file, example: --env=testing will use .env.testing file")
}

// RegisterDefaultCmd 注册默认命令
func RegisterDefaultCmd(rootCmd *cobra.Command, subCmd *cobra.Command) {
	cmd, _, err := rootCmd.Find(os.Args[1:])
	firstArg := helpers.FirstElement(os.Args[1:])
	if err == nil && cmd.Use == rootCmd.Use && firstArg != "-h" && firstArg != "--help" {
		args := append([]string{subCmd.Use}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}
}

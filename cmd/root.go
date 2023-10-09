package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gokmp",
		Short: "tool to talk with your KMP meters",
		Run:   root,
	}
)

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("gokmp")
	viper.SetEnvKeyReplacer(strings.NewReplacer(`-`, `_`))
	flags := rootCmd.PersistentFlags()
	flags.String("port", "/dev/ttyUSB0", "serial port")
	viper.BindPFlags(flags)
}

func root(cmd *cobra.Command, args []string) {
	fmt.Println("the root command does nothing, use the subcommands")
}

func Execute() {
	rootCmd.Execute()
}

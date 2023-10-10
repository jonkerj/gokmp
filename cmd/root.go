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
	flags.String("port", "", "serial port")
	flags.String("serial-pid", "6001", "USB PID of serial port")
	flags.String("serial-vid", "0403", "USB VID of serial port")
	flags.String("serial-serial", "", "USB Serial number of serial port")
	viper.BindPFlags(flags)
}

func root(cmd *cobra.Command, args []string) {
	fmt.Println("the root command does nothing, use the subcommands")
}

func Execute() {
	rootCmd.Execute()
}

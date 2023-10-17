package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gokmp",
		Short: "tool to talk with your KMP meters",
		Run:   root,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			level := slog.LevelInfo
			if viper.GetBool("verbose") {
				fmt.Println("setting verbose")
				level = slog.LevelDebug
			}

			handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
			logger := slog.New(handler)
			slog.SetDefault(logger)
		},
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
	flags.Bool("verbose", false, "Verbose logging")
	if err := viper.BindPFlags(flags); err != nil {
		panic(err)
	}
}

func root(cmd *cobra.Command, args []string) {
	fmt.Println("the root command does nothing, use the subcommands")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

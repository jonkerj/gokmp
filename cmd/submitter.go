package cmd

import (
	"context"
	"time"

	"github.com/jonkerj/gokmp/internal/serial"
	"github.com/jonkerj/gokmp/internal/submitter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	submitCmd = &cobra.Command{
		Use:   "submitter",
		Short: "Poll and send to influxdb",
		Run:   submit,
	}
)

func init() {
	flags := submitCmd.PersistentFlags()
	flags.Duration("interval", time.Minute*15, "Interval between polls, expressed as golang duration")
	flags.String("influxdb-org", "", "InfluxDB Organization")
	flags.String("influxdb-url", "http://influxdb.influxdb:8086", "InfluxDB URL")
	flags.String("influxdb-token", "notme:notmypassword", "InfluxDB token")
	flags.String("influxdb-bucket", "iioflux/autogen", "InfluxDB bucket")
	rootCmd.AddCommand(submitCmd)

	err := viper.BindPFlags(flags)
	if err != nil {
		panic(err)
	}
}

func submit(cmd *cobra.Command, args []string) {
	portName, err := serial.GetPortName(viper.GetString("port"), viper.GetString("serial-vid"), viper.GetString("serial-pid"), viper.GetString("serial-serial"))
	if err != nil {
		panic(err)
	}

	s, err := submitter.NewSubmitter(
		context.TODO(),
		portName,
		viper.GetString("influxdb-url"),
		viper.GetString("influxdb-token"),
		viper.GetString("influxdb-org"),
		viper.GetString("influxdb-bucket"),
		viper.GetDuration("interval"),
	)
	if err != nil {
		panic(err)
	}
	s.Run()
}

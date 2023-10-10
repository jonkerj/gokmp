package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/jonkerj/gokmp/internal/serial"
	"github.com/jonkerj/gokmp/internal/submitter"
	"github.com/jonkerj/gokmp/pkg/client"
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
	port, err := serial.Open(viper.GetString("port"), viper.GetString("serial-vid"), viper.GetString("serial-pid"), viper.GetString("serial-serial"))
	if err != nil {
		panic(err)
	}

	c := client.NewSerialClient(port)
	sn, err := c.GetSerialNo()
	if err != nil {
		panic(err)
	}

	typ, version, err := c.GetType()
	if err != nil {
		panic(err)
	}

	fmt.Printf("type: %02x, version: %02x, serial: %d\n", typ, version, sn)
	s, err := submitter.NewSubmitter(
		context.TODO(),
		c,
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

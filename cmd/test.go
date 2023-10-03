package cmd

import (
	"fmt"
	"time"

	"github.com/jonkerj/gokmp/pkg/application"
	"github.com/jonkerj/gokmp/pkg/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.bug.st/serial"
)

var (
	testCmd = &cobra.Command{
		Use:   "test",
		Short: "Test heat meter",
		Run:   testMeter,
	}
)

func init() {
	rootCmd.AddCommand(testCmd)
}

func testMeter(cmd *cobra.Command, args []string) {
	port, err := serial.Open(
		viper.GetString("port"),
		&serial.Mode{
			BaudRate: 1200,
			DataBits: 8,
			StopBits: 2,
			Parity:   serial.NoParity,
		},
	)
	if err != nil {
		panic(err)
	}

	err = port.SetReadTimeout(100 * time.Millisecond)
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
	fetchEm := []application.RegisterID{
		application.RegisterHeatEnergy,
		application.RegisterVolumeRegister1,
		application.RegisterCurrentInTemp,
		application.RegisterCurrentReturnTemp,
	}
	regs, err := c.GetRegister(fetchEm)
	if err != nil {
		panic(err)
	}

	for _, r := range regs {
		fmt.Println(r)
	}
}

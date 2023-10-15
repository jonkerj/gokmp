package serial

import (
	"errors"
	"fmt"
	"time"

	"go.bug.st/serial"
	"go.bug.st/serial/enumerator"
)

var (
	PortNotFoundError    = errors.New("port not found")
	NoPortsError         = errors.New("no serial ports found")
	NoMatchingPortsError = errors.New("no matching ports found")
)

func getPortName(portName, vid, pid, serialNo string) (string, error) {
	name := portName
	if name != "" {
		return name, nil
	}

	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		return "", fmt.Errorf("error enumerating ports: %v", err)
	}

	if len(ports) == 0 {
		return "", NoPortsError
	}

	for _, port := range ports {
		if port.IsUSB && port.VID == vid && port.PID == pid && port.SerialNumber == serialNo {
			return port.Name, nil
		}
	}

	return "", NoMatchingPortsError
}

func Open(givenPortName, vid, pid, serialNo string) (serial.Port, error) {
	portName, err := getPortName(givenPortName, vid, pid, serialNo)
	if err != nil {
		return nil, err
	}

	port, err := serial.Open(
		portName,
		&serial.Mode{
			BaudRate: 1200,
			DataBits: 8,
			StopBits: 2,
			Parity:   serial.NoParity,
		},
	)
	if err != nil {
		return nil, err
	}

	err = port.SetReadTimeout(100 * time.Millisecond)
	if err != nil {
		return nil, err
	}

	return port, nil
}

package client

import (
	"fmt"
	"log/slog"

	"github.com/jonkerj/gokmp/pkg/application"
	"github.com/jonkerj/gokmp/pkg/datalink"
	"go.bug.st/serial"
)

type (
	SerialClient struct {
		port serial.Port
	}
)

func NewSerialClient(port serial.Port) *SerialClient {
	return &SerialClient{
		port: port,
	}
}

func (s *SerialClient) command(command application.Command) (application.Command, error) {
	commandFrame := command.ToFrame()
	commandBytes := commandFrame.EncodeFrame()

	written, err := s.port.Write(commandBytes)
	if err != nil {
		return nil, fmt.Errorf("could not write to serial port: %w", err)
	}

	if written < len(commandBytes) {
		return nil, fmt.Errorf("short write: %d bytes written, but command was %d bytes long", written, len(commandBytes))
	}

	buff := make([]byte, 1024) // probbaly too much
	bytesRead := []byte{}

	for {
		lenRead, err := s.port.Read(buff)
		if err != nil {
			return nil, fmt.Errorf("could not read from port: %w", err)
		}

		if lenRead == 0 {
			break
		}

		bytesRead = append(bytesRead, buff[:lenRead]...)
	}

	bytesRead = bytesRead[len(commandBytes):] // remove command echo

	if len(bytesRead) < 5 {
		return nil, fmt.Errorf("short read: %d bytes read", len(bytesRead))
	}

	responseFrame, err := datalink.DecodeFrame(bytesRead)
	if err != nil {
		return nil, fmt.Errorf("could not decode data into frame: %w", err)
	}

	response, err := command.FromFrame(*responseFrame)
	if err != nil {
		return nil, fmt.Errorf("could not decode frame into response: %w", err)
	}

	return response, nil
}

func (s *SerialClient) GetType() (uint16, uint16, error) {
	slog.Debug("GetType")
	command := application.NewGetType()
	response, err := s.command(command)
	if err != nil {
		return 0, 0, err
	}

	gt := response.(application.GetType)
	return gt.Type, gt.Version, nil
}

func (s *SerialClient) GetSerialNo() (uint32, error) {
	slog.Debug("GetSerialNo")
	command := application.NewGetSerialNo()
	response, err := s.command(command)
	if err != nil {
		return 0, err
	}

	return response.(application.GetSerialNo).Serial, nil
}

func (s *SerialClient) GetRegister(registerIds []application.RegisterID) ([]application.Register, error) {
	slog.Debug("GetRegister", "registers", fmt.Sprintf("%v", registerIds))
	command := application.NewGetRegister(registerIds)
	response, err := s.command(command)
	if err != nil {
		return nil, err
	}

	return response.(application.GetRegister).Registers, nil
}

package submitter

import (
	"context"
	"errors"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/domain"
	"github.com/jonkerj/gokmp/pkg/application"
	gokmpclient "github.com/jonkerj/gokmp/pkg/client"
)

type (
	Submitter struct {
		context        context.Context
		gokmp          *gokmpclient.SerialClient
		influxWriteAPI api.WriteAPIBlocking
		interval       time.Duration
	}
)

func NewSubmitter(ctx context.Context, gokmp *gokmpclient.SerialClient, url, token, org, bucket string, interval time.Duration) (*Submitter, error) {
	client := influxdb2.NewClient(url, token)
	health, err := client.Health(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting status from influxdb: %v", err)
	}

	if health.Status != domain.HealthCheckStatusPass {
		return nil, fmt.Errorf("influxdb server was not healthy, status=%v", health.Status)
	}

	return &Submitter{
		context:        ctx,
		gokmp:          gokmp,
		influxWriteAPI: client.WriteAPIBlocking(org, bucket),
		interval:       interval,
	}, nil
}

func (s *Submitter) poll() error {
	regs, err := s.gokmp.GetRegister([]application.RegisterID{
		application.RegisterHeatEnergy,
		application.RegisterVolumeRegister1,
		application.RegisterCurrentInTemp,
		application.RegisterCurrentReturnTemp,
	})
	if err != nil {
		return fmt.Errorf("error getting KMP registers: %v", err)
	}

	if len(regs) != 4 {
		return errors.New("4 registers expected")
	}

	for _, reg := range regs {
		fmt.Printf("got reg: %v\n", reg)
	}

	p := influxdb2.NewPointWithMeasurement("heat").
		AddTag("location", "gareelhoek").
		AddField("energy", regs[0].Value*10e9).
		AddField("volume", regs[1].Value).
		AddField("t_in", regs[2].Value).
		AddField("t_out", regs[3].Value).
		SetTime(time.Now())

	if err = s.influxWriteAPI.WritePoint(s.context, p); err != nil {
		return fmt.Errorf("error submitting to influxdb: %v", err)
	}

	return nil
}

func (s *Submitter) Run() {
	ticker := time.NewTicker(s.interval)
	for ; true; <-ticker.C {
		s.poll()
	}
}

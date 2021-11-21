package app

import (
	"context"
	"fmt"
	"github.com/damek86/go-windhager-infowin/pkg/api"
	"strconv"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

const (
	dataPointName = "infowin"
	orgName       = ""
)

type Config struct {
	ScanInterval time.Duration
	InfluxDB     string
	ServerURL    string
	AuthToken    string
}

type InfluxExporter struct {
	client        influxdb2.Client
	InfoWinClient api.Client
	cfg           Config
}

func NewInfluxExporter(cfg Config, client api.Client) *InfluxExporter {
	return &InfluxExporter{
		client:        influxdb2.NewClient(cfg.ServerURL, cfg.AuthToken),
		cfg:           cfg,
		InfoWinClient: client,
	}
}

func (i *InfluxExporter) Work() {
	fmt.Println("start executor")
	i.work()
	for true {
		select {
		case <-time.After(i.cfg.ScanInterval):
			i.work()
		}
	}
}

func (i *InfluxExporter) work() {
	points, err := i.InfoWinClient.GetDataPoints()
	if err != nil {
		fmt.Println(err)
	}

	writeAPI := i.client.WriteAPIBlocking(orgName, i.cfg.InfluxDB)
	dataPoints := make(map[string]interface{})
	for name, datapoint := range points {
		//fmt.Printf("%s: %s %s\n", name, datapoint.Value, datapoint.Unit)
		if datapoint.Unit == "°C" || datapoint.Unit == "t" {
			if s, err := strconv.ParseFloat(datapoint.Value, 64); err == nil {
				dataPoints[name] = s
			}
		}
		if datapoint.Unit == "%" || datapoint.Unit == "" || datapoint.Unit == "h" {
			if s, err := strconv.ParseInt(datapoint.Value, 10, 32); err == nil {
				dataPoints[name] = s
			}
		}
	}

	p := influxdb2.NewPoint(dataPointName,
		nil,
		dataPoints, time.Now())
	err = writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		fmt.Println(err)
		// do not handle error since writeAPI already print a message
	}
}

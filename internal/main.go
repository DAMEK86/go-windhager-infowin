package main

import (
	"fmt"
	"github.com/damek86/go-windhager-infowin/internal/app"
	"github.com/damek86/go-windhager-infowin/pkg/api"
	"os"
	"strconv"
	"time"
)

func main() {
	client := api.NewClient(
		getEnv("INFOWIN_URL", "192.168.4.100"),
		getEnv("INFOWIN_USERNAME", api.DefaultCustomerUsername),
		getEnv("INFOWIN_PASSWORD", api.DefaultCustomerUserPassword),
	)
	go app.StartHealthEndpoint()
	app.NewInfluxExporter(getAppConfig(), client).Work()
}

func getAppConfig() app.Config {
	scanInterval, err := strconv.Atoi(getEnv("SCAN_INTERVAL", "30"))
	if err != nil {
		panic(err)
	}
	return app.Config{
		ScanInterval: time.Duration(scanInterval) * time.Second,
		InfluxDB:     getEnv("INFLUXDB_DB", "infowin"),
		ServerURL:    fmt.Sprintf("%s:%s", getEnv("INFLUXDB_HOST", "http://localhost"), getEnv("INFLUXDB_PORT", "8086")),
		AuthToken:    fmt.Sprintf("%s:%s", getEnv("INFLUXDB_USER", "root"), getEnv("INFLUXDB_PASSWORD", "root")),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

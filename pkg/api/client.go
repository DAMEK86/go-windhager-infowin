package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	dac "github.com/xinsnake/go-http-digest-auth-client"
)

const (
	DefaultServiceUsername      string = "Service"
	DefaultServicePassword      string = "123"
	DefaultCustomerUsername     string = "USER"
	DefaultCustomerUserPassword string = "123"

	BaseDatapointPath string = "api/1.0/datapoint"
	BaseResourcePath  string = "res/xml"
)

/*Client abstracts http client*/
type Client interface {
	GetDataPoint(datapointPath string) (*Datapoint, error)
	GetDataPoints() (map[string]Datapoint, error)
}

type client struct {
	httpClient dac.DigestTransport
	url        string
}

/*NewClient creates a new instance of Client*/
func NewClient(baseURL, username, password string) Client {
	t := dac.NewTransport(username, password)
	return &client{
		httpClient: t,
		url:        fmt.Sprintf("http://%s/%s", baseURL, BaseDatapointPath),
	}
}

func (c *client) GetDataPoint(datapointPath string) (*Datapoint, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.url, datapointPath), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	point := &Datapoint{}
	err = json.NewDecoder(resp.Body).Decode(point)

	if err != nil {
		return nil, fmt.Errorf("failed decode http response: %s", err)
	}

	return point, nil
}

func (c *client) GetDataPoints() (map[string]Datapoint, error) {
	results := map[string]Datapoint{}
	for _, i := range c.getDataPoints() {
		point, err := c.GetDataPoint(i.Oid)
		if err != nil {
			fmt.Println(err)
			continue
		}
		results[i.Name] = *point
	}

	return results, nil
}

func (c *client) getDataPoints() []DatapointRequest {
	return []DatapointRequest{
		{
			Name: "circuit_1_mode",
			Oid:  "/1/15/0/3/50/0",
		},
		{
			Name: "circuit_2_mode",
			Oid:  "/1/15/1/3/50/0",
		},
		{
			Name: "outside_temp",
			Oid:  "/1/15/0/0/0/0",
		},
		{
			Name: "circuit_1_flow_temp_actual",
			Oid:  "/1/15/0/0/2/0",
		},
		{
			Name: "circuit_1_flow_temp_set_point",
			Oid:  "/1/15/0/1/2/0",
		},
		{
			Name: "circuit_1_mixing_valve",
			Oid:  "/1/15/0/1/21/0",
		},
		{
			Name: "circuit_2_flow_temp_actual",
			Oid:  "/1/15/1/0/2/0",
		},
		{
			Name: "circuit_2_flow_temp_set_point",
			Oid:  "/1/15/1/1/2/0",
		},
		{
			Name: "circuit_2_mixing_valve",
			Oid:  "/1/15/1/1/21/0",
		},
		{
			Name: "burner_starts",
			Oid:  "/1/60/0/2/80/0",
		},
		{
			Name: "burner_operating_hours",
			Oid:  "/1/60/0/2/81/0",
		},
		{
			Name: "burner_current_boiler_output",
			Oid:  "/1/60/0/0/9/0",
		},
		{
			Name: "pellet_consumption_since_bulk_fill",
			Oid:  "/1/60/0/23/100/0",
		},
		{
			Name: "pellet_consumption_total",
			Oid:  "/1/60/0/23/103/0",
		},
		{
			Name: "burner_operation_state",
			Oid:  "/1/60/0/2/1/0",
		},
		{
			Name: "buffer_temp_tpe",
			Oid:  "/1/16/1/21/65/0",
		},
		{
			Name: "buffer_temp_tpa",
			Oid:  "/1/16/1/21/66/0",
		},
		{
			Name: "buffer_temp_set_point",
			Oid:  "/1/16/1/1/15/0",
		},
		{
			Name: "burner_daily_store",
			Oid:  "1/60/0/43/42/0",
		},
		{
			Name: "burner_current_power_level",
			Oid:  "/1/60/0/0/9/0",
		},
	}
}

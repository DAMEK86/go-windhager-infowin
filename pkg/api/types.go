package api

import (
	"encoding/json"
	"fmt"
	"time"
)

type Timestamp struct {
	time.Time
}

// UnmarshalJSON decodes non RFC based timestamp into a time.Time object
func (p *Timestamp) UnmarshalJSON(bytes []byte) error {
	var value string
	err := json.Unmarshal(bytes, &value)

	if err != nil {
		fmt.Printf("error decoding timestamp: %s\n", err)
		return err
	}

	*&p.Time, err = time.Parse("2006-01-02 15:04:05", value)
	return err
}

type DatapointRequest struct {
	Name string
	Oid  string
}

type Datapoint struct {
	Oid            string    `json:"OID"`
	GroupNr        int       `json:"groupNr"`
	MaxValue       string    `json:"maxValue"`
	MemberNr       int       `json:"memberNr"`
	MinValue       string    `json:"minValue"`
	Name           string    `json:"name"`
	Step           string    `json:"step"`
	StepId         int       `json:"stepId"`
	SubtypeId      int       `json:"subtypeId"`
	DateTime       Timestamp `json:"timestamp"`
	TypeId         int       `json:"typeId"`
	Unit           string    `json:"unit"`
	UnitId         int       `json:"unitId"`
	Value          string    `json:"value"`
	WriteProtected bool      `json:"writeProt"`
}

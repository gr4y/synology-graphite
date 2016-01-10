package plugins

import (
	g "github.com/soniah/gosnmp"
	"log"
)

type LoadPlugin struct{}

func (p LoadPlugin) FetchData(snmp g.GoSNMP) map[string]float64 {
	result, err := snmp.Get([]string{".1.3.6.1.4.1.2021.10.1.5.1", ".1.3.6.1.4.1.2021.10.1.5.2", ".1.3.6.1.4.1.2021.10.1.5.3"})
	if err != nil {
		log.Fatalf("Get() err: %v", err)
	}
	return map[string]float64{
		"load.shortterm": float64(result.Variables[0].Value.(int)) / 100,
		"load.midterm":   float64(result.Variables[1].Value.(int)) / 100,
		"load.longterm":  float64(result.Variables[2].Value.(int)) / 100,
	}
}

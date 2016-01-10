package plugins

import (
	g "github.com/soniah/gosnmp"
	"log"
)

type CPUPlugin struct{}

func (p CPUPlugin) FetchData(snmp g.GoSNMP) map[string]float64 {
	oids := []string{
		".1.3.6.1.4.1.2021.11.50.0",
		".1.3.6.1.4.1.2021.11.51.0",
		".1.3.6.1.4.1.2021.11.52.0",
		".1.3.6.1.4.1.2021.11.53.0",
		".1.3.6.1.4.1.2021.11.54.0",
		".1.3.6.1.4.1.2021.11.55.0",
		".1.3.6.1.4.1.2021.11.56.0",
	}
	result, err := snmp.Get(oids)
	if err != nil {
		log.Fatalf("Get() err: %v", err)
	}
	return map[string]float64{
		"cpu-0.cpu-user":      float64(result.Variables[0].Value.(uint)),
		"cpu-0.cpu-nice":      float64(result.Variables[1].Value.(uint)),
		"cpu-0.cpu-system":    float64(result.Variables[2].Value.(uint)),
		"cpu-0.cpu-idle":      float64(result.Variables[3].Value.(uint)),
		"cpu-0.cpu-wait":      float64(result.Variables[4].Value.(uint)),
		"cpu-0.cpu-kernel":    float64(result.Variables[5].Value.(uint)),
		"cpu-0.cpu-interrupt": float64(result.Variables[6].Value.(uint)),
	}
}

package plugins

import (
	"fmt"
	g "github.com/soniah/gosnmp"
	"log"
)

type DiskPlugin struct{}

func (p DiskPlugin) FetchData(snmp g.GoSNMP) map[string]float64 {
	metrics := map[string]float64{}
	for key, value := range getTemperatures(snmp) {
		metrics[fmt.Sprintf("disk.disk-%v.temperature", key)] = value
	}
	getDiskUsage(snmp)
	return metrics
}

func getDiskUsage(snmp g.GoSNMP) map[int64]float64 {
	result, err := snmp.GetBulk([]string{".1.3.6.1.2.1.25.2.3.1.1"}, 64, 64)
	if err != nil {
		log.Fatalf("Get() err: %v", err)
	}
	log.Println(result.Variables)
	return map[int64]float64{}
}

func getTemperatures(snmp g.GoSNMP) map[int]float64 {
	result, err := snmp.Get([]string{".1.3.6.1.4.1.6574.2.1.1.6.0", ".1.3.6.1.4.1.6574.2.1.1.6.1"})
	if err != nil {
		log.Fatalf("Get() err: %v", err)
	}
	temps := map[int]float64{}
	for i, variable := range result.Variables {
		temps[i] = float64(variable.Value.(int))
	}
	return temps
}

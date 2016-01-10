package plugins

import (
	g "github.com/soniah/gosnmp"
)

type Plugin interface {
	FetchData(snmp g.GoSNMP) map[string]float64
}

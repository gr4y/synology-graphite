package cmd

import (
	"fmt"
	"github.com/codegangsta/cli"
	p "github.com/gr4y/synology-graphite/plugins"
	g "github.com/soniah/gosnmp"
	"log"
	"net"
	"time"
)

var CmdFetchData = func(c *cli.Context) {
	diskStation := c.String("ds")
	interval := c.Duration("interval")
	prefix := c.String("prefix")

	carbonHost := c.String("carbon-host")
	carbonPort := c.String("carbon-port")
	carbonAddr := fmt.Sprintf("%s:%s", carbonHost, carbonPort)

	// SNMP Configuration
	snmp := g.GoSNMP{
		Target:    diskStation,
		Port:      161,
		Community: "public",
		Version:   g.Version2c,
		Timeout:   time.Duration(2) * time.Second,
	}

	// Plugins
	plugins := []p.Plugin{p.DiskPlugin{}, p.LoadPlugin{}, p.CPUPlugin{}} //, p.MemoryPlugin{}, p.NetworkPlugin{}}

	// Execute plugins every interval
	for now := range time.Tick(interval) {
		snmp.Connect()
		defer snmp.Conn.Close()
		for _, plugin := range plugins {
			data := plugin.FetchData(snmp)
			for key, value := range data {
				metric := fmt.Sprintf("%s.%s %v %d\n\r", prefix, key, value, now.Unix())
				send(metric, carbonAddr)
			}
		}

	}

}

func send(metric string, carbonAddr string) {
	conn, err := net.Dial("tcp", carbonAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(conn, metric)
	conn.Close()
}

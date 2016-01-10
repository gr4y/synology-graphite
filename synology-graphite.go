package main

import (
	"github.com/codegangsta/cli"
	"github.com/gr4y/synology-graphite/cmd"
	"os"
	"time"
)

func main() {
	a := cli.NewApp()
	a.Name = "synology-graphite"
	a.Usage = "Gets SNMP Data from Synology DiskStation, converts and sends it into Graphite"
	a.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "diskstation,ds",
			Usage: "DiskStation Hostname or IP address",
		},
		cli.StringFlag{
			Name:  "carbon-host,ch",
			Value: "localhost",
			Usage: "Carbon Hostname or IP address",
		},
		cli.StringFlag{
			Name:  "carbon-port,cp",
			Value: "2003",
			Usage: "Carbon Port",
		},
		cli.DurationFlag{
			Name:  "interval,i",
			Value: 60 * time.Second,
			Usage: "Interval",
		},
		cli.StringFlag{
			Name:  "prefix",
			Value: "metrics.nas",
			Usage: "Prefix",
		},
	}
	a.Action = cmd.CmdFetchData
	a.Run(os.Args)
}

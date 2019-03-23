package main

import (
	a "github.com/mrvocalplay/homation/automation"
	"github.com/mrvocalplay/homation/log"
)

func main() {
	a.Temperatures()
	a.Volts()
	a.Clock()

	log.Log(log.MakeLog(a.Temperatures()))
	// rpi.GetCPUTemp()
	// rpi.GetCPUSpannung()
	// log.Log("tet")
}

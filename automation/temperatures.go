package automation

import (
	"fmt"

	"github.com/mrvocalplay/homation/monitoring/rpi"
)

// PowerOff Kill Server
func PowerOff() {
	fmt.Println("SHUTTING DOWN")
}

// Warn Warn Message
func Warn(currentTemp string) {
	fmt.Println("Temperatures Critical: " + currentTemp)
}

func Temperatures() {
	currentTemp := rpi.GetCPUTemp()
	if rpi.CutStr(currentTemp, "temp=", "'C") > 75 {
		PowerOff()
	} else if currentTemp > 65 {
		Warn(currentTemp)
	}

	// log.Log(log.MakeLog(rpi.GetCPUTemp(), rpi.GetCPUSpannung()))
	// rpi.GetCPUTemp()
	// rpi.GetCPUSpannung()
	// log.Log("tet")
}

package automation

import (
	"fmt"

	mon "github.com/mrvocalplay/monitoring"
	"github.com/mrvocalplay/monitoring/rpi"
)

// PowerOff Kill Server
func PowerOff() {
	fmt.Println("SHUTTING DOWN")
}

// Warn Warn Message
func Warn(currentTemp string) {
	fmt.Println("Temperatures Critical: " + currentTemp)
}

func Temperatures() string {
	currentTemp := mon.CutStr(rpi.GetCPUTemp(), "temp=", "'C")
	fmt.Println(currentTemp)
	return currentTemp
}

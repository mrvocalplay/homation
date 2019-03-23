package automation

import (
	"fmt"

	mon "github.com/mrvocalplay/monitoring"
)

// PowerOff Kill Server
func PowerOff() {
	fmt.Println("SHUTTING DOWN")
	// fmt.Println(mon.RunCmd("sudo shutdown -h now"))
}

// Warn Warn Message
func Warn(currentTemp string) {
	fmt.Println("Temperatures Critical: " + currentTemp)
}

func Temperatures(currentTemp string) {
	temp := mon.ToInt(mon.CutStr(currentTemp, "temp=", "'C"))
	if temp >= 80 {
		PowerOff()
	} else if temp >= 50 {
		Warn(currentTemp)
	}
}

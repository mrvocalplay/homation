package automation

import (
	"fmt"

	"github.com/mrvocalplay/monitoring/rpi"
)

// PowerOff Kill Server
func PowerOff() {
	fmt.Println("SHUTTING DOWN")
}

// Warn Warn Message
// func Warn() {
// 	fmt.Println("Temperatures Critical: " + rpi.GetTemp())
// }

func Temp() {
	// currentTemp := rpi.GetTemp()
	// if currentTemp > 75 {
	// 	PowerOff()
	// } else if currentTemp > 65 {
	// 	Warn()
	// }

	rpi.GetTemp()
}

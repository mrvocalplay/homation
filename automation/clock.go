package automation

import (
	"fmt"

	mon "github.com/mrvocalplay/monitoring"
	"github.com/mrvocalplay/monitoring/rpi"
)

func Clock() {
	currentArmClock := mon.CutStr(rpi.GetArmClock(), "frequency(45)=", "")
	fmt.Println(currentArmClock)
}

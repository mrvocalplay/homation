package automation

import (
	"fmt"

	mon "github.com/mrvocalplay/monitoring"

	"github.com/mrvocalplay/monitoring/rpi"
)

func Volts() {
	currentCPUVolts := mon.CutStr(rpi.GetCPUVolts(), "volt=", "V")
	fmt.Println(currentCPUVolts)
}

package main

import (
	"fmt"
	"strings"

	a "github.com/mrvocalplay/homation/automation"
	"github.com/mrvocalplay/monitoring/rpi"
)

func main() {
	os := strings.TrimRight(rpi.GetOs(), "\r\n")
	temp := strings.TrimRight(rpi.GetCPUTemp(), "\r\n")
	volts := strings.TrimRight(rpi.GetCPUVolts(), "\r\n")
	clock := strings.TrimRight(rpi.GetArmClock(), "\r\n")
	str := (os + ", " + temp + ", " + volts + ", " + clock)
	fmt.Println(str)
	a.Temperatures(temp)
	// notify.EMail(str)

	// log.Log(log.MakeLog(a.Temperatures()))
}

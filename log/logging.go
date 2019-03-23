package log

import (
	"log"
	"os"

	"github.com/mrvocalplay/monitoring/rpi"
)

func Log(message string) {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print(message)
}

func MakeLog(CPUTemp string, CPUSpannung string) string {
	os := rpi.GetOs()
	log := (os + " " + CPUTemp + " " + CPUSpannung)
	return log
}

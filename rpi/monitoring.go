package rpi

import (
	"fmt"
	"os/exec"
	"runtime"
)

// GetTemp returns the CPU and GPU Temperature from the RPI
func GetTemp() {
	// temperature := "CPUTEMP"
	temperature, err := exec.Command("vcgencmd measure_temp").Output()
	if err != nil {
		fmt.Println("error occured")
		fmt.Printf("%s", err)
	}
	fmt.Println(temperature)
	// return temperature
}

// GetOs gives you your OS and Arcitecture
func GetOs() string {
	os := runtime.GOOS
	arch := runtime.GOARCH
	info := os + " " + arch
	return info
}

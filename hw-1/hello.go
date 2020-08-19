package main

import (
	"fmt" // Using fmt module
	"os"  // Using os module

	"github.com/beevik/ntp" // Using NTP
)

func main() {
	time, err := ntp.Time("time.nist.gov")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	const layout = "3:04:05 PM (MST)"
	fmt.Println("Current Local Time:")
	fmt.Println(time.Local().Format(layout))
	os.Exit(0)
}

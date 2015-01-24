package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	MIN  = 60
	HOUR = MIN * 60
	DAY  = HOUR * 24
)

func main() {
	// duration is time in milliseconds
	if len(os.Args) != 2 || os.Args[1] == "-h" || os.Args[1] == "--help" || os.Args[1] == "-help" {
		fmt.Printf("usage: %s [-h] <milliseconds>\n", os.Args[0])
		os.Exit(1)
	}
	duration, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%s\n\nusage: %s [-h] <milliseconds>\n", err, os.Args[0])
		os.Exit(1)
	}
	seconds := float64(duration) / 1000

	if seconds < 1 {
		fmt.Printf("%dms\n", duration)
	} else if seconds < MIN {
		fmt.Printf("%.1fs\n", seconds)
	} else if seconds < HOUR {
		fmt.Printf("%dm %ds\n", int(seconds)/MIN, int(seconds)%MIN)
	} else if seconds < DAY {
		if int(seconds)%MIN >= 30 {
			// round up to nearest minute
			seconds += MIN
		}
		fmt.Printf("%dh %dm\n", int(seconds)/HOUR, int(seconds)%HOUR/MIN)
	} else {
		if int(seconds)%MIN >= 30 {
			// round up to nearest minute
			seconds += MIN
		}
		fmt.Printf("%dd %dh %dm\n", int(seconds)/DAY, int(seconds)%DAY/HOUR, int(seconds)%HOUR/MIN)
	}
}
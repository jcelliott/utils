package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/cloudfoundry/gosigar"
	"github.com/wsxiaoys/terminal"
)

func main() {
	terminal.Stdout.Color("b").Print("Uptime:\t ").Reset().Print(uptime()).Nl()
	terminal.Stdout.Color("b").Print("Loadavg: ").Reset().Print(loadavg()).Nl()
	terminal.Stdout.Color("b").Print("Root:\t ").Reset().Print(rootUsage()).Nl()
}

func uptime() string {
	uptime := sigar.Uptime{}
	uptime.Get()
	return uptime.Format()
}

func loadavg() string {
	concreteSigar := sigar.ConcreteSigar{}
	avg, err := concreteSigar.GetLoadAverage()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return "error"
	}
	return fmt.Sprintf("%.2f, %.2f, %.2f", avg.One, avg.Five, avg.Fifteen)
}

func rootUsage() string {
	usage := sigar.FileSystemUsage{}
	usage.Get("/")
	return fmt.Sprintf("%s / %s (%.0f%%)",
		strings.TrimSpace(sigar.FormatSize(usage.Used*1024)),
		strings.TrimSpace(sigar.FormatSize(usage.Total*1024)),
		usage.UsePercent(),
	)
}

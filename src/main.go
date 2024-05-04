package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

func FormatDuration(totalSeconds int) string {
	seconds := totalSeconds % 60
	minutes := (totalSeconds / 60) % 60
	hours := (totalSeconds / 3600) % 24
	days := (totalSeconds / 86400) % 30
	months := (totalSeconds / 2592000) % 12
	years := totalSeconds / 31536000
	weeks := (totalSeconds / 604800) % 52

	result := ""

	if years > 0 {
		result += fmt.Sprintf("%d years, ", years)
	}
	if months > 0 {
		result += fmt.Sprintf("%d months, ", months)
	}
	if weeks > 0 {
		result += fmt.Sprintf("%d weeks, ", weeks)
	}
	if days > 0 {
		result += fmt.Sprintf("%d days, ", days)
	}
	if hours > 0 {
		result += fmt.Sprintf("%d hours, ", hours)
	}
	if minutes > 0 {
		result += fmt.Sprintf("%d minutes, ", minutes)
	}
	result += fmt.Sprintf("%d seconds", seconds) // Seconds are always shown

	return result
}

func FetchSystem() (string, string, string, string, string) {
	hostInfo, hostErr := host.Info()
	cpuInfo, cpuErr := cpu.Info()
	memInfo, memError := mem.VirtualMemory()

	sysUptime := FormatDuration(int(hostInfo.Uptime))

	if hostErr != nil {
		fmt.Println("Error fetching host info:", hostErr)
	}

	if cpuErr != nil {
		fmt.Println("Error fetching CPU info:", cpuErr)
	}

	if memError != nil {
		fmt.Println("Error fetching memory info:", cpuErr)
	}

	os := fmt.Sprintf("OS: %s %s", hostInfo.Platform, hostInfo.KernelArch)
	kernel := fmt.Sprintf("Kernel: %s", hostInfo.KernelVersion)
	uptime := fmt.Sprintf("Uptime: %s", sysUptime)
	cpu := fmt.Sprintf("CPU: %s (%s)", cpuInfo[0].ModelName, fmt.Sprint(cpuInfo[0].Cores))
	memory := fmt.Sprintf("Memory: %sMiB / %sMiB", fmt.Sprint(memInfo.Used/1048576), fmt.Sprint(memInfo.Total/1048576))

	return os, kernel, uptime, cpu, memory
}

func main() {
	os, kernel, uptime, cpu, memory := FetchSystem()

	fmt.Println(os)
	fmt.Println(kernel)
	fmt.Println(uptime)
	fmt.Println(cpu)
	fmt.Println(memory)
}

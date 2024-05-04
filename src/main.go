package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

func FetchSystem() (string, string, string, string, string) {
	hostInfo, hostErr := host.Info()
	cpuInfo, cpuErr := cpu.Info()
	memInfo, memError := mem.VirtualMemory()

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
	uptime := fmt.Sprintf("Uptime: %s seconds", fmt.Sprint(hostInfo.Uptime))
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

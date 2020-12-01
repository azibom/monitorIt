package cpuService

import "github.com/shirou/gopsutil/cpu"

func GetCpuInfoInArray() []float64 {
	percentage, _ := cpu.Percent(0, true)
	return percentage
}

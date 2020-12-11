package systemInfoService

import (
	"monitoring/internal/package/services/cpuService"
	"monitoring/internal/package/services/memoryService"
)

type SystemInfo struct {
	CPU    []float64
	MEMORY float64
}

func GetSystemResourceInfo() SystemInfo {
	percentage := cpuService.GetCpuInfoInArray()
	swapMemory := memoryService.GetMemoryInfo()

	systemInfo := SystemInfo{percentage, swapMemory.UsedPercent}
	return systemInfo
}

package memoryService

import "github.com/shirou/gopsutil/mem"

func GetMemoryInfo() *mem.SwapMemoryStat {
	swapMemory, _ := mem.SwapMemory()
	return swapMemory
}

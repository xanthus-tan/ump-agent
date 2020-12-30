package metrics

import (
	"github.com/shirou/gopsutil/v3/mem"
)

// MemStat 内存状态
type MemStat struct {
	Total     uint64 `json:"total"`
	Available uint64 `json:"available"`
	// RAM used by programs
	//
	// This value is computed from the kernel specific values.
	Used uint64 `json:"used"`
	// Percentage of RAM used by programs
	//
	// This value is computed from the kernel specific values.
	UsedPercent float64 `json:"usedPercent"`
	// This is the kernel's notion of free memory; RAM chips whose bits nobody
	// cares about the value of right now. For a human consumable number,
	// Available is what you really want.
	Free      uint64 `json:"free"`
	SwapTotal uint64 `json:"swapTotal"`
	SwapFree  uint64 `json:"swapFree"`
}

// GetMemState 获取当前内存状态
func GetMemState() (*MemStat, error) {
	v, err := mem.VirtualMemory()
	msg := new(MemStat)
	msg.Total = v.Total
	msg.Available = v.Available
	msg.Used = v.Used
	msg.UsedPercent = v.UsedPercent
	msg.Free = v.Free
	msg.SwapTotal = v.SwapTotal
	msg.SwapFree = v.SwapFree
	return msg, err
}

// GetMemorySize 获取内存容量
func (l *Linux) GetMemorySize() uint64 {
	memoryStat, err := GetMemState()
	if err != nil {
		return 0
	}
	return memoryStat.Total
}

package metrics

import (
	"ump-agent/message"

	"github.com/shirou/gopsutil/v3/mem"
)

// MemStat 内存状态
type memStat struct {
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

// MemMsg 内存数据项
type MemMsg struct {
	Header  message.Header `json:"header"`
	MemStat memStat        `json:"body"`
}

// GetMemState 获取当前内存状态
func GetMemState() (*MemMsg, error) {
	v, err := mem.VirtualMemory()
	msg := new(MemMsg)
	msg.Header.MsgType = message.TYPEMETRICS
	msg.Header.Item = message.ITEMMEM
	msg.MemStat.Total = v.Total
	msg.MemStat.Available = v.Available
	msg.MemStat.Used = v.Used
	msg.MemStat.UsedPercent = v.UsedPercent
	msg.MemStat.Free = v.Free
	msg.MemStat.SwapTotal = v.SwapTotal
	msg.MemStat.SwapFree = v.SwapFree
	return msg, err
}

// GetMemorySize 获取内存容量
func (l *Linux) GetMemorySize() uint64 {
	memoryStat, err := GetMemState()
	if err != nil {
		return 0
	}
	return memoryStat.MemStat.Total
}

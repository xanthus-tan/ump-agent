package metrics

import (
	"runtime"

	"github.com/shirou/gopsutil/v3/cpu"
)

// CPUMertics cpu指标
type CPUMertics struct {
	CPUStats []cpu.TimesStat `json:"cpuStats"`
}

// GetCPUTotalStat 获取当前CPU Times
func GetCPUTotalStat() (*CPUMertics, error) {
	msg := new(CPUMertics)
	timestat, err := cpu.Times(false)
	msg.CPUStats = timestat
	return msg, err
}

// GetCpus 获取cpu核心数
func (l *Linux) GetCpus() int {
	return runtime.NumCPU()
}

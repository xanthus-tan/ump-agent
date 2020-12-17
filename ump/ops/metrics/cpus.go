package metrics

import (
	"runtime"
	"ump-agent/message"

	"github.com/shirou/gopsutil/v3/cpu"
)

// CPUMsg cpu消息体
type CPUMsg struct {
	Header   message.Header  `json:"header"`
	CPUStats []cpu.TimesStat `json:"body"`
}

// GetCPUTotalStat 获取当前CPU Times
func GetCPUTotalStat() (*CPUMsg, error) {
	msg := new(CPUMsg)
	msg.Header.MsgType = message.TYPEMETRICS
	msg.Header.Item = message.ITEMCPU
	timestat, err := cpu.Times(false)
	msg.CPUStats = timestat
	return msg, err
}

// GetCpus 获取cpu核心数
func (l *Linux) GetCpus() int {
	return runtime.NumCPU()
}

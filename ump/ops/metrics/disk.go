package metrics

import (
	"ump-agent/message"

	"github.com/shirou/gopsutil/disk"
)

// DiskStat 简要硬盘状态信息
type diskStat struct {
	Fstype      string  `json:"fstype"`
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

// DiskMsg 本地磁盘消息体
type DiskMsg struct {
	Header   message.Header `json:"header"`
	DiskStat diskStat       `json:"body"`
}

// GetDiskStat 获取分区磁盘状态
func GetDiskStat() (*DiskMsg, error) {
	msg := new(DiskMsg)
	v, err := disk.Usage("/")
	msg.Header.MsgType = message.TYPEMETRICS
	msg.Header.Item = message.ITEMDISK
	msg.DiskStat.Total = v.Total
	msg.DiskStat.Fstype = v.Fstype
	msg.DiskStat.Free = v.Free
	msg.DiskStat.Used = v.Used
	msg.DiskStat.UsedPercent = v.UsedPercent
	return msg, err
}

// DiskPartitions 硬盘分区
func DiskPartitions() ([]disk.PartitionStat, error) {
	return disk.Partitions(false)
}

// GetDiskSize 获取磁盘容量
func (l *Linux) GetDiskSize() uint64 {
	pList, err := DiskPartitions()
	if err != nil {
		return 0
	}
	var diskSize uint64
	for _, p := range pList {
		usageStat, err := disk.Usage(p.Mountpoint)
		if err != nil {
			return 0
		}
		diskSize = diskSize + usageStat.Total
	}
	return diskSize
}

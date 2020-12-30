package metrics

import (
	"github.com/shirou/gopsutil/disk"
)

// DiskStat 简要硬盘状态信息
type DiskStat struct {
	Fstype      string  `json:"fstype"`
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

// GetDiskStat 获取分区磁盘状态
func GetDiskStat() (*DiskStat, error) {
	msg := new(DiskStat)
	v, err := disk.Usage("/")

	msg.Total = v.Total
	msg.Fstype = v.Fstype
	msg.Free = v.Free
	msg.Used = v.Used
	msg.UsedPercent = v.UsedPercent
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

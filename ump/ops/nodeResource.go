package ops

import (
	"encoding/json"
	"os"
	"runtime"
	"ump-agent/message"
	"ump-agent/ump/ops/metrics"
)

// resource 节点资源
type resource struct {
	Cpus       int    `json:"cpus"`
	MemorySize uint64 `json:"memSize"`
	DiskSize   uint64 `json:"diskSize"`
}

// Computer 计算机基本信息
type computer struct {
	HostName string `json:"hostname"`
	Arch     string `json:"arch"`
	Platform string `json:"os"`
}

// GeneralMsg 主机通用消息体
type GeneralMsg struct {
	Header   message.Header `json:"header"`
	Computer computer       `json:"computer"`
	Resource resource       `json:"resource"`
}

// OpsResource 基本信息接口
type OpsResource interface {
	GetCpus() int          // CPU核数
	GetMemorySize() uint64 // 内存容量
	GetDiskSize() uint64   // 本地硬盘容量
}

// GeneralInfo 通用资源
func generalInfo() *GeneralMsg {
	var b OpsResource
	b = new(metrics.Linux)
	cpus := b.GetCpus()
	memorySize := b.GetMemorySize()
	diskSize := b.GetDiskSize()
	arch := runtime.GOARCH
	hostname, _ := os.Hostname()
	msg := new(GeneralMsg)
	msg.Header.MsgType = message.TYPENODE
	msg.Header.Item = message.ITEMGENERAL
	msg.Computer.Platform = runtime.GOOS
	msg.Computer.HostName = hostname //主机名称
	msg.Computer.Arch = arch
	msg.Resource.Cpus = cpus
	msg.Resource.MemorySize = memorySize
	msg.Resource.DiskSize = diskSize
	return msg
}

// CollectInfo 信息采集
func CollectInfo(metricsType string) (string, error) {
	var s interface{}
	switch metricsType {
	case "cpu":
		s, _ = metrics.GetCPUTotalStat()
	case "memory":
		s, _ = metrics.GetMemState()
	case "disk":
		s, _ = metrics.GetDiskStat()
	}
	j, err := json.Marshal(s)
	return string(j), err
}

// CollectGeneralInfo 收集主机通用信息方法
func CollectGeneralInfo() string {
	c := generalInfo()
	j, _ := json.Marshal(c)
	return string(j)
}

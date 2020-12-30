package ops

import (
	"encoding/json"
	"os"
	"runtime"
	"ump-agent/ump/header"
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
	Header   header.Header `json:"header"`
	Computer computer      `json:"computer"`
	Resource resource      `json:"resource"`
}

// MetricsMsgPkg 指标消息体
type MetricsMsgPkg struct {
	Header header.Header `json:"header"`
	Body   interface{}   `json:"body"`
}

// OpsResource 基本信息接口
type OpsResource interface {
	GetCpus() int          // CPU核数
	GetMemorySize() uint64 // 内存容量
	GetDiskSize() uint64   // 本地硬盘容量
}

// GeneralInfo 通用资源
func generalInfo(code string) *GeneralMsg {
	var b OpsResource
	b = new(metrics.Linux)
	cpus := b.GetCpus()
	memorySize := b.GetMemorySize()
	diskSize := b.GetDiskSize()
	arch := runtime.GOARCH
	hostname, _ := os.Hostname()
	msg := new(GeneralMsg)
	msg.Header.MsgType = header.TYPENODE
	msg.Header.Item = header.ITEMGENERAL
	msg.Header.ActionCode = code
	msg.Computer.Platform = runtime.GOOS
	msg.Computer.HostName = hostname //主机名称
	msg.Computer.Arch = arch
	msg.Resource.Cpus = cpus
	msg.Resource.MemorySize = memorySize
	msg.Resource.DiskSize = diskSize
	return msg
}

// CollectInfo 信息采集
func CollectInfo(metricsType string) (*MetricsMsgPkg, error) {
	var s interface{}
	var err error
	metricsMsgPkg := new(MetricsMsgPkg)
	metricsMsgPkg.Header.MsgType = header.TYPEMETRICS
	switch metricsType {
	case "cpu":
		s, err = metrics.GetCPUTotalStat()
		metricsMsgPkg.Header.Item = header.ITEMCPU
	case "memory":
		s, err = metrics.GetMemState()
		metricsMsgPkg.Header.Item = header.ITEMMEM
	case "disk":
		s, err = metrics.GetDiskStat()
		metricsMsgPkg.Header.Item = header.ITEMDISK
	}
	metricsMsgPkg.Body = s
	return metricsMsgPkg, err
}

// CollectGeneralInfo 收集主机通用信息方法
func CollectGeneralInfo(code string) string {
	c := generalInfo(code)
	j, _ := json.Marshal(c)
	return string(j)
}

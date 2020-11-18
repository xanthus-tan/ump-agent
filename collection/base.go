package collection

import (
	"os"
	"runtime"
)

// Computer 计算机基本信息
type Computer struct {
	HostName string
	Arch     string
	Platform string
	Cpus     int
	Memory   int
	Disk     int
}

// CollectBaseINFO 收集主机基本信息方法
func CollectBaseINFO() *Computer {
	c := new(Computer)
	arch := runtime.GOARCH
	platform := runtime.GOOS
	hostname, _ := os.Hostname()
	cpus := runtime.NumCPU() //cpu核心数
	c.HostName = hostname    //主机名称
	c.Cpus = cpus
	c.Arch = arch
	c.Platform = platform
	return c
}

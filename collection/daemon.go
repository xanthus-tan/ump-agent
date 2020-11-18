package collection

import "fmt"

//Daemon 守护进程运行
func Daemon() {
	fmt.Println("run...daemon")
	c := CollectBaseINFO()
	fmt.Println((*c).HostName)
	fmt.Println((*c).Cpus)
	fmt.Println((*c).Arch)
	fmt.Println((*c).Platform)
}

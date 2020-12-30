package collector

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"regexp"
	"time"
	"ump-agent/message"
	"ump-agent/ump/ops"
)

// ConsoleServer console服务器
type ConsoleServer struct {
	Addr string
	Port string
}
type nodeTime struct {
	NodeTime string `json:"time"`
}

// heartBeat 心跳
type heartBeat struct {
	Header message.Header `json:"header"`
	Body   nodeTime       `json:"body"`
}

// Work 采集器
func Work() {
	server := new(ConsoleServer)
	server.Addr = "192.168.178.18"
	server.Port = "5388"
	initTicker := time.NewTicker(5 * time.Second)
	heartBeatCode := message.INITCODE
	var replay string
	var connError error
	for range initTicker.C {
		scn := time.Now().Unix()
		replay, connError = server.heartBeatToConsole(heartBeatCode, scn)
		re := regexp.MustCompile("\r*\n")
		replay = re.ReplaceAllString(replay, "")
		log.Println("replay=>" + replay)
		if connError != nil {
			fmt.Println(connError)
		}
		if replay == message.REGISTEREQUESTCODE {
			log.Print("node 注册.....")
			server.registerNodeInfo()
			continue
		} else if replay == message.REGISTERSUCCESSCODE {
			heartBeatCode = message.HEARTBEATCODE
		} else if replay == message.CONSOLEACTIVE {
			server.nodeMetrics(scn)
			heartBeatCode = message.HEARTBEATCODE
		}

	}
}

// HeartBeat 心跳
func (server *ConsoleServer) heartBeatToConsole(code string, scn int64) (string, error) {
	connAddr := server.Addr + ":" + server.Port
	conn, connerr := net.Dial("tcp", connAddr)
	if connerr != nil {
		fmt.Println(connerr)
	}
	defer conn.Close()
	nodeTime := time.Now().Format("2006-01-02 15:04:05")
	h := new(heartBeat)
	h.Header.MsgType = message.TYPEPROBE
	h.Header.Scn = scn
	h.Header.ActionCode = code
	h.Body.NodeTime = nodeTime
	j, _ := json.Marshal(h)
	msg := string(j)
	fmt.Fprintln(conn, msg)
	message, _ := bufio.NewReader(conn).ReadString('\n')
	return message, connerr
}

// nodeMetrics 采集指标信息
func (server *ConsoleServer) nodeMetrics(scn int64) error {
	connAddr := server.Addr + ":" + server.Port
	conn, err := net.Dial("tcp", connAddr)
	defer conn.Close()
	cpuMsgPkg, _ := ops.CollectInfo("cpu")
	cpuMsgPkg.Header.Scn = scn
	cpuMetrics, _ := json.Marshal(cpuMsgPkg)
	fmt.Fprintln(conn, string(cpuMetrics))
	memMsgPkg, _ := ops.CollectInfo("memory")
	memMsgPkg.Header.Scn = scn
	memMetrics, _ := json.Marshal(memMsgPkg)
	fmt.Fprintln(conn, string(memMetrics))
	diskMsgPkg, _ := ops.CollectInfo("disk")
	diskMsgPkg.Header.Scn = scn
	diskMetrics, _ := json.Marshal(diskMsgPkg)
	fmt.Fprintln(conn, string(diskMetrics))
	return err
}

// nodeInfo 发送节点信息
func (server *ConsoleServer) registerNodeInfo() error {
	connAddr := server.Addr + ":" + server.Port
	conn, err := net.Dial("tcp", connAddr)
	defer conn.Close()
	fmt.Fprintln(conn, ops.CollectGeneralInfo(message.REGISTEREQUESTCODE))
	return err
}

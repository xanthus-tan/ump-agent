package collector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
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
	hbTicker := time.NewTicker(5 * time.Second)
	for range hbTicker.C {
		server.heartBeatToConsole(message.INITACTION)
		// server.nodeMetrics()
	}
}

// HeartBeat 心跳
func (server *ConsoleServer) heartBeatToConsole(code string) error {
	connAddr := server.Addr + ":" + server.Port
	conn, connerr := net.Dial("tcp", connAddr)
	if connerr != nil {
		fmt.Println(connerr)
	}
	defer conn.Close()
	scn := time.Now().Unix()
	nodeTime := time.Now().Format("2006-01-02 15:04:05")
	h := new(heartBeat)
	h.Header.MsgType = message.TYPEPROBE
	h.Header.Scn = scn
	h.Header.ActionCode = code
	h.Body.NodeTime = nodeTime
	j, _ := json.Marshal(h)
	msg := string(j)
	fmt.Fprintln(conn, msg)
	var buf bytes.Buffer
	_, err := io.Copy(&buf, conn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(buf.String())
	return connerr
}

// nodeMetrics 采集指标信息
func (server *ConsoleServer) nodeMetrics() error {
	connAddr := server.Addr + ":" + server.Port
	conn, err := net.Dial("tcp", connAddr)
	defer conn.Close()
	cpuMsg, _ := ops.CollectInfo("cpu")
	fmt.Fprintln(conn, cpuMsg)
	memMsg, _ := ops.CollectInfo("memory")
	fmt.Fprintln(conn, memMsg)
	diskMsg, _ := ops.CollectInfo("disk")
	fmt.Fprintln(conn, diskMsg)
	return err
}

// nodeInfo 发送节点信息
func (server *ConsoleServer) nodeInfo() error {
	connAddr := server.Addr + ":" + server.Port
	conn, err := net.Dial("tcp", connAddr)
	defer conn.Close()
	fmt.Fprintln(conn, ops.CollectGeneralInfo())
	return err
}

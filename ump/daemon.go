package ump

import (
	"log"
	"ump-agent/ump/net/collector"
	"ump-agent/ump/net/listener"

	"github.com/sevlyar/go-daemon"
)

//Daemon 守护进程运行
func Daemon() {
	log.Print("- - - - - - - - - - - - - - -")
	log.Print("daemon started")
	log.Print("- - - - - - - - - - - - - - -")
	cntxt := &daemon.Context{
		PidFileName: "agent.pid",
		PidFilePerm: 0644,
		LogFileName: "agent.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{"[ump-agent]"},
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release()
	go collector.Work()
	listener.ActonListener()
}

package main

import (
	"base/env"
	"base/glog"
	"base/service"
	"flag"
	"time"
)

type RoomServer struct {
	service.Service
	id uint64
}

var gRoomserver *RoomServer

func RoomServer_GetMe() *RoomServer {
	if gRoomserver == nil {
		gRoomserver = &RoomServer{}
		gRoomserver.Derived = gRoomserver
	}

	if !ConfigMgr_GetMe().Init() {
		return nil
	}

	return gRoomserver
}

func (this *RoomServer) Init() bool {

	if !StartHttpServer() {
		return false
	}
	return true
}

func (this *RoomServer) MainLoop() {
	time.Sleep(time.Second * 3)
}

func (this *RoomServer) Final() bool {

	return true
}

func (this *RoomServer) Reload() bool {

	return true
}

func (this *RoomServer) GetId() uint64 {
	return this.id
}

var (
	logfile = flag.String("logfile", "", "Log file name")
	config  = flag.String("config", "config.json", "config path")
)

func main() {

	flag.Parse()

	env.Load(*config)

	loglevel := env.Get("global", "loglevel")
	if loglevel != "" {
		flag.Lookup("stderrthreshold").Value.Set(loglevel)
	}

	logtostderr := env.Get("global", "logtostderr")
	if logtostderr != "" {
		flag.Lookup("logtostderr").Value.Set(logtostderr)
	}

	if *logfile != "" {
		glog.SetLogFile(*logfile)
	} else {
		glog.SetLogFile(env.Get("room", "log"))
	}

	RoomServer_GetMe().Main()
}

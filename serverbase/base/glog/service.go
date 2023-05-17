package glog

import (
	"flag"
)

func NewService(logfile, level, tostd string) *GlogService {
	return &GlogService{file: logfile, level: level, tostd: tostd}
}

type GlogService struct {
	file  string
	level string
	tostd string
}

func (this *GlogService) Init() bool {
	if this.level != "" {
		flag.Lookup("stderrthreshold").Value.Set(this.level)
	}

	if this.tostd != "" {
		flag.Lookup("logtostderr").Value.Set(this.tostd)
	}

	if this.file != "" {
		SetLogFile(this.file)
	}
	return true
}

func (this *GlogService) Name() string { return "glog" }

func (this *GlogService) Reload() bool { return true }

func (this *GlogService) Final() {
	Flush()
}

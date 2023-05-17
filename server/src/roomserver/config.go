package main

import (
	"base/env"
	"base/glog"
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v3"
)

type ConfigMgr struct {
	prompyMap map[string]string
}

var (
	configm      *ConfigMgr
	configmMutex sync.RWMutex
)

func NewConfigMgr() *ConfigMgr {
	c := &ConfigMgr{
		prompyMap: make(map[string]string),
	}
	return c
}

func ConfigMgr_GetMe() (c *ConfigMgr) {
	if configm == nil {
		configm = NewConfigMgr()
	}
	configmMutex.RLock()
	c = configm
	configmMutex.RUnlock()
	return
}

func ReloadConfig() bool {
	c := NewConfigMgr()
	if !c.Init() {
		return false
	}
	configmMutex.Lock()
	configm = c
	configmMutex.Unlock()
	return true
}

func (this *ConfigMgr) Init() bool {
	if !this.LoadPrompyConfig() {
		return false
	}

	glog.Info("[配置] 加载配置成功 ")

	return true
}

func (this *ConfigMgr) LoadPrompyConfig() bool {

	file, err := ioutil.ReadFile(env.Get("global", "config") + "prompy.yaml")
	if err != nil {
		glog.Info("[配置] 找不到配置文件prompy.yaml")
		return false
	}

	err = yaml.Unmarshal(file, &this.prompyMap)
	if err != nil {
		glog.Info("[配置] 解析提示语配置失败 ", err)
		return false
	}

	return true
}

func (this *ConfigMgr) GetPrompy(key string) string {
	prompy, ok := this.prompyMap[key]
	if !ok {
		return ""
	}

	return prompy
}

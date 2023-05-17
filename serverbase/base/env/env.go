package env

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var _configData map[string]map[string]string

var _evnData map[string]string

func Load(path string) bool {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln("[配置] 读取失败 ", path, ",", err)
		return false
	}
	err = json.Unmarshal(file, &_configData)
	if err != nil {
		log.Fatalln("[配置] 解析失败 ", path, ",", err)
		return false
	}
	return true
}

func Get(table, key string) string {
	t, ok := _configData[table]
	if !ok {
		return ""
	}
	val, ok := t[key]
	if !ok {
		return ""
	}
	return val
}

func GetEnv(key string) string {
	return _evnData[key]
}

func SetEnv(key, val string) {
	_evnData[key] = val
}

func Global(key string) string {
	val, ok := _configData["global"][key]
	if !ok {
		return ""
	}
	return val
}

func Set(key1, key2, val string) {
	if kmap, ok := _configData[key1]; ok {
		kmap[key2] = val
	}
}

func User(key string) string {
	val, ok := _configData["user"][key]
	if !ok {
		return ""
	}
	return val
}

func Room(key string) string {
	val, ok := _configData["room"][key]
	if !ok {
		return ""
	}
	return val
}

func Video(key string) string {
	val, ok := _configData["video"][key]
	if !ok {
		return ""
	}
	return val
}

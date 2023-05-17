package main

import (
	"sync"
	"time"
)

type MatchUser struct {
	token string
	name  string
	tm    int64
}

type MatchMgr struct {
	mutex    sync.Mutex
	userlist map[string]*MatchUser
}

var g_MatchMgr *MatchMgr

func MatchMgr_GetMe() *MatchMgr {

	if g_MatchMgr == nil {
		g_MatchMgr = &MatchMgr{
			userlist: make(map[string]*MatchUser),
		}
	}

	return g_MatchMgr
}

func (this *MatchMgr) Add(name string) string {

	token := RandToken(32)
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.userlist[token] = &MatchUser{
		token: token,
		name:  name,
	}

	return token
}

func (this *MatchMgr) GetMatchRoom(token string) []*MatchUser {

	roomPlayerCount := 1

	this.mutex.Lock()
	defer this.mutex.Unlock()

	player, ok := this.userlist[token]
	if !ok {
		return nil
	}

	now := time.Now().Unix()
	player.tm = now

	for k, v := range this.userlist {
		if now-v.tm < 3 {
			continue
		}
		// 删除过期玩家
		delete(this.userlist, k)
	}

	if len(this.userlist) < roomPlayerCount {
		return nil
	}

	var m []*MatchUser
	m = append(m, player)
	delete(this.userlist, token)

	for _, user := range this.userlist {

		m = append(m, user)
		if len(m) >= roomPlayerCount {
			break
		}
	}

	return m
}

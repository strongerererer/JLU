package main

import (
	"base/glog"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"
)

const (
	MAX_KEEPEND_TIME = 600
)

type RoomMgr struct {
	uniqueId uint64

	runmutex sync.RWMutex
	runRooms map[uint64]*Room

	endmutex sync.RWMutex
	endrooms map[uint64]int64

	tokenMutex sync.RWMutex
	tokenrooms map[string]uint64
}

var roommgr *RoomMgr

func RoomMgr_GetMe() *RoomMgr {
	if roommgr == nil {
		roommgr = &RoomMgr{
			runRooms:   make(map[uint64]*Room),
			endrooms:   make(map[uint64]int64),
			tokenrooms: make(map[string]uint64),
		}
		roommgr.Init()
	}
	return roommgr
}

func (this *RoomMgr) Init() {
	go func() {
		minTick := time.NewTicker(time.Second * 1)
		var tickId uint32
		defer func() {
			if err := recover(); err != nil {
				glog.Error("[异常] 报错 ", err, "\n", string(debug.Stack()))
			}
			minTick.Stop()
		}()
		for {
			select {
			case <-minTick.C:
				newTickId := atomic.AddUint32(&tickId, 1)
				if newTickId%60 == 0 {
					this.ChkEndRoomId()
				}
			}
		}
	}()
}

func (this *RoomMgr) AllocId() uint64 {
	this.uniqueId++
	return this.uniqueId
}

// 添加房间
func (this *RoomMgr) AddRoom(room *Room) (*Room, bool) {
	this.runmutex.Lock()
	defer this.runmutex.Unlock()
	oldroom, ok := this.runRooms[room.id]
	if ok {
		return oldroom, false
	}
	this.runRooms[room.id] = room

	return room, true
}

// 删除房间
func (this *RoomMgr) RemoveRoom(room *Room) {
	this.runmutex.Lock()
	delete(this.runRooms, room.id)
	this.runmutex.Unlock()
	this.AddEndRoomId(room.id)
	glog.Info("[房间] 删除房间 ", room.id)
}

// 添加结束房间
func (this *RoomMgr) AddEndRoomId(roomid uint64) {
	this.endmutex.Lock()
	this.endrooms[roomid] = time.Now().Unix() + MAX_KEEPEND_TIME
	this.endmutex.Unlock()
}

// 是否结束房间
func (this *RoomMgr) IsEndRoomId(roomid uint64) bool {
	this.endmutex.Lock()
	defer this.endmutex.Unlock()
	endtime, ok := this.endrooms[roomid]
	if !ok {
		return false
	}
	if endtime < time.Now().Unix() {
		delete(this.endrooms, roomid)
		return false
	}
	return true
}

// 检查结束房间列表
func (this *RoomMgr) ChkEndRoomId() {
	timenow := time.Now().Unix()
	this.endmutex.Lock()
	for roomid, endtime := range this.endrooms {
		if endtime > timenow {
			continue
		}
		delete(this.endrooms, roomid)
	}
	this.endmutex.Unlock()
}

// 获取房间列表
func (this *RoomMgr) GetRooms() (rooms []*Room) {
	this.runmutex.RLock()
	for _, room := range this.runRooms {
		rooms = append(rooms, room)
	}
	this.runmutex.RUnlock()
	return
}

// 根据id获取有效房间
func (this *RoomMgr) getRoomById(rid uint64) *Room {
	this.runmutex.RLock()
	model, ok := this.runRooms[rid]
	this.runmutex.RUnlock()
	if !ok {
		return nil
	}
	return model
}

func (this *RoomMgr) AddTokenRoom(roomid uint64, tokens []string) {

	this.tokenMutex.Lock()
	defer this.tokenMutex.Unlock()

	for _, token := range tokens {
		this.tokenrooms[token] = roomid
	}
}

func (this *RoomMgr) GetRoomByToken(token string) *Room {

	this.tokenMutex.RLock()
	roomid, ok := this.tokenrooms[token]
	this.tokenMutex.RUnlock()

	if !ok {
		return nil
	}

	return this.getRoomById(uint64(roomid))
}

func (this *RoomMgr) GetWaitRoom(name string) *Room {

	this.runmutex.RLock()
	defer this.runmutex.RUnlock()

	for _, room := range this.runRooms {

		if room.IsFulled() {
			continue
		}

		return room
	}

	return nil
}

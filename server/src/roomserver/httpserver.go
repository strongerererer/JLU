package main

import (
	"base/env"
	"base/glog"
	"encoding/json"
	"net"
	"net/http"
	"strconv"
	"time"
)

func StartHttpServer() bool {
	http.Handle("/", http.FileServer(http.Dir(env.Get("room", "static"))))
	http.HandleFunc("/time", HandleTime)
	http.HandleFunc("/room_create", HandleRoomCreate)
	http.HandleFunc("/room_chat", HandleRoomChat)
	http.HandleFunc("/room_chatlist", HandleRoomChatList)
	http.HandleFunc("/room_join", HandleRoomJoin)

	http.HandleFunc("/match_add", HandleMatchAdd)
	http.HandleFunc("/match_status", HandleMatchStatus)

	listen := env.Get("room", "listen")
	ln, err := net.Listen("tcp", listen)
	if err != nil {
		glog.Error("[启动] http绑定端口失败 ", listen)
		return false
	}
	go http.Serve(ln, nil)

	glog.Info("[启动] http绑定端口成功 ", listen)
	return true
}

func RetWithData(res http.ResponseWriter, data interface{}) {

	dataMap := make(map[string]interface{})
	dataMap["errcode"] = 0
	dataMap["data"] = data
	buf, err := json.Marshal(dataMap)
	if err != nil {
		RetWithErrCode(res, 1)
		return
	}

	res.Write(buf)
}

func RetWithErrCode(res http.ResponseWriter, errcode uint32) {

	dataMap := make(map[string]interface{})
	dataMap["errcode"] = errcode

	buf, _ := json.Marshal(dataMap)

	res.Write(buf)
}

// 获取服务器时间 /time
func HandleTime(res http.ResponseWriter, req *http.Request) {
	timenow := time.Now().Unix()
	res.Write([]byte(strconv.FormatInt(timenow, 10)))
}

func HandleRoomCreate(res http.ResponseWriter, req *http.Request) {

	room := NewRoom()
	RoomMgr_GetMe().AddRoom(room)

	RetWithData(res, room.GetInfo())
}

func HandleRoomChat(res http.ResponseWriter, req *http.Request) {
	roomid := StrToUint64(req.FormValue("roomid"))
	token := req.FormValue("token")
	chat := req.FormValue("chat")

	room := RoomMgr_GetMe().getRoomById(roomid)
	if room == nil {
		RetWithErrCode(res, 1)
		return
	}

	room.chan_Chat <- &PlayerChat{Token: token, Text: chat}
	RetWithErrCode(res, 0)
}

func HandleRoomChatList(res http.ResponseWriter, req *http.Request) {
	roomid := StrToUint64(req.FormValue("roomid"))
	chatid := StrToUint32(req.FormValue("chatid"))
	token := req.FormValue("token")

	room := RoomMgr_GetMe().getRoomById(roomid)
	if room == nil {
		RetWithErrCode(res, 1)
		return
	}

	RetWithData(res, room.GetMsgList(token, chatid))
}

func HandleMatchAdd(res http.ResponseWriter, req *http.Request) {

	name := req.FormValue("name")
	token := MatchMgr_GetMe().Add(name)

	if len(name) == 0 {
		RetWithErrCode(res, 1)
		return
	}

	dataMap := make(map[string]interface{})
	dataMap["token"] = token
	RetWithData(res, dataMap)
}

func HandleMatchStatus(res http.ResponseWriter, req *http.Request) {

	token := req.FormValue("token")

	// 查看玩家是否已在房间
	room := RoomMgr_GetMe().GetRoomByToken(token)
	if room != nil {
		RetWithData(res, room.GetInfo())
		return
	}

	// 匹配池玩家不足
	userlist := MatchMgr_GetMe().GetMatchRoom(token)
	if len(userlist) == 0 {
		RetWithErrCode(res, 1)
		return
	}

	// 创建房间
	room = NewRoom()

	var tokens []string
	for _, user := range userlist {
		room.chan_PlayerAdd <- &PlayerAdd{
			name:  user.name,
			token: user.token,
		}
		tokens = append(tokens, user.token)
	}

	RoomMgr_GetMe().AddTokenRoom(room.id, tokens)
	RoomMgr_GetMe().AddRoom(room)

	RetWithData(res, room.GetInfo())
}

func HandleRoomJoin(res http.ResponseWriter, req *http.Request) {

	name := req.FormValue("name")
	head := req.FormValue("head")

	var (
		room *Room
	)

	room = RoomMgr_GetMe().GetWaitRoom(name)

	if room == nil {
		room = NewRoom()
		glog.Info("[房间] 创建新的房间 ", name, ",", room.id)
	}

	token := RandToken(32)
	room.chan_PlayerAdd <- &PlayerAdd{
		name:  name,
		token: token,
		head:  head,
	}

	RoomMgr_GetMe().AddRoom(room)

	info := room.GetInfo()
	info["token"] = token
	RetWithData(res, info)
}

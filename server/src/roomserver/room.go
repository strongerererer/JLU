package main

import (
	"base/glog"
	"fmt"
	"sync/atomic"
	"time"
)

type PlayerChat struct {
	Token string
	Text  string
}

type ChatInfo struct {
	Id   uint32
	Name string
	Text string
	Tm   int64
}

type PlayerAdd struct {
	name  string
	head  string
	token string
}

type Room struct {
	id           uint64
	playerNum    int32
	MaxPlayerNum int32
	players      map[uint32]*Player
	tokens       map[string]uint32

	timeLoop int64

	chatList       []ChatInfo
	waitActionList []uint32

	godAgent       *GodAgent
	assistantAgent *AssistantAgent
	msgFilter      *MsgFilter

	stage    string
	isclosed int

	chan_Control   chan int //可以用于停止程序，select控制了
	chan_Chat      chan *PlayerChat
	chan_PlayerAdd chan *PlayerAdd

	allPlayerChat string //存一下这组所有玩家的发言信息
}

func NewRoom() *Room {
	room := &Room{
		id:      RoomMgr_GetMe().AllocId(),
		players: make(map[uint32]*Player),
		tokens:  make(map[string]uint32),
		//godAgent:       NewGodAgent(),
		assistantAgent: NewAssistantAgent(),
		msgFilter:      NewMsgFilter(),
		chan_Control:   make(chan int, 1),
		chan_Chat:      make(chan *PlayerChat, 10),
		chan_PlayerAdd: make(chan *PlayerAdd, 15),
		MaxPlayerNum:   1,
	}

	if !room.init() {
		return nil
	}

	room.Start()

	return room
}

func (this *Room) init() bool {

	/*
		this.AppendMsg(nil, "加入房间成功,请等待其他玩家加入")

		names := this.assistantAgent.RandName()
		occs := this.assistantAgent.RandOcc(names)

		for i, name := range names {

			id := uint32(i) + 1
			p := NewPlayer(id, name, occs[i])

			this.players[id] = p
		}

		this.AppendMsg(nil, strings.Join(names, ",")+"加入房间")

		this.AppendMsg(nil, this.GetPlayerOccList())
	*/

	//this.AppendMsg(nil, fmt.Sprint("创建房间成功,本房间需要玩家数:", this.MaxPlayerNum))
	this.msgFilter.GodMsg(fmt.Sprint("创建房间服务器成功"))

	return true
}

func (this *Room) Start() {

	go this.Loop()

	go func() {
		/*
			// 通知上帝游戏开始
			stage := this.godAgent.GameStart()
			this.SetStage(stage)

			for {
				if this.isclosed != 0 {
					return
				}

				// 阶段继续
				stage := this.godAgent.Next()
				this.SetStage(stage)
			}
		*/
	}()
}

func (this *Room) NotifyWatiActionPlayer() {
	//id := this.waitActionList[0]
	//player := this.GetPlayer(id)
	//this.msgFilter.PrivateMsg(player, "请提问")
}

func (this *Room) SetStage(s string) {
	//在这里运行两个筛选GPT获得执行人的编号

	this.stage = s
	//this.AppendMsg(nil, s)
	this.msgFilter.GodMsg(s) //上帝前端输出

	// 根据当前阶段，询问助手，那些人需要作出行为
	// 行为全部对应返回后，让上帝继续，直到游戏结束
	// 检查阶段需要谁做行为

	this.waitActionList = this.assistantAgent.CheckStageUser() //只有一个1
	this.NotifyWatiActionPlayer()

}

func (this *Room) GameEnd() {

	//this.AppendMsg(nil, "游戏结束")
	this.msgFilter.GodMsg("游戏结束")
	this.isclosed = 1
}

func (this *Room) Stop() {

	this.isclosed = 1
	this.chan_Control <- 1
}

func (this *Room) GetInfo() map[string]interface{} {

	data := make(map[string]interface{})
	data["id"] = this.id

	return data
}

func (this *Room) Loop() {

	var (
		tick = time.NewTicker(40 * time.Millisecond)
	)

	for {
		select {
		case <-tick.C:
			this.timeLoop++
			if this.timeLoop%25 == 0 {
				/*
					str := this.GodAgent.Next()
					this.AppendMsg(nil, str)

					if this.assistantAgent.IsEnd(str) {
						this.AppendMsg(nil, "游戏结束")
						return
					}
				*/
			}

		case chat := <-this.chan_Chat:
			this.onChat(chat.Token, chat.Text)
		case player := <-this.chan_PlayerAdd:
			this.onPlayerAdd(player.token, player.name, player.head)
		case <-this.chan_Control:
			return
		}
	}
}

func (this *Room) doChatGPTTest() {

	if len(this.chatList) >= 50 {
		return
	}

	glog.Info("doChatGPTTest")

	for _, p := range this.players {

		text := p.doChatGPT("随机帮我说一句话")
		if len(text) != 0 {
			this.chatList = append(this.chatList, ChatInfo{
				Id:   p.id,
				Name: p.name,
				Text: text,
				Tm:   time.Now().Unix(),
			})
		}

		glog.Info("[chat] 玩家", p.name, "说: ", text)
		return
	}
}

func (this *Room) onChat(token string, text string) {

	if len(this.waitActionList) == 0 {

		//this.AppendMsg(nil, "执行人队列为空")
		return
	}

	id, ok := this.tokens[token]
	if !ok {
		return
	}
	p := this.GetPlayer(id)
	if p == nil {
		return
	}

	if id != this.waitActionList[0] {
		//this.AppendMsg(nil, "你不是当前阶段执行人")
		this.msgFilter.PrivateMsg(p, "你不是当前阶段执行人")
		return
	}

	this.waitActionList = this.waitActionList[1:]

	/*
		glog.Info("[chat] playerszie ", len(this.players))

		for _, p := range this.players {

			glog.Info("[chat] chatgpt ", p.id, ",", p.doChatGPT(text))
		}
	*/

	//这里要改一下，把执行一次GPT该成把这次玩家输出保存在一个string里面保存着

	this.msgFilter.PlayerMsg(p, text) //玩家说话让大家都看到

	curString := fmt.Sprint(text)
	//this.allChat += curString + "\n"
	this.allPlayerChat += curString + "\n"

	//result := this.godAgent.PlayerSay(p, text)
	//this.msgFilter.PrivateMsg(p, result)

	//通知让队列最前面的那个哥们继续操作
	if len(this.waitActionList) != 0 {
		this.NotifyWatiActionPlayer()
	}
}

func (this *Room) IncPlayerNum() int32 {

	return atomic.AddInt32(&this.playerNum, 1)
}

func (this *Room) IsFulled() bool {
	return atomic.LoadInt32(&this.playerNum) >= this.MaxPlayerNum
}

func (this *Room) onPlayerAdd(token string, name string, head string) {

	this.IncPlayerNum()

	id := uint32(len(this.players)) + 1
	this.players[id] = NewPlayer(id, name, "", head)
	this.tokens[token] = id

	//this.AppendMsg(nil, fmt.Sprint(name, "加入房间成功,剩余玩家:", this.MaxPlayerNum-num))
	this.msgFilter.GodMsg(fmt.Sprint(name, "加入系统成功，可开始提问"))

	if this.IsFulled() {
		this.GameStart()
	}

	glog.Info("[房间] 玩家加入 ", token, ",", name, ",", id)
}

func (this *Room) GameStart() {

	//this.msgFilter.GodMsg("游戏开始")

	go func() {

		ids := this.GetIdList()
		var occs []string
		occs = append(occs, "游客")
		for i, id := range ids {
			p := this.GetPlayer(StrToUint32(id))
			if p == nil {
				continue
			}
			if i >= len(occs) {
				continue
			}
			p.SetOcc(occs[i])

			//this.msgFilter.PrivateMsg(p, fmt.Sprint("你的身份:", occs[i]))
		}

		//this.AppendMsg(nil, this.GetPlayerOccList())

		// 通知上帝游戏开始
		this.godAgent = NewGodAgent()
		stage := this.godAgent.GameStart()
		//根据阶段，把对应执行人加入到waitActionList
		this.SetStage(stage)

		//这里是所有需要发言的玩家发言完了以后（waitActionList==0），给上帝一个继续的命令
		for {
			if len(this.waitActionList) != 0 {
				time.Sleep(100 * time.Millisecond)
				continue
			}
			// 阶段继续，把这次所有人的信息输给上帝，然后清空存储区
			stage := this.godAgent.Say(this.allPlayerChat)
			this.allPlayerChat = ""
			this.SetStage(stage)
		}
	}()
}

func (this *Room) getChatList(chatid uint32) interface{} {

	maxId := len(this.chatList)
	dataMap := make(map[string]interface{})
	dataMap["chatId"] = maxId
	if chatid < uint32(maxId) {
		dataMap["chatList"] = this.chatList[chatid:]
	}

	//glog.Info("getChatList ", chatid, ",", maxId, ",", this.chatList)
	return dataMap
}

func (this *Room) GetMsgList(token string, chatid uint32) interface{} {

	glog.Info("[房间] MsgList ", token, ",", chatid)

	id, ok := this.tokens[token]
	if !ok {
		return nil
	}

	p := this.GetPlayer(id)
	if p == nil {
		return nil
	}

	return this.msgFilter.GetMsgData(p, chatid)
}

/*
func (this *Room) AppendMsg(p *Player, text string) {

	if len(text) == 0 {
		return
	}

	var (
		id   uint32 = 0
		name string = "系统"
	)

	if p != nil {
		id = p.id
		name = p.name
	}

	this.chatList = append(this.chatList, ChatInfo{
		Id:   id,
		Name: name,
		Text: text,
		Tm:   time.Now().Unix(),
	})
}
*/

func (this *Room) GetPlayerOccList() string {

	str := ""

	for i, p := range this.players {
		if i != 0 {
			str += " "
		}
		str += p.name + ":" + p.occ

	}

	return str
}

// GetIdOccList 返回这局游戏的所有玩家信息
func (this *Room) GetIdOccList() string {

	str := ""

	for i, p := range this.players {
		if i != 0 {
			str += " "
		}
		str += fmt.Sprint(p.id, ":", p.occ)

	}

	return str
}

func (this *Room) GetIdList() []string {

	var ids []string

	for _, p := range this.players {
		ids = append(ids, fmt.Sprint(p.id))
	}

	return ids
}

func (this *Room) GetPlayer(id uint32) *Player {
	p, ok := this.players[id]
	if !ok {
		return nil
	}

	return p
}

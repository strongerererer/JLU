package main

import "strings"

type Player struct {
	id     uint32 //阿拉伯id
	name   string //名字
	occ    string //身份
	prompt string
	head   string
}

func NewPlayer(id uint32, name string, occ string, head string) *Player {

	player := &Player{
		id:   id,   //玩家编号
		name: name, //名字
		occ:  occ,  //身份
		head: head,
	}

	player.init()

	return player
}

func (this *Player) init() {

	prompt := ConfigMgr_GetMe().GetPrompy("player")
	prompt = strings.Replace(prompt, "{{name}}", this.name, -1)
	prompt = strings.Replace(prompt, "{{occ}}", this.occ, -1)
	this.prompt = prompt
}

func (this *Player) SetOcc(occ string) {
	this.occ = occ
}

func (this *Player) GetOcc() string {
	return this.occ
}

func (this *Player) doChatGPT(task string) string {

	return LLM_GetMe().ExecChatGPT(this.prompt, task, "")
}

func (this *Player) ExecAction(stage string) string {

	prompt := strings.Replace(this.prompt, "{{stage}}", stage, -1)

	return LLM_GetMe().ExecChatGPT(prompt, "请说出你的想法", "")
}

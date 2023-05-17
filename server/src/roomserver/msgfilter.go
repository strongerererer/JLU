package main

const (
	MsgFilterType_None   = 0
	MsgFilterType_Player = 1
)

type MsgItem struct {
	Name        string
	Content     string
	MsgType     uint32
	PrivateId   uint32 // 私有消息,仅玩家ID可见
	PrivateName string
	Head        string
}

type MsgFilter struct {
	msglist []*MsgItem
}

func NewMsgFilter() *MsgFilter {
	filter := &MsgFilter{}

	return filter
}

func (this *MsgFilter) GodMsg(msg string) {

	this.msglist = append(this.msglist, &MsgItem{
		Name:    "上帝",
		Content: msg,
		MsgType: MsgFilterType_None,
	})
}

func (this *MsgFilter) PlayerMsg(player *Player, msg string) {

	this.msglist = append(this.msglist, &MsgItem{
		Name:    player.name,
		Content: msg,
		MsgType: MsgFilterType_Player,
		Head:    player.head,
	})
}

func (this *MsgFilter) PrivateMsg(player *Player, msg string) {

	this.msglist = append(this.msglist, &MsgItem{
		Name:        "上帝",
		Content:     msg,
		MsgType:     MsgFilterType_None,
		PrivateId:   player.id,
		PrivateName: player.name,
	})
}

func (this *MsgFilter) GetMsgData(player *Player, chatId uint32) map[string]interface{} {

	maxId := len(this.msglist)
	dataMap := make(map[string]interface{})
	dataMap["chatId"] = maxId

	if chatId > uint32(maxId) {
		return nil
	}

	var (
		tmplist []*MsgItem = this.msglist[chatId:]
		retlist []*MsgItem
	)

	for i, _ := range tmplist {
		item := tmplist[i]

		if item.PrivateId != 0 && item.PrivateId != player.id {
			continue
		}

		retlist = append(retlist, item)
	}

	dataMap["chatList"] = retlist

	return dataMap

}

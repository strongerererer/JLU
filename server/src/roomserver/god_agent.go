package main

import (
	"fmt"
	"github.com/sashabaranov/go-openai"
)

type GodAgent struct {
	/*	// 提示语
		prompt string
		// 上下文
		context []string
		优势：整体局势处理的相对比较明白
		劣势：一旦有不符合要求的对话，上帝的回应容易直接崩掉，对玩家回复的要求很严格
	*/

	//messages输入,可使用对话形式的GPT函数
	//优势：可以模拟出真实的上帝发言，流程过的更加顺畅。
	//劣势：长对话分析能力相对上面那种方式会弱一点，因为保存了上帝自己的发言，后续可专门用一个GPT用于分析局势，只分析玩家行为，输出谁还活着来解决此缺陷
	godMessages []openai.ChatCompletionMessage
}

func NewGodAgent() *GodAgent {
	//把system的内容放进去
	messages := make([]openai.ChatCompletionMessage, 0)

	sys := ConfigMgr_GetMe().GetPrompy("god")
	//需要在这里替换一下{{occ}}
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: sys,
	})

	agent := &GodAgent{
		godMessages: messages,
	}

	return agent
}

func (this *GodAgent) Next() string {
	//这个地方在上帝这里输入继续，可以用在两个地方：玩家回复超时或者是该角色玩家已死亡。
	//返回上帝说的话
	user := "无操作"
	this.godMessages = append(this.godMessages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: user,
	})
	text := LLM_GetMe().ExecMessagesChatGPT(this.godMessages)

	this.godMessages = append(this.godMessages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: text,
	})

	return text
}

func (this *GodAgent) GameStart() string {

	//告诉上帝游戏开始了，预期返回值是：天黑请闭眼，狼人请睁眼，选择你要杀的对象

	task := "开始"
	this.godMessages = append(this.godMessages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: task,
	})

	text := LLM_GetMe().ExecMessagesChatGPT(this.godMessages)

	this.godMessages = append(this.godMessages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: text,
	})

	return text
}

func (this *GodAgent) PlayerSay(p *Player, saytext string) string {

	task := fmt.Sprint(p.id, "号玩家:", saytext)
	/*result := LLM_GetMe().ExecChatGPT(this.prompt, task, strings.Join(this.context, "\n"))
	this.context = append(this.context, task)*/
	this.godMessages = append(this.godMessages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: task,
	})

	text := LLM_GetMe().ExecMessagesChatGPT(this.godMessages)

	this.godMessages = append(this.godMessages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: text,
	})
	return text
}
func (this *GodAgent) Say(saytext string) string {

	task := fmt.Sprint(saytext)
	/*result := LLM_GetMe().ExecChatGPT(this.prompt, task, strings.Join(this.context, "\n"))
	this.context = append(this.context, task)*/
	this.godMessages = append(this.godMessages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: task,
	})

	text := LLM_GetMe().ExecMessagesChatGPT(this.godMessages)

	this.godMessages = append(this.godMessages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: text,
	})
	return text
}

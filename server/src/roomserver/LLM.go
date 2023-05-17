package main

import (
	"base/env"
	"base/glog"
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"net/http"
	"net/url"
)

const (
	ChatGPTKey = "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"
)
const (
	LLM_ChatGPT = 1
)

type LLM struct {
	ChatGptCount uint32
}

var g_LLM *LLM

func LLM_GetMe() *LLM {
	if g_LLM == nil {
		g_LLM = &LLM{}
	}
	return g_LLM
}

func (this *LLM) ExecChatGPT(system string, user string, assistant string) string {

	this.ChatGptCount++

	if this.ChatGptCount >= 100 {
		glog.Info("[chatgpt] 调用次数过多 ", system, ",", user, ",", assistant)
		return ""
	}

	proxyUrl := env.Get("global", "proxy")
	config := openai.DefaultConfig(ChatGPTKey)
	if proxyUrl != "" {
		config.HTTPClient.Transport = &http.Transport{
			// 设置代理
			Proxy: func(req *http.Request) (*url.URL, error) {
				return url.Parse(proxyUrl)
			}}
	}

	var messages []openai.ChatCompletionMessage

	if len(system) != 0 {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleSystem,
			Content: system,
		})
	}

	if len(user) != 0 {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: user,
		})
	}

	if len(assistant) != 0 {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: assistant,
		})
	}

	client := openai.NewClientWithConfig(config)
	ctx := context.Background()
	req := openai.ChatCompletionRequest{ //设置各个要求
		Model:       openai.GPT3Dot5Turbo,
		Temperature: 0.001,
		Messages:    messages,
	}
	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		return fmt.Sprint(err)
	}

	content := resp.Choices[0].Message.Content

	glog.Info("[chatgpt] ExecChatGPT: \nsystem:\n", system, "\nuser:\n", user, "\nassistant:\n", assistant, "\nresult:\n", content)

	return content
}

// ExecMessagesChatGPT 输入参数直接是messages，达到<chat>的效果
func (this *LLM) ExecMessagesChatGPT(messages []openai.ChatCompletionMessage) string {
	this.ChatGptCount++

	if this.ChatGptCount >= 100 {
		glog.Info("[chatgpt] 调用次数过多 ")
		return ""
	}
	proxyUrl := env.Get("global", "proxy")
	config := openai.DefaultConfig(ChatGPTKey)
	if proxyUrl != "" {
		config.HTTPClient.Transport = &http.Transport{
			// 设置代理
			Proxy: func(req *http.Request) (*url.URL, error) {
				return url.Parse(proxyUrl)
			}}
	}

	client := openai.NewClientWithConfig(config)
	ctx := context.Background()
	req := openai.ChatCompletionRequest{ //设置各个要求
		Model:       openai.GPT3Dot5Turbo,
		Temperature: 0.001,
		Messages:    messages,
	}
	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		return fmt.Sprint(err)
	}
	content := resp.Choices[0].Message.Content

	glog.Info("[chatgpt] ExecMessagesChatGPT: \nsystem:\n", messages, "result:\n", content)

	return content

}

package main

type AssistantAgent struct {
}

func NewAssistantAgent() *AssistantAgent {
	agent := &AssistantAgent{}

	return agent
}

// CheckStageUser 筛选GPT
func (this *AssistantAgent) CheckStageUser() []uint32 {

	var vec []uint32
	vec = append(vec, 1)
	return vec

}

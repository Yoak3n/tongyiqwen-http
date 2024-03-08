package qianwen

import (
	"github.com/tidwall/gjson"
	"log"
	"tongyiqwen/package/model"
	"tongyiqwen/plugin"
)

type Conversations struct {
	Sub map[string][]model.Message
}

const DefaultPreset = "你是ChatGPT, 一个由OpenAI训练的大型语言模型, 你旨在回答并解决人们的任何问题，并且可以使用多种语言与人交流。"

var convs *Conversations

func init() {
	convs = new(Conversations)
	convs.Sub = make(map[string][]model.Message)
}

func NewConversation(id string, preset string, question string) string {

	newMsg := make([]model.Message, 0)
	if preset != "" {
		p, err := plugin.LoadTextPreset(preset)
		if err != nil {
			m, e := plugin.LoadMapPreset(preset)
			if e != nil {
				newMsg = append(newMsg, model.Message{Role: "system", Content: DefaultPreset})
			} else {
				newMsg = m
			}
		} else {
			newMsg = append(newMsg, model.Message{Role: "system", Content: p})
		}
	} else {
		newMsg = append(newMsg, model.Message{Role: "system", Content: DefaultPreset})
	}

	newMsg = append(newMsg, model.Message{Role: "user", Content: question})
	convs.Sub[id] = newMsg
	answer := makeQuestionBody(newMsg)
	log.Println(string(answer))
	jResult := gjson.Parse(string(answer))
	reply := jResult.Get("Data.Choices").Array()
	if len(reply) > 0 {
		r := reply[0].Get("Message.Content").String()
		convs.Sub[id] = append(convs.Sub[id], model.Message{
			Role:    "assistant",
			Content: r,
		})
		return r
	} else {
		return jResult.Get("Message").String()
	}
}

func ContinueConversation(id string, question string) string {
	newMsg := []model.Message{
		{Role: "user", Content: question},
	}
	convs.Sub[id] = append(convs.Sub[id], newMsg...)
	answer := makeQuestionBody(convs.Sub[id])
	jResult := gjson.Parse(string(answer))
	reply := jResult.Get("Data.Choices").Array()[0].Get("Message.Content").String()
	convs.Sub[id] = append(convs.Sub[id], model.Message{
		Role:    "assistant",
		Content: reply,
	})
	return reply
}

func Ask(id string, preset string, question string) string {
	if _, ok := convs.Sub[id]; !ok {
		return NewConversation(id, preset, question)
	} else if question == "/reset" || question == "/重置" {
		delete(convs.Sub, id)
		return "Conversation reset!"
	} else {
		return ContinueConversation(id, question)
	}
}

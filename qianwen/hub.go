package qianwen

import (
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"time"
	"tongyiqwen/package/model"
	"tongyiqwen/plugin"
)

type Conversations struct {
	Sub    map[string][]model.Message
	Update map[string]int64
}

const DefaultPreset = "你是ChatGPT, 一个由OpenAI训练的大型语言模型, 你旨在回答并解决人们的任何问题，并且可以使用多种语言与人交流。"

var convs *Conversations

func init() {
	convs = new(Conversations)
	convs.Sub = make(map[string][]model.Message)
	convs.Update = make(map[string]int64)
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
	convs.Sub[id] = newMsg
	convs.Update[id] = time.Now().Unix()
	if question == "" {
		return fmt.Sprintf("Load preset %s successfully", preset)
	}
	convs.Sub[id] = append(newMsg, model.Message{Role: "user", Content: question})
	count := 0
	for {
		if count > 3 {
			return "API令牌认证失败"
		}
		count++
		answer := makeQuestionBody(convs.Sub[id])
		reply, err := checkAndRefresh(answer, id)
		if err != nil {
			continue
		}
		return reply
	}
}

func ContinueConversation(id string, question string) string {
	newMsg := []model.Message{
		{Role: "user", Content: question},
	}
	convs.Sub[id] = append(convs.Sub[id], newMsg...)
	fmt.Println(convs.Sub[id])
	convs.Update[id] = time.Now().Unix()
	count := 0
	for {
		if count > 3 {
			return "API令牌认证失败"
		}
		count++
		answer := makeQuestionBody(convs.Sub[id])
		reply, err := checkAndRefresh(answer, id)
		if err != nil {
			continue
		}
		return reply
	}

}

func checkAndRefresh(answer []byte, id string) (string, error) {
	jResult := gjson.ParseBytes(answer)
	reply := jResult.Get("Data.Choices").Array()
	if len(reply) > 0 {
		r := reply[0].Get("Message.Content").String()
		convs.Sub[id] = append(convs.Sub[id], model.Message{
			Role:    "assistant",
			Content: r,
		})
		convs.Update[id] = time.Now().Unix()
		return r, nil
	} else {
		CreateToken()
		time.Sleep(time.Second)
		return "", errors.New(jResult.Get("Message").String())
	}
}

func Ask(id string, preset string, question string) string {
	if _, ok := convs.Sub[id]; !ok {
		return NewConversation(id, preset, question)
	} else if time.Since(time.Unix(convs.Update[id], 0)).Hours() > 24 {
		delete(convs.Sub, id)
		return NewConversation(id, preset, question)
	} else if preset != "" {
		return NewConversation(id, preset, question)
	} else if question == "/reset" || question == "/重置" {
		delete(convs.Sub, id)
		return "Conversation reset!"
	} else {
		return ContinueConversation(id, question)
	}
}

// AskWithOpenAIStyle TODO
func AskWithOpenAIStyle() {

}

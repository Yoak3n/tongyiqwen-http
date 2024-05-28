package qianwen

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"tongyiqwen/config"
	"tongyiqwen/package/ali_model"
	"tongyiqwen/package/openai_model"
	"tongyiqwen/plugin"

	"github.com/tidwall/gjson"
)

type Conversations struct {
	Sub    map[string][]openai_model.Message
	Update map[string]int64
}

const DefaultPreset = "你是ChatGPT, 一个由OpenAI训练的大型语言模型, 你旨在回答并解决人们的任何问题，并且可以使用多种语言与人交流。"

var convs *Conversations

func init() {
	convs = new(Conversations)
	convs.Sub = make(map[string][]openai_model.Message)
	convs.Update = make(map[string]int64)
}

func NewConversation(id string, preset string, question string) string {
	newMsg := make([]openai_model.Message, 0)
	if preset != "" {
		p, err := plugin.LoadTextPreset(preset)
		if err != nil {
			m, e := plugin.LoadMapPreset(preset)
			if e != nil {
				newMsg = append(newMsg, openai_model.Message{Role: "system", Content: DefaultPreset})
			} else {
				newMsg = m
			}
		} else {
			newMsg = append(newMsg, openai_model.Message{Role: "system", Content: p})
		}
	} else {
		newMsg = append(newMsg, openai_model.Message{Role: "system", Content: DefaultPreset})
	}
	convs.Sub[id] = newMsg
	convs.Update[id] = time.Now().Unix()
	if question == "" {
		return fmt.Sprintf("Load preset %s successfully", preset)
	}
	fmt.Println(newMsg)
	convs.Sub[id] = append(newMsg, openai_model.Message{Role: "user", Content: question})
	count := 0
	for {
		if count > 3 {
			return "API令牌认证失败"
		}
		count++
		answer := makeQuestionBody(convs.Sub[id])
		reply, err := checkAndRefresh(answer, id)
		if err != nil {
			fmt.Println(err)
			continue
		}
		return reply
	}
}

func ContinueConversation(id string, question string) string {
	newMsg := []openai_model.Message{
		{Role: "user", Content: question},
	}
	convs.Sub[id] = append(convs.Sub[id], newMsg...)
	fmt.Println(convs.Sub[id])
	convs.Update[id] = time.Now().Unix()
	count := 0
	for {
		count++
		answer := makeQuestionBody(convs.Sub[id])
		reply, e := checkAndRefresh(answer, id)
		if e != nil {
			if count > 3 {
				return e.Error()
			}
		} else {
			return reply
		}

	}

}

func checkAndRefresh(answer []byte, id string) (string, error) {
	jResult := gjson.ParseBytes(answer)
	reply := jResult.Get("output.choices").Array()
	if len(reply) > 0 {
		r := reply[0].Get("message.content").String()
		convs.Sub[id] = append(convs.Sub[id], openai_model.Message{
			Role:    "assistant",
			Content: r,
		})
		convs.Update[id] = time.Now().Unix()
		return r, nil
	} else {
		time.Sleep(time.Second)
		return "", errors.New(jResult.Get("output.text").String())
	}
}

func Reset(id string) {
	delete(convs.Sub, id)
	delete(convs.Update, id)
}

func Ask(id string, preset string, question string) string {
	if _, ok := convs.Sub[id]; !ok {
		return NewConversation(id, preset, question)
	} else if time.Since(time.Unix(convs.Update[id], 0)).Hours() > 24 {
		delete(convs.Sub, id)
		return NewConversation(id, preset, question)
	} else if preset != "" {
		return NewConversation(id, preset, question)
	} else {
		return ContinueConversation(id, question)
	}
}

// AskWithOpenAIStyle TODO
func AskWithOpenAIStyle(data *openai_model.RequestBody) (*openai_model.ResponseBody, error) {
	if data.Id == "" {
		data.Id = genRequestId()
	}
	client := http.Client{Timeout: time.Second * 120}

	conf := config.GetConfig()
	body := &RequestBody{
		Stream:    true,
		RequestId: data.Id,
		Parameters: &Parameters{
			ResultFormat: "text",
		},
	}
	if data.Prompt != "" {
		oip, _ := plugin.LoadTextPreset(data.Model)
		body.Input.Messages = append(body.Input.Messages, openai_model.Message{
			Role:    "system",
			Content: oip,
		}, openai_model.Message{
			Role:    "user",
			Content: data.Prompt,
		})
	} else {
		for _, v := range data.Messages {
			body.Input.Messages = append(body.Input.Messages, openai_model.Message{
				Role:    v.Role,
				Content: v.Content,
			})
		}
	}

	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(b)
	req, err := http.NewRequest("POST", QianwenApiUrl, buf)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//if data.Stream == true {
	//	req.Header.Set("Accept", "text/event-stream")
	//}else{
	//
	//}
	//req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", conf.APIkey))

	res, e := client.Do(req)
	if e != nil {
		return nil, err
	}
	defer res.Body.Close()
	content := make([]byte, 0)

	fmt.Println(string(content))
	resp := &ali_model.ResponseBody{}
	err = json.Unmarshal(content, resp)
	if len(resp.Data.Choices) <= 0 || err != nil {
		return nil, errors.New("no data")
	}
	message := &openai_model.ResponseBody{
		FinishReason: resp.Data.Choices[0].FinishReason,
		Object:       "chat.completion",
		Id:           data.Id,
		Model:        resp.Data.Usage[0].ModelId,
		Text:         resp.Data.Choices[0].Message.Content,
		Created:      time.Now().Unix(),
		Usage: &openai_model.Usage{
			PromptTokens:     resp.Data.Usage[0].InputTokens,
			CompletionTokens: resp.Data.Usage[0].OutputTokens,
			TotalTokens:      resp.Data.Usage[0].InputTokens + resp.Data.Usage[0].OutputTokens,
		},
	}
	for _, v := range resp.Data.Choices {
		choice := &openai_model.Choice{
			Index:        v.Index,
			FinishReason: v.FinishReason,
			Message: openai_model.Message{
				Content: v.Message.Content,
				Role:    v.Message.Role,
			},
		}
		message.Choices = append(message.Choices, choice)
	}
	return message, nil
}

package qianwen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"tongyiqwen/config"
	"tongyiqwen/package/openai_model"

	"github.com/google/uuid"
)

const QianwenApiUrl = "https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation"

func genRequestId() string {
	return uuid.New().String()
}

func makeQuestionBody(msg []openai_model.Message) []byte {
	conf := config.GetConfig()
	input := Input{
		Messages: msg,
	}
	client := http.Client{Timeout: time.Second * 120}
	body := &RequestBody{
		RequestId: genRequestId(),
		Input:     input,
		Model:     conf.Model,
		Parameters: &Parameters{
			ResultFormat: "message",
		},
	}
	b, err := json.Marshal(body)
	if err != nil {
		return []byte(err.Error())
	}
	buf := bytes.NewBuffer(b)
	req, err := http.NewRequest("POST", QianwenApiUrl, buf)
	if err != nil {
		fmt.Println(err)
		return []byte(err.Error())
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	//req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", conf.APIkey))
	// 为了防止token过期，这里需要重试

	res, e := client.Do(req)
	if e != nil {
		return []byte(e.Error())
	}
	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		time.Sleep(time.Second)
	}
	content, e := io.ReadAll(res.Body)
	fmt.Println(string(content))
	if e != nil {
		res.Body.Close()
		return []byte(e.Error())
	} else {
		res.Body.Close()
		return content
	}

}

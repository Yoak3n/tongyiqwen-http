package qianwen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"time"
	"tongyiqwen/config"
	"tongyiqwen/package/model"
)

const QianwenApiUrl = "https://bailian.aliyuncs.com/v2/app/completions"

func genRequestId() string {
	return uuid.New().String()
}

func makeQuestionBody(msg []model.Message) []byte {

	conf := config.GetConfig()
	client := http.Client{Timeout: time.Second * 120}
	body := &RequestBody{
		RequestId: genRequestId(),
		AppId:     conf.AppID,
		Messages:  msg,
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
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", conf.Token))
	content := make([]byte, 0)
	count := 0
	// 为了防止token过期，这里需要重试
	for {
		count++
		if count > 10 {
			return []byte("重试次数过多，退出")
		}
		res, e := client.Do(req)
		if e != nil {
			return []byte(err.Error())
		}
		if res.StatusCode != http.StatusOK {
			res.Body.Close()
			CreateToken()
			time.Sleep(time.Second * 5)
			continue
		}
		content, e = io.ReadAll(res.Body)
		if err != nil {
			res.Body.Close()
			return []byte(err.Error())
		} else {
			res.Body.Close()
			return content
		}
	}
}

<p align="center"><a href="https://hub.docker.com/r/yoaken/tongyiqwen-http"><img src="https://img.shields.io/docker/pulls/yoaken/tongyiqwen-http?&logo=docker" alt="Docker Pulls"></a>
<a href="./LICENSE"><img src="https://img.shields.io/github/license/Yoak3n/tongyiqwen-http" alt="License"></a></p>


## Preparation
type Aliyun large model authorization keys in `config.yaml`,reference `config.example.yaml`

## Usage
### Ask
Post `/v1/ask` with question to quickly ask tongyiqianwen 
```json
{
  "id": "test",
  "preset": "translator",
  "question": "test"
}
```
if `preset` is not blank,server will create a new conversation,so make it empty after beginning conversation

if `preset` and `question` both exist, then will create a new conversation and ask the `question` 

### Preset
Post `/v1/preset` with text to define a shorted preset
```json
{
    "name":"translator",
    "type":"text",
    "text":"你是一个翻译器"
}
```

Post `/v1/preset` with map to upload openai-style preset(Note the capitalized beginnings of fields)
```json
{
    "name":"translator",
    "type":"map",
    "map":[
        {"Role":"system","Content":"你是一个翻译器"},
        {"Role":"user","Content":"翻译时不能添加任何与提供的文本无关的内容"},
        {"Role":"assistant","Content":"好的，我现在就开始工作，请给出指令"}
    ]
}
```

Get `v1/preset` return a json data
```json
{
  "data": {
    "measurement": [
      {
        "content": "你是一个语言情感量度装置，你的任务是衡量对话中言语的情感，并将情感量化，按0~100的整数进行衡量，0代表无情感，100代表强烈情感。你的回答中只包含一个整数，代表情感度量，不包含其他文字。",
        "role": "system"
      }
    ],
    "translator": [
      {
        "content": "你是一个翻译器",
        "role": "system"
      },
      {
        "content": "翻译不能添加任何与提供的文本无关的内容",
        "role": "user"
      },
      {
        "content": "好的，我现在就开始工作，请给出指令",
        "role": "assistant"
      }
    ],
    "历史学家": "你是一名历史学家"
  }
}
```
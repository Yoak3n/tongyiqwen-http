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
    "content":"你是一个翻译器"
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
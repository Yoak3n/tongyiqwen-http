package plugin

import (
	"errors"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"os"
	"tongyiqwen/package/ali_model"
)

type Preset struct {
	Prompt string
}

var preset *Preset
var p *viper.Viper

func init() {
	p = viper.New()
	p.SetConfigType("json")
	p.SetConfigName("preset")
	p.AddConfigPath(".")
	p.AddConfigPath("../")
	err := p.ReadInConfig()
	p.WatchConfig()
	if err != nil {
		return
	}
}

// LoadTextPreset 从preset.json中加载字符串预设
func LoadTextPreset(name string) (string, error) {
	b := p.GetString(name)
	preset = &Preset{Prompt: b}
	if b == "" {
		return "", errors.New("not found preset")
	}
	return b, nil
}

// LoadMapPreset 从preset.json中加载消息预设
func LoadMapPreset(name string) ([]ali_model.Message, error) {
	file, err := os.ReadFile("preset.json")
	if err != nil {
		return nil, err
	}
	r := gjson.ParseBytes(file)
	b := r.Get(name).Array()
	if b == nil {
		return nil, errors.New("not found preset")
	}
	messages := make([]ali_model.Message, 0)
	l := len(b)
	for index, result := range b {
		if index == l {
			if role := result.Get("Role").String(); role != "assistant" {
				break
			}
		}
		message := ali_model.Message{Role: result.Get("Role").String(), Content: result.Get("Content").String()}
		messages = append(messages, message)

	}
	return messages, nil
}

func PushNewTextPreset(name string, content string) error {
	p.Set(name, content)
	err := p.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func PushNewMapPreset(name string, Map []ali_model.Message) error {
	p.Set(name, Map)
	err := p.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func GetAllPreset() (map[string]any, error) {
	return p.AllSettings(), nil
}

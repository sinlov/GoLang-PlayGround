package go_parse

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
)

type Graphics struct {
	Height int `json:"height,omitempty" yaml:"height,omitempty" toml:"height" xml:"height"`
	Width  int `json:"width,omitempty" yaml:"width,omitempty" toml:"width"`
}

type MultiStruct struct {
	Id       int64    `json:"id,omitempty" yaml:"id,omitempty" toml:"id"`
	Name     string   `json:"name,omitempty" yaml:"name,omitempty" toml:"name"`
	Graphics Graphics `json:"graphics,omitempty" yaml:"graphics,omitempty" toml:"graphics"`
}

func multiJsonParse(multi *MultiStruct) (string, error) {
	marshal, err := json.Marshal(multi)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}

func multiYamlParse(multi *MultiStruct) (string, error) {
	marshal, err := yaml.Marshal(multi)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}

func multiTomlParse(multi *MultiStruct)  {

}

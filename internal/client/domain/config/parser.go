package config

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

/* __________________________________________________ */

// Parser represents a configuration format parser.
type Parser interface {
	Unmarshal([]byte) (ConfigObject, error)
	Marshal(ConfigObject) ([]byte, error)
}

/* __________________________________________________ */

// JsonParser implements a JSON parser.
type JsonParser struct{}

// ProvideJsonParser returns a JSON Parser.
func ProvideJsonParser() *JsonParser {
	return &JsonParser{}
}

// Unmarshal parses the given JSON bytes.
func (j *JsonParser) Unmarshal(bytes []byte) (ConfigObject, error) {
	var out = ConfigObject{}
	if err := json.Unmarshal(bytes, &out); err != nil {
		return ConfigObject{}, err
	}
	return out, nil
}

// Marshal marshals the given config map to JSON bytes.
func (j *JsonParser) Marshal(object ConfigObject) ([]byte, error) {
	return json.Marshal(object)
}

/* __________________________________________________ */

// YamlParser implements a YAML parser.
type YamlParser struct{}

// ProvideYamlParser returns a YAML Parser.
func ProvideYamlParser() *YamlParser {
	return &YamlParser{}
}

// Unmarshal parses the given YAML bytes.
func (y *YamlParser) Unmarshal(bytes []byte) (ConfigObject, error) {
	var out = ConfigObject{}
	if err := yaml.Unmarshal(bytes, &out); err != nil {
		return ConfigObject{}, err
	}
	return out, nil
}

// Marshal marshals the given config map to YAML bytes.
func (y *YamlParser) Marshal(object ConfigObject) ([]byte, error) {
	return yaml.Marshal(object)
}

/* __________________________________________________ */

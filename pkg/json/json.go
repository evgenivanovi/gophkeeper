package json

import (
	"bytes"
	"encoding/json"

	"github.com/evgenivanovi/gpl/std"
)

func Json(data interface{}) (string, error) {
	return asJSON(data, std.Empty, std.Empty)
}

func MustJson(data interface{}) string {
	result, err := Json(data)
	if err != nil {
		panic(err)
	}
	return result
}

func PrettyJson(data interface{}) (string, error) {
	return asJSON(data, std.Empty, std.Tab)
}

func MustPrettyJson(data interface{}) string {
	result, err := PrettyJson(data)
	if err != nil {
		panic(err)
	}
	return result
}

func asJSON(data interface{}, prefix, indent string) (string, error) {
	buf := new(bytes.Buffer)

	enc := json.NewEncoder(buf)
	enc.SetIndent(prefix, indent)

	err := enc.Encode(data)
	if err != nil {
		return std.Empty, err
	}

	return buf.String(), nil
}

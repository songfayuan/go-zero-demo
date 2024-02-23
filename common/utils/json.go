package utils

import (
	"bytes"
	"encoding/json"
)

// Encode jsonEncode
func Encode(obj interface{}) (*bytes.Buffer, error) {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false) // 禁用escapeHtml转义
	err := jsonEncoder.Encode(obj)
	if err != nil {
		return nil, err
	}

	return bf, nil
}

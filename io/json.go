package io

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

//SaveJSON ...
func SaveJSON(filePath string, data interface{}) (err error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return
	}

	if err = ioutil.WriteFile(filePath, bytes, 0644); err != nil {
		return
	}

	return
}

//LoadJSON loads data from a local file.
func LoadJSON(filePath string, dataPtr interface{}) (err error) {
	body, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	//Refer to https://stackoverflow.com/questions/31398044/got-error-invalid-character-%C3%AF-looking-for-beginning-of-value-from-json-unmar
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))

	if err = json.Unmarshal(body, dataPtr); err != nil {
		return
	}
	return
}

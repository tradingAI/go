package io

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
)

func SaveGOB(filePath string, data interface{}) (err error) {
	buf, err := SerializeGOB(data)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(filePath, buf, 0644)
	return
}

func LoadGOB(filePath string, dataPtr interface{}) (err error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	err = DeserializeGOB(data, dataPtr)
	return
}

func SerializeGOB(data interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(data); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func DeserializeGOB(data []byte, dataPtr interface{}) (err error) {
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	err = decoder.Decode(dataPtr)
	return
}

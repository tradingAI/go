package io

import (
	"encoding/xml"
	"io/ioutil"
)

//LoadXML ...
func LoadXML(filePath string, dataPtr interface{}) (err error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	err = xml.Unmarshal(content, dataPtr)
	return
}

//SaveXML ...
func SaveXML(filePath string, data interface{}) (err error) {
	bytes, err := xml.Marshal(data)
	if err != nil {
		return
	}
	if err = ioutil.WriteFile(filePath, bytes, 0644); err != nil {
		return
	}
	return
}

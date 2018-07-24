package pkg

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
)

// FileNotExist check file exist
func FileNotExist(name string) bool {
	_, err := os.Stat(name)
	return os.IsNotExist(err)
}

// ReadFile read file if not exist will create a file named name
func ReadFile(name string) (bytes []byte, err error) {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return
	}

	bytes, err = ioutil.ReadAll(f)
	return
}

// WriteFile write data to file
func WriteFile(name string, m map[string]interface{}) (err error) {
	out, err := json.Marshal(m)
	if err != nil {
		return
	}
	bytes := &bytes.Buffer{}
	json.Indent(bytes, out, "", "\t")
	err = ioutil.WriteFile(name, bytes.Bytes(), os.ModePerm)
	return
}

// ClearFile ...
func ClearFile(name string) {
	ioutil.WriteFile(name, []byte("{}"), os.ModePerm)
}

// DeleteFile ...
func DeleteFile(name string) (err error) {
	err = os.Remove(name)
	if err != nil && FileNotExist(name) {
		err = nil
	}
	return
}

// CreateDir ...
func CreateDir(path string) (err error) {
	if FileNotExist(path) {
		err = os.MkdirAll(path, os.ModePerm)
	}
	return
}

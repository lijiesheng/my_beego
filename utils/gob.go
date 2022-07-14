package utils

import (
	"bytes"
	"encoding/gob"
)

// 解码 字节流数据转换为结构体对象
func Decode(value string, r interface{}) error {
	b := []byte(value) // 字符串转换为 byte数组
	network := bytes.NewBuffer(b)
	decoder := gob.NewDecoder(network)  // 获取数据
	if err := decoder.Decode(r); err != nil {
		return err
	}
	return nil

	//network := bytes.NewBuffer([]byte(value))
	//dec := gob.NewDecoder(network)
	//return dec.Decode(r)
}

// 编码 把结构体数据编码成字节流数据
func Encode(value interface{}) (string, error) {
	network := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(network)
	err := enc.Encode(value)
	if err != nil {
		return "", err
	}
	return network.String(), nil
}
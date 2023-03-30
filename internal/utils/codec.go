package utils

import (
	"bytes"
	"encoding/gob"
	"go.uber.org/zap"
	"toolbox/internal/logs"
)

func Encode(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(v); err != nil {
		logs.Warn("encode object failed",
			zap.Error(err),
			zap.Any("obj", v))
		return nil, err
	}
	return buf.Bytes(), nil
}

func Decode(data []byte, v interface{}) error {
	dec := gob.NewDecoder(bytes.NewBuffer(data))
	if err := dec.Decode(v); err != nil {
		logs.Warn("decode object failed",
			zap.Error(err),
			zap.Binary("data", data))
		return err
	}
	return nil
}

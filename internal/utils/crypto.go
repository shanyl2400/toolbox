package utils

import (
	"crypto"
	_ "crypto/sha1"
	"encoding/hex"
	"gopkg.in/mgo.v2/bson"
)

func NewID() string {
	return bson.NewObjectId().Hex()
}

func Hash(text string) string {
	sha1 := crypto.SHA1.New()
	sha1.Write([]byte(text))
	return hex.EncodeToString(sha1.Sum(nil))
}

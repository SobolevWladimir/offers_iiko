package base

import (
	"crypto/sha1"
	"encoding/hex"

	uuid "github.com/gofrs/uuid"
)

func UUID() string {
	u4, _ := uuid.NewV4()
	return u4.String()
}
func SHA1(text string) string {

	hasher := sha1.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
func StringToInterface(str []string) []interface{} {
	result := make([]interface{}, len(str))
	for i, v := range str {
		result[i] = v
	}
	return result
}
func IntToInterface(str []int) []interface{} {
	result := make([]interface{}, len(str))
	for i, v := range str {
		result[i] = v
	}
	return result
}

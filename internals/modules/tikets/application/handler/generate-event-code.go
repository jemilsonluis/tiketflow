package handler

import (
	"math/rand"
	"strings"
)


func randomDigits(leng int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	r := rand.New(rand.NewSource(int64(rand.Int())))

	b := make([]byte, leng)

	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}

	return string(b)
}

func GenerateEventCode(value string) string {
	var code string
	shorName := strings.Split(value, " ")

	for _, value := range shorName {
		code += value[0:1]
	}

	return code + "-"+ randomDigits(6)
}

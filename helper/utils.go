package helper

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

var numRunes = []rune("0123456789")
func RandNumStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = numRunes[rand.Intn(len(numRunes))]
	}
	return string(b)
}

// 取得明天一早00:00的nano時間
func GetNextEarlyNano() (nextEarlyNano int64) {
	now := time.Now()

	nano := int64(1000000000)
	exceed := int64(now.Hour()) * 60 * 60 * nano
	exceed += int64(now.Minute()) * 60 * nano
	exceed += int64(now.Second()) * nano
	exceed += int64(now.Nanosecond())

	nextEarlyNano = now.AddDate(0, 0, 1).UnixNano() - exceed
	return
}

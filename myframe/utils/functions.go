package utils

import (
	"os"
	"io"
	"strconv"
	"crypto/sha1"
	"fmt"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"time"
	"github.com/bitly/go-simplejson"
)

func DuplicateUsers(users []uint64) []uint64 {
	res := []uint64{}
	user_set := make(map[uint64]int)
	for _, user := range users {
		user_set[user] = 1
	}
	for user := range user_set {
		res = append(res, user)
	}
	return res
}
func WritePidToFile(strFileName string) error {
	fw, err := os.OpenFile(strFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	_, err = io.WriteString(fw, strconv.Itoa(os.Getpid()))
	if err != nil {
		return err
	}
	return nil
}
func Sha1Encode(st string) string {
	h := sha1.New()
	h.Write([]byte(st))
	sum := h.Sum(nil)
	return fmt.Sprintf("%x", sum)
}
func Md5sum(src string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(src))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
func ToJSONObject(object interface{}) (*simplejson.Json, error) {
	j, _ := json.Marshal(object)
	return simplejson.NewJson(j)
}
func GetRand(r int) int {
	res := rand.New(rand.NewSource(time.Now().UnixNano()))
	return res.Intn(r)
}

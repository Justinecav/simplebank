package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

//RandInt generates random integer
func RandInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

//RandString generates random string
func RandString(n int) string {

	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

//RandomOwner generates names in a random manner
func RandomOwner() string {
	return RandString(6)
}

//RandomMoney generates money in random manner
func RandomMoney() int64 {
	return RandInt(0, 1000)
}

//RandCurrency genrates currency in random manner
func RandomCurrency() string {
	currencies := []string{INR, USD, EUR}

	n := len(currencies)

	return currencies[rand.Intn(n)]
}

//RansomEmail generates email in random manner
func RandomEmail() string {
	return fmt.Sprintf("%s@yahoo.com", RandString(6))
}

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

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomEmail() string {
	return fmt.Sprintf("%s@%s.%s", RandomString(6), RandomString(4), RandomString(3))
}

func RandomCountry() int32 {
	currencies := []int32{254, 256, 257, 258, 250, 230}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomGender() string {
	currencies := []string{"F", "M"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomFullName() string {
	return fmt.Sprintf("%s %s", RandomString(5), RandomString(5))
}

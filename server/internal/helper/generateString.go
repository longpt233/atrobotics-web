package helper

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	numberSet = "0123456789"
	OTP_LEN   = 6
	letters   = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func GenerateOtp() string {
	var otpValue strings.Builder
	for i := 0; i < OTP_LEN; i++ {
		random := rand.Intn(len(numberSet))
		otpValue.WriteString(string(numberSet[random]))
	}
	inRune := []rune(otpValue.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}

func GenerateOrderCode() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 4)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return "AT" + getDateMonthYear() + string(b)
}
func getDateMonthYear() string {
	now := time.Now()
	date := fmt.Sprint(now.Day())
	month := fmt.Sprint(int(now.Month()))
	if len(month) < 2 {
		month = "0" + month
	}
	if len(date) < 2 {
		date = "0" + date
	}
	year := fmt.Sprint(now.Year())
	return date + month + year[len(year)-2:]
}

package services

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateOTP(phone string) string {
	rand.Seed(time.Now().UnixNano()+ int64(len(phone)))
	lastFive := phone[len(phone)-5:]
	seed,_:=strconv.Atoi(lastFive)
	r := rand.New(rand.NewSource(int64(seed)))
	otp := ""
	for i:= 0; i<6; i++{
		otp += strconv.Itoa(r.Intn(10))
	}
	return  otp
}
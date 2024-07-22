package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(min,max int) int{
	return min + rand.Intn(max-min+1)
}

func RandomString(n int) string{
	var sb strings.Builder
	k := len(alphabet)

	for i:=0;i<n;i++{
		c:=alphabet[rand.Intn(k)]
		sb.WriteByte(c)

	}
	return sb.String()
}

func RandomOwner()string{
	return RandomString(6)
}

func RandomBalance()int{
	return RandomInt(100,10000)
}

func RandomCurrency()string{
	currencies := []string{"INR","USD","CHF","EURO"}

	return currencies[RandomInt(0,len(currencies)-1)]
}
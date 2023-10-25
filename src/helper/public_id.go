package helper

import (
	"crypto/rand"
	"math/big"
	"strings"
)

func GeneratePublicID(len int) string {
	result := insertSeparator(generateRandomString(len), "-", 4)

	return result
}

func GenerateUniqueDefault(intial string, len int) string {
	val := generateRandomString(len)

	combined := intial + "" + val
	return strings.ToUpper(combined)
}

func generateRandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyz123456789"
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[randomIndex.Int64()]
	}

	return strings.ToLower(string(result))
}

func insertSeparator(str, separator string, interval int) string {
	var sb strings.Builder
	length := len(str)

	for i := 0; i < length; i++ {
		if i != 0 && i%interval == 0 {
			sb.WriteString(separator)
		}
		sb.WriteByte(str[i])
	}

	return sb.String()
}

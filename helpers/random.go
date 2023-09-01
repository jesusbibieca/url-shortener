package helpers

import "math/rand"

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func RandomString(n int) string {
	var letter = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, n)
	for i := range result {
		result[i] = letter[rand.Intn(len(letter))]
	}
	return string(result)
}

func RandomEmail() string {
	return RandomString(6) + "@" + RandomString(6) + ".com"
}

func RandomUrl() string {
	return "http://" + RandomString(10) + ".com"
}

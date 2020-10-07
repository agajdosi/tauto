package identity

import (
	"math/rand"
	"time"
)

func seed() {
	rand.Seed(time.Now().UTC().UnixNano())

	return
}

//GenerateName generates a random name
func GenerateName(sex string) (string, string) {
	seed()
	if sex == "M" {
		names := []string{"Jiří", "Jan", "Lukáš", "Petr", "Daniel", "Ondřej", "Adam", "Filip", "Pavel", "Luděk", "Ivan", "Alois", "Jonáš", "Vojtěch"}
		surnames := []string{"Novotný", "Novák", "Černý", "Svoboda", "Dvořák", "Procházka", "Kučera", "Veselý", "Krejčí", "Horák"}

		name := names[rand.Intn(len(names))]
		surname := surnames[rand.Intn(len(surnames))]

		return name, surname
	}

	return "", ""
}

//GeneratePassword generates random password
func GeneratePassword(n int) string {
	seed()
	var letters = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//GenerateBirthday generates a random birthday
func GenerateBirthday() time.Time {
	seed()
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2000, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min

	return time.Unix(sec, 0)
}

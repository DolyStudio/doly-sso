package utils

import "github.com/google/uuid"

func GenerateRandomString() string {
	id := uuid.New()
	return id.String()
}

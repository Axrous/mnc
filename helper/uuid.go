package helper

import "github.com/google/uuid"

func GenerateUUID() string {
	return uuid.NewString()
}
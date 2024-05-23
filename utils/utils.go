package utils

import "github.com/google/uuid"

func StringToUUID(s string) uuid.UUID {
	id, _ := uuid.Parse(s)
	return id
}
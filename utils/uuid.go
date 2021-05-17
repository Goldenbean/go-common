package utils

import (
	guuid "github.com/google/uuid"
	uuid "github.com/satori/go.uuid"
)

func NewUuid() string {
	return guuid.New().String()
}

func SatoriNewUuid() uuid.UUID {
	ret := uuid.Must(uuid.NewV4())
	return ret
}

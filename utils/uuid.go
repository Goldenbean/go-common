package utils

import (
	"fmt"

	guuid "github.com/google/uuid"
	uuid "github.com/satori/go.uuid"
)

func NewUuid() string {
	return guuid.New().String()
}

// Generate : create UUID
func Generate() uuid.UUID {
	//	ret := uuid.NewV4()
	ret := uuid.Must(uuid.NewV4())
	fmt.Printf("UUIDv4: %s\n", ret)

	return ret
}

package common

import (
	"fmt"

	"github.com/satori/go.uuid"
)

// Generate : create UUID
func Generate() uuid.UUID {
	//	ret := uuid.NewV4()
	ret := uuid.Must(uuid.NewV4())
	fmt.Printf("UUIDv4: %s\n", ret)

	return ret
}

func main() {
	Generate()
}

package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUuid(t *testing.T) {
	uuid := NewUuid()
	fmt.Println(" ", uuid)
	assert.Equal(t, 36, len(uuid))
}
func TestSatoriNewUuid(t *testing.T) {
	uuid := SatoriNewUuid()
	fmt.Printf("SatoriNewUuid: %+v \n", uuid)
	//assert.Equal(t, 36, len(uuid))
}

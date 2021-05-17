package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUuid(t *testing.T) {
	uuid := NewUuid()
	fmt.Println(" ", uuid)
	assert.Equal(t, 36, len(uuid))

}

package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func ShouldHaveValidOrderId(t *testing.T) {
	order := Order{}

	assert.Error(t, order.Validate(), "ID is invalid")
}

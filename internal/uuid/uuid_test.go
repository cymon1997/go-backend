package uuid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	uuid := New()
	assert.True(t, IsValid(uuid))
	assert.False(t, IsValid(""))
	assert.False(t, IsValid("12345-abcd"))
}

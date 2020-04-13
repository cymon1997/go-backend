package uuid

import (
	"fmt"
	"github.com/google/uuid"
)

func New() string {
	return fmt.Sprint(uuid.New())
}

func IsValid(id string) bool {
	_, err := uuid.Parse(id)
	if err != nil {
		return false
	}
	return true
}

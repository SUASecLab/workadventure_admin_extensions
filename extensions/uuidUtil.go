package extensions

import (
	"github.com/google/uuid"
)

func IsUUIDValid(requestedUUID string) bool {
	_, err := uuid.Parse(requestedUUID)

	if err != nil {
		return false
	}

	return true
}

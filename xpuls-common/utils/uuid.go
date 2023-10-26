package utils

import (
	"fmt"
	"github.com/google/uuid"
)

func NewUUIDWithPrefix(prefix string) (string, error) {
	id, err := uuid.NewUUID()

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s-%s", prefix, id.String()), nil
}

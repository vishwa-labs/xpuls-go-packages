package utils

import (
	"fmt"
	"github.com/google/uuid"
	"xpuls.io/xpuls-go-packages/common/types"
)

func NewUUIDWithPrefix(prefix types.UniqueIdPrefixes) (string, error) {
	id, err := uuid.NewUUID()

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s-%s", prefix, id.String()), nil
}

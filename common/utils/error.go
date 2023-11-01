package utils

import (
	"github.com/pkg/errors"
	"xpuls-go-packages/common/constants"
)

func IsNotFound(err error) bool {
	return errors.Is(err, constants.ErrNotFound)
}
